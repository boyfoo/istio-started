package main

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"io/ioutil"
	"log"
	"os"
)

func pubKey() []byte {
	f, err := os.Open("./mypub.pem")
	if err != nil {
		log.Fatalln(err)
	}

	all, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	return all
}

func main() {

	key, err := jwk.ParseKey(pubKey(), jwk.WithPEM(true))
	if err != nil {
		log.Fatalln(err)
	}

	if publicKey, ok := key.(jwk.RSAPublicKey); ok {
		b, err := json.Marshal(publicKey)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b))
	}
}
