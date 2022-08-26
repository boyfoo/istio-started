package main

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

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

func pubKey() []byte {
	return read("./mypub.pem")
}

func read(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	all, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	return all
}

func token() {

	privkey, err := jwk.ParseKey(read("./myrsa.pem"), jwk.WithPEM(true))
	if err != nil {
		fmt.Printf("failed to parse JWK: %s\n", err)
		return
	}

	fmt.Println("PRIKEY OK")

	//pKey, ok := privkey.(jwk.RSAPublicKey)
	//
	//if !ok {
	//	log.Fatalln("jwk.RSAPublicKey", err)
	//}

	token, err := jwt.NewBuilder().Issuer("user@jtthink.com").Subject(strconv.FormatInt(time.Now().UnixNano(), 10)).Build()
	if err != nil {
		log.Fatalln(123123, err)
	}
	signed, err := jwt.Sign(token, jwa.RS256, privkey)
	if err != nil {
		fmt.Printf("failed to sign token: %s\n", err)
		return
	}
	fmt.Println(string(signed))
}
