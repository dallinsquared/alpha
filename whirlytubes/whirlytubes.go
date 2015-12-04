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
	verify() (bool, error)
	Init(string) error
}

type TcpAddress struct {
	Addr string
	Cnxn net.TCPConn
}

func (addr TcpAddress) verify() (b bool, err error) {
	b, err = regexp.MatchString(".*:.*", addr.Addr) //fix connection for address, store connection in address
	return
}

func (addr TcpAddress) Init(a string) error {
	addr.Addr = a
	b, err := addr.verify()
	if err != nil {
		return err
	}
	if !b {
		addr.Addr = nil
		return fmt.Errorf("Address Initialization: malformed address: %s\n\tshould be I.P.add.ress:port\n", a)
	}
	return nil
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
	msg, err := bufio.NewReader(conn).ReadString('\n')
	return
}
