package certs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"os"
	"time"
)

// CrtTemplate gives a simple interface to pass details
// of the certificate to the generate function.
type CrtTemplate struct {
	Organization []string
	IP           []net.IP
	SerialNumber *big.Int
	NotBefore    time.Time
	NotAfter     time.Time
	IsCa         bool
}

// GenerateCrtKey generates key and crt for the given path.
func GenerateCrtKey(keyPath, crtPath string, c CrtTemplate) error {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return errors.New("error creating private key" + err.Error())
	}

	template := x509.Certificate{
		SerialNumber: c.SerialNumber,

		Subject: pkix.Name{
			Organization: c.Organization,
		},

		NotBefore: c.NotBefore,
		NotAfter:  c.NotAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,

		IPAddresses: c.IP,
		IsCA:        c.IsCa,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return errors.New("error creating crt file" + err.Error())
	}

	certOut, err := os.Create(crtPath)
	if err != nil {
		return errors.New("error opening to crt file" + err.Error())
	}

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return errors.New("error writing to crt file" + err.Error())

	}

	if err := certOut.Close(); err != nil {
		return errors.New("error closing crt file" + err.Error())
	}

	keyOut, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.New("error opening key file" + err.Error())
	}

	pemBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}

	if err = pem.Encode(keyOut, &pemBlock); err != nil {
		return errors.New("error writing to key file" + err.Error())
	}

	if err := keyOut.Close(); err != nil {
		return errors.New("error closing key file" + err.Error())
	}

	return nil
}
