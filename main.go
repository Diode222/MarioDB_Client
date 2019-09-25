package main

import (
	"fmt"
	"github.com/Diode222/MarioDB_Client/requestPackage"
	"net"
	"sync"
)

var conn connSyncObj = connSyncObj{
	Lock: sync.RWMutex{},
	Conn: getConn(),
}

type connSyncObj struct {
	Lock sync.RWMutex
	Conn net.Conn
}

func getConn() net.Conn {
	c, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		panic(err)
	}
	return c
}

func main() {
	//listener := client.Listener("127.0.0.1", 49999, false, 1)
	//listener.Init()
	fmt.Println("来不来")

	//conn.Lock.Lock()
	//conn.Conn.Write([]byte("how to close this?"))
	//conn.Lock.Unlock()

	var wg sync.WaitGroup
	wg.Add(2)

	pack1 := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   3,
		DBNameLength:   4,
		KeysLength:     10,
		ValuesLength:   10,
		StartsLength:   0,
		LimitsLength:   0,
		PrefixesLength: 0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("GET"),
		DBName:         []byte("LVYA"),
		Keys:           []byte("xixi##haha"),
		Values:         []byte("niubi##xyz"),
		Starts:         nil,
		Limits:         nil,
		Prefixes:       nil,
		Settings:       nil,
		Reserved:       nil,
	}

	pack2 := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   3,
		DBNameLength:   4,
		KeysLength:     10,
		ValuesLength:   10,
		StartsLength:   0,
		LimitsLength:   0,
		PrefixesLength: 0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("GET"),
		DBName:         []byte("yang"),
		Keys:           []byte("yesy##nono"),
		Values:         []byte("xiyou##qwe"),
		Starts:         nil,
		Limits:         nil,
		Prefixes:       nil,
		Settings:       nil,
		Reserved:       nil,
	}

	go func() {
		conn.Lock.Lock()
		pack1.Pack(conn.Conn)
		conn.Lock.Unlock()
		wg.Done()
	}()
	//pack1 := requestPackage.RequestDBEventPackage{
	//	Version:        [2]byte{'V', '1'},
	//	MethodLength:   3,
	//	DBNameLength:   4,
	//	KeysLength:     10,
	//	ValuesLength:   10,
	//	StartsLength:   0,
	//	LimitsLength:   0,
	//	PrefixesLength: 0,
	//	SettingsLength: 0,
	//	ReservedLength: 0,
	//	Method:         []byte("GET"),
	//	DBName:         []byte("LVYA"),
	//	Keys:           []byte("xixi##haha"),
	//	Values:         []byte("niubi##xyz"),
	//	Starts:         nil,
	//	Limits:         nil,
	//	Prefixes:       nil,
	//	Settings:       nil,
	//	Reserved:       nil,
	//}
	//pack1.Pack(conn)

	go func() {
		conn.Lock.Lock()
		pack2.Pack(conn.Conn)
		conn.Lock.Unlock()
		wg.Done()
	}()
	wg.Wait()

	//pack2 := requestPackage.RequestDBEventPackage{
	//	Version:        [2]byte{'V', '1'},
	//	MethodLength:   3,
	//	DBNameLength:   4,
	//	KeysLength:     10,
	//	ValuesLength:   10,
	//	StartsLength:   0,
	//	LimitsLength:   0,
	//	PrefixesLength: 0,
	//	SettingsLength: 0,
	//	ReservedLength: 0,
	//	Method:         []byte("GET"),
	//	DBName:         []byte("yang"),
	//	Keys:           []byte("yesy##nono"),
	//	Values:         []byte("xiyou##qwe"),
	//	Starts:         nil,
	//	Limits:         nil,
	//	Prefixes:       nil,
	//	Settings:       nil,
	//	Reserved:       nil,
	//}
	//
	//pack2.Pack(conn)
	//
	//buf := new(bytes.Buffer)
	//pack2.Pack(buf)
	//
	//zzz := new(requestPackage.RequestDBEventPackage)
	//for _, b := range buf.Bytes() {
	//	fmt.Print(b)
	//	fmt.Print(", ")
	//}
	//fmt.Println()
	//zzz.Unpack(buf)
	//
	//xxx := []byte{86, 49, 0, 3, 0, 4, 0, 10, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 71, 69, 84, 121, 97, 110, 103, 121, 101, 115, 121, 35, 35, 110, 111, 110, 111, 120, 105, 121, 111, 117, 35, 35, 113, 119, 101}
	//
	//yyy := new(requestPackage.RequestDBEventPackage)
	//yyy.Unpack(bytes.NewReader(xxx))
	//
	//fmt.Println(string(yyy.Version[0]))
	//fmt.Println(string(yyy.DBName))
	//fmt.Println(yyy.KeysLength)
	//fmt.Println(">>>>>>>>>>>>>>>")
	//
	//fmt.Println(string(zzz.Version[0]))
	//fmt.Println(string(zzz.DBName))
}
