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
	passwordType  = "Argon2"
	maxSaltSize   = 16
	argon2KeySize = 32
	argon2Time    = 4
	argon2Threads = 4
	argon2Memory  = 16 * 1024
)

type PwdAlgo interface {
	Hash(p string) (h string, err error)
	Compare(hash, password string) (bool, error)
}

type IArgon2 interface {
	Key(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32) []byte
	GenerateSalt(size int) (s string, err error)
	EncodeToString([]byte) string
}

type Argon2 struct {
}

type PwdArgon2 struct {
	Crypto IArgon2
}

func NewPwdAlgo(algo string) (pa PwdAlgo, err error) {
	if strings.EqualFold("argon2", algo) {
		return PwdArgon2{Argon2{}}, nil
	}

	return nil, errors.New(fmt.Sprintf("unsupported password algorithm %s", algo))
}

func (pwd PwdArgon2) Hash(p string) (h string, err error) {
	salt, err := pwd.Crypto.GenerateSalt(maxSaltSize)

	if err != nil {
		return "", err
	}

	unEncodedHash := pwd.Crypto.Key([]byte(p), []byte(salt), argon2Time, argon2Memory, argon2Threads, argon2KeySize)
	encodedHash := pwd.Crypto.EncodeToString(unEncodedHash)

	hash := fmt.Sprintf("%s$%d$%d$%d$%d$%s$%s",
		passwordType, argon2Time, argon2Memory, argon2Threads, argon2KeySize, salt, encodedHash)

	return hash, nil
}

func (pwd PwdArgon2) Compare(hash, password string) (bool, error) {
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

	calculatedKey := pwd.Crypto.Key([]byte(password), salt, uint32(time), uint32(memory), uint8(threads), uint32(keySize))

	if subtle.ConstantTimeCompare(key, calculatedKey) != 1 {
		return false, errors.New("password did not match")
	}

	return true, nil
}

func (c Argon2) GenerateSalt(size int) (s string, err error) {
	unencodedSalt := make([]byte, size)
	_, err = rand.Read(unencodedSalt)
	if err != nil {
		return s, err
	}
	return base64.StdEncoding.EncodeToString(unencodedSalt), nil
}

func (c Argon2) Key(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32) []byte {
	return argon2.Key(password, salt, time, memory, threads, keyLen)
}

func (c Argon2) EncodeToString(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}
