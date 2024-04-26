package test

import (
	"fmt"
	"os/exec"
	"testing"
)

// Order 包含命令描述和执行函数

// GenerateOrderDic 生成当前文件中的所有命令的字典

func TestPytontog(t *testing.T) {
	// Python 脚本路径
	scriptPath := "hello.py"

	// 参数列表
	args := []string{"Alice", "Bob", "Charlie"}

	// 创建命令
	cmd := exec.Command("python", append([]string{scriptPath}, args...)...)

	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing Python script:", err)
		return
	}

	// 打印输出
	fmt.Println("Python script output:")
	fmt.Println(string(output))
}
