package main

import (
	"fmt"
	"github.com/Diode222/MarioDB_Client/client"
	"github.com/Diode222/MarioDB_Client/requestPackage"
	"log"
	"strings"
	"sync"
)

// TODO Once a package sended, handReceive() should be used once.
func main() {
	connSync := client.GetConnSync("139.155.46.62:45555")
	packs, _ := connSync.ReceiveResponsePackages()
	for _, p := range packs {
		pError := string(p.Error)
		if string(p.Status) == "Error" && strings.Contains(pError, "MAXCLIENTCOUNT") {
			log.Printf(pError)
			connSync.Conn.Close()
			return
		} else if string(p.Status) == "OK" {
			continue
		}
	}
	fmt.Println("来不来")

	pack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   8,
		DBNameLength:   4,
		KeysLength:     10,
		ValuesLength:   0,
		StartLength:    0,
		LimitLength:    0,
		PrefixLength:   0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("BATCHGET"),
		DBName:         []byte("LVYA"),
		Keys:           []byte("abcde##adc"),
		Values:         nil,
		Start:          nil,
		Limit:          nil,
		Prefix:         nil,
		Settings:       nil,
		Reserved:       nil,
	}

	createPack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   6,
		DBNameLength:   4,
		KeysLength:     0,
		ValuesLength:   0,
		StartLength:    0,
		LimitLength:    0,
		PrefixLength:   0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("CREATE"),
		DBName:         []byte("LVYA"),
		Keys:           nil,
		Values:         nil,
		Start:          nil,
		Limit:          nil,
		Prefix:         nil,
		Settings:       nil,
		Reserved:       nil,
	}

	connSync.Lock.Lock()
	go createPack.Pack(connSync.Conn)
	connSync.Lock.Unlock()
	packages, _ := connSync.ReceiveResponsePackages()
	for _, p := range packages {
		fmt.Println("create")
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
		fmt.Println(string(p.Error))
	}

	batchPutPack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   8,
		DBNameLength:   4,
		KeysLength:     23,
		ValuesLength:   23,
		StartLength:    0,
		LimitLength:    0,
		PrefixLength:   0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("BATCHPUT"),
		DBName:         []byte("LVYA"),
		Keys:           []byte("abcde##adc##heihei##abc"),
		Values:         []byte("12345##123##nono##diode"),
		Start:          nil,
		Limit:          nil,
		Prefix:         nil,
		Settings:       nil,
		Reserved:       nil,
	}

	connSync.Lock.Lock()
	batchPutPack.Pack(connSync.Conn)
	connSync.Lock.Unlock()
	packages, _ = connSync.ReceiveResponsePackages()
	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
		fmt.Println(string(p.Error))
	}

	batchDeletePack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   11,
		DBNameLength:   4,
		KeysLength:     10,
		ValuesLength:   0,
		StartLength:    0,
		LimitLength:    0,
		PrefixLength:   0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("BATCHDELETE"),
		DBName:         []byte("LVYA"),
		Keys:           []byte("abcde##adc"),
		Values:         nil,
		Start:          nil,
		Limit:          nil,
		Prefix:         nil,
		Settings:       nil,
		Reserved:       nil,
	}

	connSync.Lock.Lock()
	batchDeletePack.Pack(connSync.Conn)
	connSync.Lock.Unlock()
	packages, _ = connSync.ReceiveResponsePackages()
	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
		fmt.Println(string(p.Error))
	}

	putPack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   3,
		DBNameLength:   4,
		KeysLength:     5,
		ValuesLength:   5,
		StartLength:    0,
		LimitLength:    0,
		PrefixLength:   0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("PUT"),
		DBName:         []byte("LVYA"),
		Keys:           []byte("abcde"),
		Values:         []byte("12346"),
		Start:          nil,
		Limit:          nil,
		Prefix:         nil,
		Settings:       nil,
		Reserved:       nil,
	}

	connSync.Lock.Lock()
	putPack.Pack(connSync.Conn)
	connSync.Lock.Unlock()
	packages, _ = connSync.ReceiveResponsePackages()
	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
		fmt.Println(string(p.Error))
	}

	rangePack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   5,
		DBNameLength:   4,
		KeysLength:     0,
		ValuesLength:   0,
		StartLength:    3,
		LimitLength:    5,
		PrefixLength:   0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("RANGE"),
		DBName:         []byte("LVYA"),
		Keys:           nil,
		Values:         nil,
		Start:          []byte("abc"),
		Limit:          []byte("abcde"), // [Start, Limit)
		Prefix:         nil,
		Settings:       nil,
		Reserved:       nil,
	}

	connSync.Lock.Lock()
	rangePack.Pack(connSync.Conn)
	connSync.Lock.Unlock()
	packages, _ = connSync.ReceiveResponsePackages()
	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
		fmt.Println(string(p.Reserved))
		fmt.Println(string(p.Error))
	}

	seekRangePack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   9,
		DBNameLength:   4,
		KeysLength:     5,
		ValuesLength:   0,
		StartLength:    0,
		LimitLength:    0,
		PrefixLength:   0,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("SEEKRANGE"),
		DBName:         []byte("LVYA"),
		Keys:           []byte("abcdf"),
		Values:         nil,
		Start:          nil,
		Limit:          nil, // [Start, Limit)
		Prefix:         nil,
		Settings:       nil,
		Reserved:       nil,
	}

	connSync.Lock.Lock()
	seekRangePack.Pack(connSync.Conn)
	connSync.Lock.Unlock()
	packages, _ = connSync.ReceiveResponsePackages()
	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
		fmt.Println(string(p.Reserved))
		fmt.Println(string(p.Error))
	}

	prefixRangePack := &requestPackage.RequestDBEventPackage{
		Version:        [2]byte{'V', '1'},
		MethodLength:   11,
		DBNameLength:   4,
		KeysLength:     0,
		ValuesLength:   0,
		StartLength:    0,
		LimitLength:    0,
		PrefixLength:   3,
		SettingsLength: 0,
		ReservedLength: 0,
		Method:         []byte("PREFIXRANGE"),
		DBName:         []byte("LVYA"),
		Keys:           nil,
		Values:         nil,
		Start:          nil,
		Limit:          nil, // [Start, Limit)
		Prefix:         []byte("abc"),
		Settings:       nil,
		Reserved:       nil,
	}

	connSync.Lock.Lock()
	prefixRangePack.Pack(connSync.Conn)
	connSync.Lock.Unlock()
	packages, _ = connSync.ReceiveResponsePackages()
	for _, p := range packages {
		fmt.Println(string(p.Status))
		fmt.Println(string(p.Values))
		fmt.Println(string(p.Reserved))
		fmt.Println(string(p.Error))
	}

	doneChan := make(chan bool, 1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
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
		fmt.Println("status: " + string(p.Status))
		fmt.Println("values: " + string(p.Values))
		fmt.Println("error: " + string(p.Error))
	}

	wg.Done()
	<-doneChan
}
