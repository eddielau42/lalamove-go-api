package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"regexp"
)

// Signature	生成数字签名
func Signature(secret, message string) (sign string) {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(message))
	sign = hex.EncodeToString(hash.Sum(nil))
	return
}

// UniqueID	生成唯一标识
func UniqueID() string {
	b := make([]byte,48)
    if _,err := io.ReadFull(rand.Reader,b); err != nil {
        return ""
    }
	
	h := md5.New()
    h.Write([]byte(base64.URLEncoding.EncodeToString(b)))
    return hex.EncodeToString(h.Sum(nil))
}

// CheckPhone 校验手机号
func CheckPhone(phone string) bool {
	// reg: 
	if matched, err := regexp.MatchString(`^\+[1-9]\d{1,14}$`, phone); err == nil {
		return matched
	}
	return false
}
