package initializers

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"

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
			"GIN_MODE": "debug",
			"WORKING_FOLDER": wd,
			"DB_PATH": wd + "/data/users.db", // ":memory:"
			"JWT_SECRET": "os.Environ() placeholder",
			"ACCESS_TIME": "3600*24*7", // 1 week
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
	if os.Getenv("GIN_MODE") == "debug" {
		secret = "TEST_kdølfjgp948wteøsrdkghpwe7thgui"
	}
	return secret
}

func GenerateSecret() (string, error) {
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