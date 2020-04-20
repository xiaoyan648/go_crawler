package models

import (
	"github.com/astaxie/goredis"
	"testing"
)

func TestRedisConnect(t *testing.T) {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	_, err := client.Ping()
	if err != nil {
		t.Error(err)
	}

}
