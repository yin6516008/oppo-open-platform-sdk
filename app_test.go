package oppo

import (
	"fmt"
	"testing"
)

func TestGetAppDetail(t *testing.T) {
	client, err := NewOppoClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetAppDetailParams{
		PkgName: "com.gengcon.android.jxc",
	}

	appDetail, _, err := client.GetAppDetail(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(appDetail)
}
