package main

import (
	"math/big"
	"net"
	"os"
	"time"

	"github.com/koraygocmen/gravitational/certs"
)

// checkKeyCrt checks if the crt and key file exists
// on the specified path, if not it will generate a new
// crt and key file.
func checkKeyCrt() {
	if _, err := os.Stat(grpcServerCrtFile); err != nil {
		if os.IsNotExist(err) {
			crt := certs.CrtTemplate{
				Organization: []string{"Gravitational"},
				IP:           []net.IP{net.ParseIP("127.0.0.1")},
				SerialNumber: big.NewInt(1),
				NotBefore:    time.Now(),
				NotAfter:     time.Now().Add(24 * 365 * time.Hour),
				IsCa:         true,
			}

			err := certs.GenerateCrtKey(grpcServerKeyFile, grpcServerCrtFile, crt)
			if err != nil {
				errs <- err
			}
		}
	}
}
