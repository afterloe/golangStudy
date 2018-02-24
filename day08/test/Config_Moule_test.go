package test

import (
	"testing"
	"../config"
	"fmt"
)

func Test_ConfigModule_readConfig(t *testing.T) {
	val := config.Get("name")
	fmt.Println(val)
	if "cynomys_go" != val {
		t.Error("get config fail.")
	}
}