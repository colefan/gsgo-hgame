package config

import (
	"fmt"
	"testing"
)

func TestServerConfig(t *testing.T) {
	if false == GetServerConfig().InitConfig() {
		t.Fatal()
	}
	fmt.Println(GetServerConfig().xmlconf)
}
