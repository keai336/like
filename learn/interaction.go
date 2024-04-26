package learn

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
)

func Interact(t *testing.T) {
	// 要执行的命令
	//cmd := exec.Command("python3", "btime.py")
	//t.Log(cmd.Stdout)
	//cmd.Stdin = strings.NewReader("-r ")
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//cmd.Start()
	//cmd.Wait()\
	cmd := exec.Command("python", "btime.py")
	in, _ := cmd.StdinPipe()
	cmd.Stdout = os.Stdout
	cmd.Start()
	go func() {
		for i := 0; i < 10; i++ {
			//t.Log(cmd.Output())
			var userInput string
			//fmt.Println("请输入一些文本:")
			fmt.Scanln(&userInput)
			io.WriteString(in, userInput)
			//time.Sleep(time.Second)
		}
		//io.WriteString(in, "exit()\n")
	}()

	cmd.Wait()
}
