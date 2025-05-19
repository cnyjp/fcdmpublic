package model

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"time"
)

// LocalSocketRequest Request body for local socket
type LocalSocketRequest struct {
	Command  string `json:"command"`
	Sequence string `json:"sequence"`
	Data     string `json:"data"`
}

func (lsr *LocalSocketRequest) GenSuccessResponse(message string, data string) LocalSocketResponse {
	return LocalSocketResponse{
		Command:  lsr.Command,
		Success:  true,
		Sequence: lsr.Sequence,
		Data:     data,
		Message:  message,
	}
}

func (lsr *LocalSocketRequest) GenErrorResponse(code string, message string, data string) LocalSocketResponse {
	return LocalSocketResponse{
		Command:  lsr.Command,
		Success:  false,
		Sequence: lsr.Sequence,
		Code:     code,
		Data:     data,
		Message:  message,
	}
}

// LocalSocketResponse
// Response body for local socket
type LocalSocketResponse struct {
	Command     string   `json:"command"`
	Success     bool     `json:"success"`
	Sequence    string   `json:"sequence"`
	Code        string   `json:"code"` //error code
	Data        string   `json:"data"`
	Message     string   `json:"message"`
	MessageVars []string `json:"messageVars,omitempty"`
}

// StartLocalSocketServer
// start a new local socket server
func StartLocalSocketServer(handleFunc SocketCommandFunc) (*LocalSocketServer, error) {
	return StartLocalSocketServerNamed(handleFunc, genLocalSocketName())
}

func StartLocalSocketServerNamed(handleFunc SocketCommandFunc, name string) (*LocalSocketServer, error) {
	server := LocalSocketServer{
		name:     name,
		listener: nil,
		quitchan: nil,
	}
	err := server.start(handleFunc)
	if nil != err {
		return nil, err
	}
	return &server, nil
}

func (lss *LocalSocketServer) start(handleFunc SocketCommandFunc) error {
	//start li
	err := lss.platformListen()
	if nil != err {
		return err
	}
	go func() {
		for {
			conn, err := lss.listener.Accept()
			if nil != err {
				return //stop handle, if error
			}

			defer func() {
				conn.Close()
			}()

			var buffer [1024 * 1024]byte

			n, err := conn.Read(buffer[0:])

			rqab := buffer[0:n]

			request := LocalSocketRequest{}

			err = json.Unmarshal(rqab, &request)

			response := LocalSocketResponse{}
			if nil != err {
				response = request.GenErrorResponse("", err.Error(), "")
			} else {
				response = handleFunc(request)
			}

			rsab, err := json.Marshal(response)

			if nil != err {
				if nil == rsab {
					rqab = []byte{} //if error, return an empty string
				}
			}

			conn.Write(rsab)

			conn.Close()
		}
	}()
	return nil
}

func (lss *LocalSocketServer) Close() error {
	//关闭监听
	defer func() {
		recover()
	}()
	if nil != lss {
		lss.listener.Close()
	}
	return nil
}

type SocketCommandFunc func(request LocalSocketRequest) (response LocalSocketResponse)

type LocalSocketServer struct {
	name     string
	listener net.Listener //server listener
	quitchan chan int
}

type LocalSocketClient struct {
	name           string
	defaultTimeout int //
}

func CreateLocalSocketClient() *LocalSocketClient {
	return &LocalSocketClient{
		name: genLocalSocketName(),
	}
}

func CreateLocalSocketClientByPid(pid int) *LocalSocketClient {
	return &LocalSocketClient{
		name: genLocalSocketNameByPid(pid),
	}
}

func CreateLocalSocketClientByName(name string) *LocalSocketClient {
	return &LocalSocketClient{
		name: genLocalSocketNameByName(name),
	}
}

func (lsc *LocalSocketClient) SendRequest(ctx context.Context, request LocalSocketRequest, timeout time.Duration) LocalSocketResponse {
	conn, err := lsc.platformDail(ctx, timeout)
	if err != nil {
		return request.GenErrorResponse("", err.Error(), "")
	}
	defer func() {
		conn.Close()
	}()

	if timeout > 0 {
		conn.SetDeadline(time.Now().Add(timeout))
	} else {
		conn.SetDeadline(time.Now().Add(time.Second * 30)) //default use 30 second for timeout
	}

	var response LocalSocketResponse
	errchan := make(chan error)
	go func() {

		rba, _ := json.Marshal(request)

		_, err = conn.Write(rba)

		if nil != err {
			response = request.GenErrorResponse("", err.Error(), "")
			errchan <- err
			return
		}

		rsab, readError := io.ReadAll(conn)

		if nil != err {
			//read error, but the content maybe right, so try to unmarshal read content.
		}

		err = json.Unmarshal(rsab, &response)

		if err != nil {
			if nil != readError {
				response = request.GenErrorResponse("", readError.Error(), "")
				errchan <- readError
			} else {
				response = request.GenErrorResponse("", err.Error(), "")
				errchan <- err
			}
			return
		}
		errchan <- nil
	}()
	var ctx1 context.Context
	if nil == ctx {
		ctx1 = context.Background()
	} else {
		ctx1 = context.WithValue(ctx, "request", "request")
	}
	select {
	case err = <-errchan:
		{
			return response
		}
	case <-ctx1.Done():
		{
			return request.GenErrorResponse("", ctx1.Err().Error(), "")
		}
	}
}
