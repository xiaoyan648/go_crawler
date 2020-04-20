package fetcher

import (
	"testing"
)

func TestFetch(t *testing.T) {
	url := "http://www.zhenai.com/zhenghun";
	_, err := Fetch(url)
	if err != nil {
		t.Errorf("Fetch data error: %v",err)
	}
}
