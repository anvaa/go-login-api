package initializers

import (
	"log"
	"crypto/rand"
	"encoding/base64"
	"github.com/joho/godotenv"
)

func LoadEnv(wd string) {
	err := godotenv.Load(wd + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func WriteEnv(wd string) {
	err := godotenv.Write(
		map[string]string{
			"PORT": "8090",
			"JWT_SECRET": GetSecret(),
			"WORKING_FOLDER": wd,
			"DB_PATH": wd + "/data/data.db", // ":memory:"
		},
		wd + "/.env",
	)
	if err != nil {
		log.Fatal("Error writing .env file")
	}
}


func GetSecret() string {
	secret, err := GenerateSecret()
	if err != nil {
		log.Fatal("Error generating JWT secret")
	}
	// uncomment next line to use a static secret when testing
	// secret = "TEST_kdølfjgp948wteøsrdkghpwe7thgui"
	return secret
}

func GenerateSecret() (string, error) {
	// Generating a random byte array of 32 bytes
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