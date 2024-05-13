package model

import (
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
	listener net.Listener //server listenrer
	quitchan chan int
}

type LocalSocketClient struct {
	name string
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

func CreateLocalSocketClientNamed(name string) *LocalSocketClient {
	return &LocalSocketClient{
		name: genLocalSocketNameByName(name),
	}
}

func (lsc *LocalSocketClient) SendRequest(request LocalSocketRequest) LocalSocketResponse {
	conn, err := lsc.platformDail()

	if err != nil {
		return request.GenErrorResponse("", err.Error(), "")
	}
	defer func() {
		conn.Close()
	}()

	conn.SetReadDeadline(time.Now().Add(time.Second * 20))

	rba, err := json.Marshal(request)

	_, err = conn.Write(rba)

	if nil != err {
		return request.GenErrorResponse("", err.Error(), "")
	}

	rsab, err := io.ReadAll(conn)

	if nil != err {
		//read error, but the content maybe right, so try to unmarshal read content.
	}

	response := LocalSocketResponse{}

	err = json.Unmarshal(rsab, &response)

	if err != nil {
		return request.GenErrorResponse("", err.Error(), "")
	}

	return response
}
