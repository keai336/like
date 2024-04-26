package learn

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestErrout(t *testing.T) {
	arg := " spotdl kioku"
	// 执行的命令和参数
	args := strings.Split(arg, " ")
	cmd := exec.Command("python3", args...)

	// 执行命令并获取输出
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("命令执行错误:", string(combinedOutput))
		return
	}

	// 打印命令输出（包括标准输出和标准错误输出）
	fmt.Println("终端显示的所有输出:", string(combinedOutput))
}
