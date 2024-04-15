package appsec

import (
	"appconf"
	
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"math/big"

	"log"
	"os"
	"strconv"
	"time"
)

// GetSecret returns the JWT secret
func GetSecret() string {
	
	if appconf.GetVal("gin_mode") == "debug" {
		return appconf.GetVal("jwt_secret")
	}

	secret, err := generateSecret()
	if err != nil {
		log.Fatal("Error generating JWT secret")
	}

	return secret
}

func generateSecret() (string, error) {
	// Since we want a 64-character secret and each character is 8 bits,
	// we need to generate 32 bytes and then encode it using base64
	const byteLength = 32

	secretBytes := make([]byte, byteLength)
	_, err := rand.Read(secretBytes)
	if err != nil {
		return "", err
	}

	// Encoding the random bytes to base64
	secretBase64 := base64.RawURLEncoding.EncodeToString(secretBytes)
	return secretBase64, nil
}


func GenerateTLS(keyFile string, certFile string, keySize string) error {
	log.Println("Generating TLS keys ... " + keySize + " bits")

	privKey, err := rsa.GenerateKey(rand.Reader, stringToBits(keySize))
	if err != nil {
		return err
	}

	keyOut, err := os.OpenFile(keyFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer keyOut.Close()

	privBytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		return err
	}

	keyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}

	if err := pem.Encode(keyOut, keyBlock); err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Organization: []string{"gobox"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365), // one year

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	if err != nil {
		return err
	}

	certOut, err := os.OpenFile(certFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer certOut.Close()

	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}

	if err := pem.Encode(certOut, certBlock); err != nil {
		return err
	}

	log.Println(keyFile)
	log.Println(certFile)
	
	return nil
}

func stringToBits(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}