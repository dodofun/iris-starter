package jwt

import (
	"testing"
)

func TestCreateJWTToken(t *testing.T) {
	token, err := CreateToken(map[string]interface{}{"userId": "123456"}, 3600)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	value, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2In0.EcekJCyYcfUHWZXhw71rlrr8adnyqvqutB8cWAgm5O4")
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
}
