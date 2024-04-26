package dialog

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestParseOrder(t *testing.T) {
	path := "aiapi.py"
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个 Scanner 来读取文件内容
	scanner := bufio.NewScanner(file)

	// 创建一个字符串变量来保存文件的前10行
	var content string

	// 逐行读取文件内容
	lineNumber := 0
	for scanner.Scan() {
		// 将当前行添加到字符串变量中
		content += scanner.Text() + "\n"

		// 每读取一行，计数器加1
		lineNumber++

		// 如果已经读取了10行，则退出循环
		if lineNumber >= 10 {
			break
		}
	}

	// 检查 Scan 函数的错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时发生错误:", err)
		return
	}

	// 输出保存的内容
	//fmt.Println("前10行内容为:")
	//fmt.Println(content)
	r := regexp.MustCompile("(?s)\"\"\"\n<<\n(.+?)>>\n\"\"\"")
	result := r.FindAllStringSubmatch(content, -1)
	init_value := result[0][1]
	lines := strings.Split(init_value, "\n")

	// 初始化 Order 结构体
	order := Order{}

	// 遍历每行，提取字段值
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			switch parts[0] {
			case "name":
				order.name = parts[1]
			case "parse":
				order.parse = parts[1]
			case "describe":
				order.describe = parts[1]
			}
		}
	}
	fmt.Println(order)

}
