package main

import (
	"encrypt/codec"
	"flag"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	os.Exit(intMain())
}

func intMain() int {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %snum licenseName privateKeyFile\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) != 3 {
		flag.Usage()
		return 1
	}
	NUM := flag.Args()[0]
	License := flag.Args()[1]
	PRVFile := flag.Args()[2]
	//MAC3 := flag.Args()[3]
	//filepath := flag.Args()[4]
	mac := NUM + "@" + "This license is generated and authorized by HarmonyCloud company"
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	createCh := make(chan error, 1)
	go func() {
		createCh <- makelicense(mac, License, PRVFile)
	}()

	select {
	case <-interrupt:
		return 130
	case err := <-createCh:
		if err != nil {
			return 1
		}
	}
	return 0
}

func makelicense(mac string, License string, PRVFile string) error {
	//priErr := codec.RSA.InitPRKByFile("/home/lily/Go/src/ostar/private.pem")
	priErr := codec.RSA.InitPRKByFile(PRVFile)
	if priErr != nil {
		fmt.Println("init error:", priErr)
	}

	str, err := codec.RSA.String(mac, codec.MODE_PRIKEY_ENCRYPT)
	//return err
	//fmt.Println("prikey encrypt:", str, err)
	fmt.Println("rsa encrypt base64:" + mac)

	file, err := os.Create(License)
	if err != nil {
		fmt.Println(err)
		return err
	}

	file.WriteString(str)
	defer file.Close()
	return err
}
