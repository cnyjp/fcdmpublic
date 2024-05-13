//go:build windows
// +build windows

package model

import (
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

func (lsc *LocalSocketClient) platformDail() (net.Conn, error) {
	timeout := time.Second * 10
	conn, err := winio.DialPipe(lsc.name, &timeout)
	return conn, err
}
