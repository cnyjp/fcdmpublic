//go:build windows
// +build windows

package model

import (
	"context"
	"errors"
	"github.com/Microsoft/go-winio"
	"net"
	"os"
	"strconv"
	"time"
)

// generate local socket name
func genLocalSocketName() string {
	return genLocalSocketNameByPid(os.Getpid())
}
func genLocalSocketNameByPid(pid int) string {
	return genLocalSocketNameByName(strconv.Itoa(pid))
}
func genLocalSocketNameByName(name string) string {
	return `\\.\pipe\fcdm_plugin_pipe_` + name
}

func (lss *LocalSocketServer) platformListen() error {
	if lss.listener != nil {
		return errors.New("the LocalSocket is already listened, can not listen again")
	}

	l, err := winio.ListenPipe(lss.name, nil)
	if nil != err {
		return err
	}
	lss.listener = l
	return nil
}

func (lsc *LocalSocketClient) platformDail(ctx context.Context, timeout time.Duration) (conn net.Conn, err error) {
	errchan := make(chan error)
	var ctx1 context.Context
	if nil == ctx {
		ctx1 = context.Background()
	} else {
		ctx1 = context.WithValue(ctx, "dail", "dail")
	}
	go func() {
		//if nil == ctx{
		//	if 0 == timeout{
		//		//conn,err = winio.DialPipeContext(ctx1, lsc.name)
		//		conn,err = winio.DialPipe(lsc.name, &timeout)
		//	}else{
		//		conn,err = winio.DialPipe(lsc.name, &timeout)
		//	}
		//}else{
		//	if 0 == timeout {
		//		conn,err = winio.DialPipe(lsc.name, &timeout)
		//	}else{
		//		conn, err = winio.DialPipe(lsc.name, &timeout)
		//	}
		//}
		if 0 == timeout {
			timeout = time.Second * 30 //default use 30 seconds for timeout
		}
		defer close(errchan)
		conn, err = winio.DialPipe(lsc.name, &timeout)
		errchan <- err
	}()
	select {
	case err = <-errchan:
		{
			return conn, err
		}
	case <-ctx1.Done():
		{
			err = ctx1.Err()
			return conn, err
		}
	}
}
