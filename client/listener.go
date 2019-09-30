package client

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB_Client/responsePackage"
	"log"
	"net"
	"sync"
)

var buffer *bytes.Buffer = bytes.NewBuffer(nil)

var conn *connSyncObj
var once sync.Once

func GetConnSync(address string) *connSyncObj {
	once.Do(func() {
		conn = &connSyncObj{
			Lock: sync.RWMutex{},
			Conn: getConn(address),
		}
	})
	return conn
}

type connSyncObj struct {
	Lock sync.RWMutex
	Conn net.Conn
}

func getConn(address string) net.Conn {
	c, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *connSyncObj) ReceiveResponsePackages() ([]*responsePackage.ResponseDBEventPackage, error) {
	data := make([]byte, 4096, 4096)
	var err error
	var n int = 0
	for  {
		n, _ = c.Conn.Read(data)
		if n > 0 {
			break
		}
	}

	if n < responsePackage.ResponseDBEventPackageHeaderLength {
		buffer.Write(data)
		return nil, errors.New(fmt.Sprintf("Did not receive full response package, response: %s", string(data)))
	}

	buffer.Write(data[:n])

	packages, consumeBytesLength, err:=  responsePackage.ResponseDBEventPackageParser().Parse(buffer)
	if err != nil {
		log.Print(err)
	}
	var consumed []byte = make([]byte, consumeBytesLength, consumeBytesLength)
	_, err = buffer.Read(consumed)
	if err != nil {
		buffer = bytes.NewBuffer(nil)
		log.Printf("Reset the length of buffer failed, consumeBytesLength: %d, buffer length: %d", consumeBytesLength, len(buffer.Bytes()))
	}

	return packages, err
}
