package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct{
	HttpVersion string
	RequestTarget string
	Method string
}

type Request struct {
	RequestLine RequestLine
}

var SEPARATOR = "\r\n"
var INCOMPELETE_REQUEST_LINE = fmt.Errorf("incomplete request line")
var UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported HTTP version")

func validHTTP(httpparts []string) bool {
	return len(httpparts) == 2 &&
		httpparts[0] == "HTTP" &&
		httpparts[1] == "1.1" 
}

func RequestFromReader(reader io.Reader)(*Request,error){
	data , err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("could not read request: "),
			err ,
		)
	}

	string := string(data)
	rl , _ , err := parseRequestLine(string)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("could not parse request line: "),
			err,
		)
	}
	return &Request{
		RequestLine: *rl,
	}, err
}

func parseRequestLine(line string)(*RequestLine , string , error){
	idx := strings.Index(line, SEPARATOR)
	if idx == -1{
		return nil, line , nil 
	}

	startline := line[:idx]
	restOfMSG := line[idx+len(SEPARATOR):]

	parts := strings.SplitN(startline, " ", 3)
	if len(parts) != 3 {
		return nil, restOfMSG, INCOMPELETE_REQUEST_LINE
	}

	httpParts := strings.Split(parts[2], "/")

	if !validHTTP(httpParts) {
		return nil, restOfMSG, UNSUPPORTED_HTTP_VERSION 
	}
	r := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   httpParts[1],
	}

	return r,restOfMSG , nil 
}