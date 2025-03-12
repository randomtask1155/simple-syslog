package main

import (
	"log"
	"os"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	/*certFile := os.Getenv("SIMPLE_SYSLOG_CERT")
	keyFile := os.Getenv("SIMPLE_SYSLOG_KEY")
	certPem, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Fatal(err)
	}
	keyPem, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("%s\n", keyPem))
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert}}*/
	server := syslog.NewServer()
	//server.SetFormat(syslog.RFC5424)
	server.SetFormat(noFormat)
	server.SetHandler(handler)
	err := server.ListenTCP("0.0.0.0:" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			log.Println(logParts)
		}
	}(channel)

	server.Wait()
}
