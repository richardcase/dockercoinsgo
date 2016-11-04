package certs

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func LoadCertificatesFromFile(certPath string, keyPath string) (tls.Certificate, error) {
	pair, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return pair, fmt.Errorf("Error loading X509 key pair.\nCert: %s\nKey: %s\n%v", certPath, keyPath, err)
	}

	return pair, nil
}

func CACertPoolFromFile(caCert string) (*x509.CertPool, error) {
	pem, err := ioutil.ReadFile(caCert)
	if err != nil {
		return nil, fmt.Errorf("Error reading CA certificate.\nCert: %s\n%v", caCert, err)
	}
	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(pem)
	if !ok {
		return nil, fmt.Errorf("Bad CA certificate")
	}

	return certPool, nil
}
