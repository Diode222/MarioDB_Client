package main

import (
	"fmt"
	"github.com/Diode222/MarioDB_Client/client"
	"github.com/Diode222/MarioDB_Client/requestPackage"
	"log"
	"sync"
)

func main() {
	connSync := client.GetConnSync("127.0.0.1:50000")
	fmt.Println("来不来")

	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//pack1 := &requestPackage.RequestDBEventPackage{
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
	//
	//pack2 := &requestPackage.RequestDBEventPackage{
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
	//go func() {
	//	connSync.Lock.Lock()
	//	pack1.Pack(connSync.Conn)
	//	connSync.Lock.Unlock()
	//	wg.Done()
	//}()
	//
	//go func() {
	//	connSync.Lock.Lock()
	//	pack2.Pack(connSync.Conn)
	//	connSync.Lock.Unlock()
	//	wg.Done()
	//}()
	//wg.Wait()
	//
	//packages, err := connSync.ReceiveResponsePackages()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, p := range packages {
	//	fmt.Println(string(p.Status))
	//	fmt.Println(string(p.Values))
	//}

	pack := &requestPackage.RequestDBEventPackage{
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

	doneChan := make(chan bool, 1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go handleSend(pack, connSync, doneChan, wg)
		go handleReceive(connSync, doneChan, wg)
	}
	wg.Wait()
}

func handleSend(pack *requestPackage.RequestDBEventPackage, connSync *client.ConnSyncObj, doneChan chan bool, wg *sync.WaitGroup) {
	doneChan <- true
	fmt.Println("send")
	err := pack.Pack(connSync.Conn)
	if err != nil {
		log.Printf("packages pack failed, err: %s", err.Error())
	}

	wg.Done()
}

func handleReceive(connSync *client.ConnSyncObj, doneChan chan bool, wg *sync.WaitGroup) {
	packages, err := connSync.ReceiveResponsePackages()
	fmt.Println("receive")
	if err != nil {
		log.Printf("packages analysis failed, err: %s", err.Error())
		<-doneChan
	}
	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
	}

	wg.Done()
	<-doneChan
}
