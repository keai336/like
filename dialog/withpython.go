package dialog

import (
	"fmt"
	"os/exec"
	"strings"
)

type PythonScript struct {
	path string
	Order
}

func GeneralRunPython(para string) string {
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
