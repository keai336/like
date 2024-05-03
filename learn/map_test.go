package learn

import (
	"fmt"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	a := make(map[string]string)
	a["1"] = "1"
	fmt.Println(a)
	go func(ma *map[string]string) {
		delete(*ma, "1")
	}(&a)
	time.Sleep(time.Second)
	fmt.Println(a)
}
