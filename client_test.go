package oppo

import (
	"fmt"
	"net/url"
	"testing"
)

func TestRefreshToken(t *testing.T) {
	client, err := NewOppoClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	err = refreshToken(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(client.token)

}

func TestUrlValues(t *testing.T) {
	params := map[string]string{
		"Name": "yin",
		"Age":  "ming",
	}
	p := url.Values{}

	for k, v := range params {
		p.Set(k, v)
	}
	fmt.Println(p.Encode())

}
