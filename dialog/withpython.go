package dialog

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type PythonScript struct {
	path string
	Order
}

func GeneralRunPython(para string) string {
	fmt.Println(para)
	args := strings.Split(para, " ")

	//fmt.Println(args)
	//cmd0 := exec.Command("cmd", "/C", "dir", "/B")
	//// 创建命令
	//fullCmd := strings.Join(cmd0.Args, " ")
	//fmt.Println(fullCmd)
	//output, err0 := cmd0.Output()
	//if err0 != nil {
	//	fmt.Println("命令执行失败:", err0)
	//	return fmt.Sprintf("%s", err0)
	//}
	//
	//// 输出命令的输出结果
	//fmt.Println(string(output))

	cmd := exec.Command("python3", args...)
	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("%s", string(output))

	} else {
		//println(string(output))
		return fmt.Sprintf("%s", string(output))
	}
	// 打印输出
}

func GeneralOrder(para string) string {
	name := para
	//outputFile := "output.mp3" // 默认输出文件名为 output.mp3
	// 构造 spotdl 命令
	cmd := exec.Command("python3", "-m", "spotdl", name, "--output", fmt.Sprintf("test/%s.{output-ext}", name))
	//cmd := exec.Command("spotdl", name)

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return fmt.Sprintf("%s", err)
	}

	// 输出提示信息
	fmt.Printf("Successfully downloaded track from %s\n", name)
	return name
}

func PyscriptToOderpre(path string) *Orderpre {
	{
		// 打开文件
		file, err := os.Open(path)
		if err != nil {
			panic(fmt.Sprintf("无法打开文件: %s", err))
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
			panic(fmt.Sprintf("读取文件时发生错误:%s", err))
		}

		// 输出保存的内容
		//fmt.Println("前10行内容为:")
		//fmt.Println(content)
		r := regexp.MustCompile("(?s)\"\"\"\n<<(.+?)>>\n\"\"\"")
		if r.MatchString(content) {
			result := r.FindAllStringSubmatch(content, -1)
			init_value := result[0][1]
			lines := strings.Split(init_value, "\n")

			// 初始化 Order 结构体
			orderpre := new(Orderpre)
			for _, line := range lines {
				parts := strings.Split(line, ":")
				if len(parts) == 2 {
					switch parts[0] {
					case "name":
						orderpre.name = parts[1]
					case "parse":
						orderpre.parse = parts[1]
					case "describe":
						orderpre.describe = parts[1]
					case "paracheck":
						orderpre.paracheckstr = parts[1]

					}
				}
			}
			orderpre.path = path
			fmt.Printf("%s :%s 初始化成功\n", orderpre.name, orderpre.path)
			return orderpre

		}
		fmt.Printf("%s 初始化失败\n", path)
		return nil

		// 遍历每行，提取字段值

	}

}
