//go:build unix
// +build unix

package model

import (
	"errors"
	"net"
	"os"
	"strconv"
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
	//check if the local socket exists, if exists, try to delete the socket. if can not delete, then return error.
	stat, err := os.Stat(lss.name)

	if nil != err {
		if os.IsExist(err) {
			err := os.Remove(lss.name)
			if nil != err { //todo:complete err handle
				return err
			}
		}
	}

	if nil != stat {
		//file exist， maybe it is not removed by last, try to remove it
		os.Remove(lss.name)
	}

	addr, err := net.ResolveUnixAddr("unix", lss.name)
	if err != nil {
		return err
	}
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		return err
	}
	lss.listener = listener
	return nil
}

func (lsc *LocalSocketClient) platformDail() (net.Conn, error) {
	conn, err := net.DialTimeout("unix", lsc.name, 10)
	return conn, err
}
