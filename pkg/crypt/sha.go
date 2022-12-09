package crypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Sha1 return the sha1 value (SHA-1 hash algorithm) of string
func Sha1(data string) string {
	s := sha1.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum([]byte("")))
}

// Sha256 return the sha256 value (SHA256 hash algorithm) of string
func Sha256(data string) string {
	s := sha256.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum([]byte("")))
}

// Sha512 return the sha512 value (SHA512 hash algorithm) of string
func Sha512(data string) string {
	s := sha512.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum([]byte("")))
}

// HmacSha1 return the hmac hash of string use sha1
func HmacSha1(data, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha256 return the hmac hash of string use sha256
func HmacSha256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha512 return the hmac hash of string use sha512
func HmacSha512(data, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}
