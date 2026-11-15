package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("unix", "/tmp/airlock.sock")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		io.Copy(conn, os.Stdin)
		conn.(*net.UnixConn).CloseWrite()
	}()

	io.Copy(os.Stdout, conn)
}
