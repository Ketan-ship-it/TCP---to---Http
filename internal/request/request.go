package request

import (
	"bytes"
	"fmt"
	"io"
)

type parserState string
const (
	STATEINIT parserState = "init"
	STATEDONE parserState = "Done"
	STATEERROR parserState = "Error"
)

type RequestLine struct{
	HttpVersion string
	RequestTarget string
	Method string
}

type Request struct {
	RequestLine RequestLine
	State parserState
}

var SEPARATOR = []byte("\r\n")
var INCOMPELETE_REQUEST_LINE = fmt.Errorf("incomplete request line")
var UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported HTTP version")
var ERROR_STATE = fmt.Errorf("cannot parse request in error state")

func newRequest() *Request {
	return &Request{
		State: STATEINIT,
	}
}

func validHTTP(httpparts [][]byte) bool {
	return len(httpparts) == 2 &&
		string(httpparts[0]) == "HTTP" &&
		string(httpparts[1]) == "1.1" 
}

func RequestFromReader(reader io.Reader)(*Request,error){
	request := newRequest()

	buffer := make([]byte , 1024)
	buflen := 0
	for !request.isCompelete() && !request.isError() {
		data , err := reader.Read(buffer[buflen:])
		if err != nil {
			return nil , err 
		}
		buflen += data
		readN,err := request.parse(buffer[:buflen])
		if err != nil {
			return nil, err 
		}
		copy(buffer , buffer[readN:buflen])
		buflen -= readN
		
	}
	return request , nil
}

func (r *Request) parse(data []byte) (int,error) {
	read:=0
OUTER:
	for {
		switch r.State {
			case STATEERROR:
				return 0, ERROR_STATE	
			case STATEINIT:
				rl , n , err := parseRequestLine(data[read:])
				if err != nil {
					r.State = STATEERROR
					return 0, err
				}
				if n==0 {
					break OUTER
				}
				r.RequestLine = *rl
				read += n
				r.State = STATEDONE
			case STATEDONE:
				break OUTER
		}
	}
	return read ,nil
}

func (r *Request) isCompelete() bool {
	return r.State == STATEDONE
}

func (r *Request) isError() bool {
	return r.State == STATEERROR
}

func parseRequestLine(line []byte)(*RequestLine , int , error){
	idx := bytes.Index(line, SEPARATOR)
	if idx == -1{
		return nil, 0 , nil 
	}

	startline := line[:idx]
	read := idx+len(SEPARATOR)

	parts := bytes.SplitN(startline, []byte(" "), 3)
	if len(parts) != 3 {
		return nil, 0, INCOMPELETE_REQUEST_LINE
	}

	httpParts := bytes.Split(parts[2], []byte("/"))

	if !validHTTP(httpParts) {
		return nil, 0, UNSUPPORTED_HTTP_VERSION 
	}
	r := &RequestLine{
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
		HttpVersion:   string(httpParts[1]),
	}

	return r, read , nil 
}