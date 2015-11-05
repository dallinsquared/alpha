package whirlytubes

import (
	"net"
	"regexp"
)

type Address interface {
	Send(string) error
	Receive() (string, error)
	Verify() (bool, error)
}

type TcpAddress net.Conn

func (addr TcpAddress) Verify() (b bool, err error) {
	b, err = regexp.MatchString(".*:.*", addr) //fix connection for address, store connection in address
	return
}

func (addr TcpAddress) Send(msg string) (err error) {
}
