package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// Custom claims object
type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Tokes has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid session ID")
	}

	return nil
}

var key = []byte{}

func main() {

	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	pass := "12345678"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hashed password: ", string(hashedPass))

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}
	log.Println("Logged in!")
}

// -------- =^.^= ---- Hashing the password  ---- =^.^= --------
func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error whilr generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}

// -------- =^.^= ---- Hmac - Hash Message Authentication Code - Cryptography    ---- =^.^= --------
func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)

	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hasing the message: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checking while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}
