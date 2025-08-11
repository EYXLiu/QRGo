package utils

import (
	"strings"
)

func isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}

func isAlphanumeric(s string) bool {
	allowed := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"
	for _, r := range s {
		if !strings.ContainsRune(allowed, r) {
			return false
		}
	}
	return len(s) > 0
}

func IsAllowed(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}
