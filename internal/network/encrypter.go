package network

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	MathRand "math/rand"
	"time"
)

// NewCipher builds a new instance of Cipher
func NewCipher(encryptKey string) *Cipher {
	return &Cipher{
		encryptKey: encryptKey,
	}
}

// Cipher encrypts and decrypts data
type Cipher struct {
	encryptKey string
}

// Encrypt encrypts the data given
func (e *Cipher) Encrypt(data []byte) []byte {
	block, err := aes.NewCipher([]byte(e.encryptKey))
	if err != nil {
		log.Fatalln(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalln(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalln(err.Error())
	}

	return gcm.Seal(nonce, nonce, data, nil)
}

// Decrypt decrypts the data given
func (e *Cipher) Decrypt(data []byte) []byte {
	key := []byte(e.encryptKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalln(err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return plaintext
}

// GenerateHash builds a md5 encryptKey
func GenerateHash() string {
	randomText := time.Now().String() + string(MathRand.Intn(100)*MathRand.Intn(100))
	hash := md5.New()
	hash.Write([]byte(randomText))
	return hex.EncodeToString(hash.Sum(nil))
}
