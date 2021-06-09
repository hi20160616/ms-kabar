package fetcher

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	as, err := load()
	if err != nil {
		t.Error(err)
	}
	for _, a := range as {
		fmt.Println(a.Title)
	}
}
