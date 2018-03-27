package user

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strconv"
	"strings"
)

const (
	passwordType  = "argon2"
	maxSaltSize   = 16
	argon2KeySize = 32
	argon2Time    = 4
	argon2Threads = 4
	argon2Memory  = 16 * 1024
)

func generateSalt() (s string, err error) {
	unencodedSalt := make([]byte, maxSaltSize)
	_, err = rand.Read(unencodedSalt)
	if err != nil {
		return s, err
	}
	return base64.StdEncoding.EncodeToString(unencodedSalt), nil
}

func HashPassword(p string) (h string, err error) {
	salt, err := generateSalt()

	if err != nil {
		return "", err
	}

	unEncodedHash := argon2.Key([]byte(p), []byte(salt), argon2Time, argon2Memory, argon2Threads, argon2KeySize)
	encodedHash := base64.StdEncoding.EncodeToString(unEncodedHash)

	hash := fmt.Sprintf("%s$%d$%d$%d$%d$%s$%s",
		passwordType, argon2Time, argon2Memory, argon2Threads, argon2KeySize, salt, encodedHash)

	return hash, nil
}

func CompareHashWithPassword(hash, password string) (bool, error) {
	if len(hash) == 0 || len(password) == 0 {
		return false, errors.New("arguments cannot be zero length")
	}
	hashParts := strings.Split(hash, "$")
	time, _ := strconv.Atoi(hashParts[1])
	memory, _ := strconv.Atoi(hashParts[2])
	threads, _ := strconv.Atoi(hashParts[3])
	keySize, _ := strconv.Atoi(hashParts[4])
	salt := []byte(hashParts[5])
	key, _ := base64.StdEncoding.DecodeString(hashParts[6])

	calculatedKey := argon2.Key([]byte(password), salt, uint32(time), uint32(memory), uint8(threads), uint32(keySize))

	if subtle.ConstantTimeCompare(key, calculatedKey) != 1 {
		return false, errors.New("password did not match")
	}

	return true, nil
}
