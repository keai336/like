package learn

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type filepath string

func TestRead(t *testing.T) {
	var a filepath
	a = "a.dialog"
	//方案一:
	t.Log("方案一\n\n")
	fi := make([]byte, 1024)
	if contend, err := os.Open(string(a)); err != nil {
		t.Error("打开失败")
	} else {
		if n, err := contend.Read(fi); err != nil {
			t.Error("读失败")
		} else {
			t.Log("读了", n)
			t.Log(string(fi))
		}
	}
	//方案二
	if message, err := ioutil.ReadFile(string(a)); err != nil {
		t.Error("不行")
	} else {
		t.Log(string(message), "\n")
	}

	//方案三
	t.Log("方案三", "\n\n")
	file, err := os.Open(string(a))
	if err != nil {
		fmt.Printf("无法打开文件: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// 逐行读取并打印文件内容
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 检查Scanner是否在读取过程中出错
	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件出错: %v", err)
		return
	}
}
