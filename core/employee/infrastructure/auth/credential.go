package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthCustomClaims struct {
	Name       string `json:"name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	IsAdmin    bool   `json:"is_admin"`
	CreatedAt  time.Time
	jwt.StandardClaims
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func EncryptPassword(data []byte) []byte {
	block, _ := aes.NewCipher([]byte(createHash("EQVJ7UM8xJNcfsaxs$aw3Es2Z@8ewegzxZ531C$^bhEoMq!%fe")))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	hashedPassword := gcm.Seal(nonce, nonce, data, nil)
	return hashedPassword
}

func DecryptPassword(data []byte) []byte {
	key := []byte(createHash("EQVJ7UM8xJNcfsaxs$aw3Es2Z@8ewegzxZ531C$^bhEoMq!%fe"))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plainPassword, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainPassword
}

/*
type Credentials struct {
	secretKey string
	issue     string
}

func GetJWTCredentials() *Credentials {
	return &Credentials{
		secretKey: base64.StdEncoding.EncodeToString([]byte("EQVJ7UM8xJNcfsaxs$aw3Es2Z@8ewegzxZ531C$^bhEoMq!%fe")),
		issue:     "Probien",
	}
}
*/
