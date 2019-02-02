package main

import (
	"fmt"
	"strings"

	"github.com/jmoiron/jsonq"
	"github.com/CaliDog/certstream-go"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

func handleCertificate(jq jsonq.JsonQuery) {
	// pull the list of all the domains named in the leaf certificate (CN and SANs)
	domains, err := jq.ArrayOfStrings("data", "leaf_cert", "all_domains")

	// pull the certificate fingerprint and use it to get the crt.sh URL
	fingerprint, err := jq.String("data", "leaf_cert", "fingerprint")
	if err != nil {
		log.Error(err.Error())
		return
	}
	certURL := fmt.Sprintf("https://crt.sh/?q=%s", strings.Replace(fingerprint, ":", "", -1))
	log.Info("Cert data -> ", certURL)

	for _, domain := range domains {
		log.Info("Found: -> ", domain)
	}
	log.Info("-----------------------")
}


func main() {
	// The false flag specifies that we want heartbeat messages.
	stream, errStream := certstream.CertStreamEventStream(false)
	for {
		select {
			case jq := <-stream:
				handleCertificate(jq)
      
			case err := <-errStream:
				log.Error(err.Error())
		}
	}
}