package test

import (
	main2 "mywebot/dialog"
	"testing"
)

func TestInit(t *testing.T) {
	a := main2.Dialog{Initpath: "E:\\未完成\\mywebot\\learn\\a.dialog"}
	a.Init()
	t.Log(a.Diaglog)
}
