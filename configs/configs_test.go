package configs

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	if err := get(); err != nil {
		t.Error(err)
	}
	fmt.Println(Data)
}
