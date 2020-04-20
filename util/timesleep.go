package util

import (
	"math/rand"
	"time"
)

func RandomTimeSleep(n int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(n)))
}
