package common

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func GenerateTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func IsEmptyString(s string) bool {
	return s == ""
}

func ContainsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
