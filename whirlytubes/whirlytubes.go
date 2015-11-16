package whirlytubes

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
)

type Address interface {
	Send(string) error
	Receive() (string, error)
	Verify() (bool, error)
}

type TcpAddress struct {
	Addr string
	Cnxn net.TCPConn
}

func (addr TcpAddress) Verify() (b bool, err error) {
	b, err = regexp.MatchString(".*:.*", addr.Addr) //fix connection for address, store connection in address
	return
}

func (addr TcpAddress) Send(msg string) (err error) {
	conn, err := net.Dial("tcp", addr.Addr)
	if err != nil {
		return
	}
	_, err = fmt.Fprint(conn, msg) //add writing capabilities
}

func (addr TcpAddress) Receive() (msg string, err error) {
	conn, err := net.Dial("tcp", addr.Addr)
	if err != nil {
		return "", err
	}
}
