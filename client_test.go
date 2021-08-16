package oppo

import (
	"fmt"
	"strconv"
	"testing"
	"time"
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

}

type TestParseParamsToSortQeuryStruct struct {
	Name   string `json:"name"`
	Age    int    `json:"age,string"`
	IsMale bool   `json:"is_male,string"`
	Action string `json:"action,omitempty"`
}

func TestParseParamsToSortQeury(t *testing.T) {
	params := &TestParseParamsToSortQeuryStruct{
		Name:   "tom",
		Age:    18,
		IsMale: false,
		Action: "no",
	}
	now := strconv.FormatInt(time.Now().Unix(), 10)
	token := "fawefljafljeaw"
	secret := "fawfsafawefsa"

	result, err := ParseParamsToSortQeury(params, token, now, secret)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)

}
