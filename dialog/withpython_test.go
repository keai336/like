package dialog

import "testing"

//func Wrap(order) func(para string) string {
//	return func(para string) string {
//
//	}
//}

func TestE(t *testing.T) {
	a := PythonScript{
		path:  "E:\\未完成\\mywebot\\test\\hello.py",
		Order: Order{name: "hello", run: GeneralRunPython},
	}
	t.Log(a.run(a.path + " hhh aaa"))
}
