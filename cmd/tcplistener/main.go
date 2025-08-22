package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const inputFilePath = "message.txt"

func getLinesChannel(f io.ReadCloser) <- chan string {
	out:= make(chan string ,1)
	go func() {
		defer close(out)
		defer f.Close()

		str :=""
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}

			b = b[:n]
			if i:=bytes.IndexByte(b,'\n'); i!=-1{
				str += string(b[:i])
				b=b[i+1:]
				out <- str
				str=""
			}
			str+= string(b)
		}
		if str != "" {
			out <- str
		}
	}()
	return out
}

func main() {
	Listener, err := net.Listen("tcp",":42069")
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}
	
	for {
		conn , err := Listener.Accept()
		if err != nil {
			log.Printf("error accepting connection: %s\n", err)
			continue
		}
		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("read: %s\n", line)
		}
	}
	
}