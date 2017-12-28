package main

import (
	"encrypt/codec"
	"log"
)

func test_rsa() {
	pubErr, priErr := codec.RSA.InitByFile("/home/lily/Go/src/ostar/public.pem", "/home/lily/Go/src/ostar/private.pem")
	log.Println("init error:", pubErr, priErr)

	str, err := codec.RSA.String("OKOK", codec.MODE_PRIKEY_ENCRYPT)
	log.Println("prikey encrypt:", str, err)
	str, err = codec.RSA.String(str, codec.MODE_PUBKEY_DECRYPT)
	log.Println("pubkey decrypt:", str, err)

}
