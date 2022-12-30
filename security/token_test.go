package security

import (
	b64 "encoding/base64"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	lengths := []int{1, 1, 4, 5, 16, 32, 32, 64}
	for _, v := range lengths {
		Init(v)
		if len(SecretKey) != v {
			t.Errorf("Init() --> length of secret key wrong: %d Expected: %d", len(SecretKey), v)
		}
	}
}

func TestGenerateRandomString(t *testing.T) {
	lengths := []int{1, 1, 4, 5, 16, 32, 32, 64}
	for _, v := range lengths {
		if str, err := GenerateRandomString(v); err != nil {
			t.Errorf("GenerateRandomString() --> %v", err)
		} else {
			l := len([]rune(str))
			if l != v {
				t.Errorf("GenerateRandomString() --> Length Expected: %d Actual: %d", v, l)
			}
		}
	}

	str1, _ := GenerateRandomString(32)
	str2, _ := GenerateRandomString(32)
	if str1 == str2 {
		t.Errorf("GenerateRandomString() --> Expected different but both are equal: %v", str1)
	}

	lengths = []int{-5, -9, 0}
	for _, v := range lengths {
		str, err := GenerateRandomString(v)
		if err == nil {
			t.Error("GenerateRandomString() --> Should fail if length <=0.")
		}
		if str != "" {
			t.Errorf("GenerateRandomString() --> Should return empty string if length is <=0. Actual: %v", str)
		}
	}
}

func TestGenerateJWT(t *testing.T) {
	token, _ := GenerateJWT("test")
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		t.Errorf("GenerateJWT() --> Token should have 3 parts and has only %d\n", len(tokenParts))
	}

	// Test header
	a, _ := b64.StdEncoding.DecodeString(tokenParts[0])
	actualHeader := string(a)
	expectedHeader := `{"alg":"HS256","typ":"JWT"}`
	if actualHeader != expectedHeader {
		t.Errorf("GenerateJWT() --> Header actual: %s Header expected: %s", actualHeader, expectedHeader)
	}

	// Test Payload
	a, _ = b64.StdEncoding.DecodeString(tokenParts[1])
	actualPayload := string(a)
	expectedPayload := `{"name":"test","iss":"Chat App","exp":`
	if !strings.Contains(actualPayload, expectedPayload) {
		t.Errorf("GenerateJWT() --> Payload actual: %s Payload expected: %s", actualPayload, expectedPayload)
	}

	// TODO - test sign
}
