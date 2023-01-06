package darajago

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"os"
)

func openSSlEncrypt(data, certPath string) (string, error) {
	// read the certificate
	certFile, err := os.Open(certPath)
	if err != nil {
		return "", err
	}
	defer certFile.Close()
	certBytes, err := io.ReadAll(certFile)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(certBytes)
	if block == nil {
		return "", errors.New("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}

	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, cert.PublicKey.(*rsa.PublicKey), []byte(data), nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil

}
