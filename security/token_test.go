package security

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	lengths := []int{1, 1, 4, 5, 16, 32, 32, 64}
	for _, v := range lengths {
		if str, err := GenerateRandomString(v); err != nil {
			t.Errorf("GenerateRandomString() --> %v", err)
		} else {
			fmt.Println(str)
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
