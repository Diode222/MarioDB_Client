package main

import (
	"fmt"
	"github.com/Diode222/MarioDB_Client/client"
	"github.com/Diode222/MarioDB_Client/requestPackage"
	"sync"
)



func main() {
	connSync := client.GetConnSync("127.0.0.1:50000")
	fmt.Println("来不来")

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
		connSync.Lock.Lock()
		pack1.Pack(connSync.Conn)
		connSync.Lock.Unlock()
		wg.Done()
	}()

	go func() {
		connSync.Lock.Lock()
		pack2.Pack(connSync.Conn)
		connSync.Lock.Unlock()
		wg.Done()
	}()
	wg.Wait()

	packages, err := connSync.ReceiveResponsePackages()
	if err != nil {
		panic(err)
	}

	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
	}
}
