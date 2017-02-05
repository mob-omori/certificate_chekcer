package main

import (
	"crypto/x509"
	"fmt"
	"net/http"
)

func main() {
	//url := "https://socialnews.rakuten.co.jp/"
	url := "https://github.com"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	_, _ = confirmCertificate(response.TLS.PeerCertificates)

}

func confirmCertificate(certificates []*x509.Certificate) (string, error) {
	for _, v := range certificates {
		fmt.Println(v.MaxPathLen, v.OCSPServer, v.IsCA, v.NotBefore, v.NotAfter)
		fmt.Println(v.Subject.CommonName)
	}

	return "", nil
}
