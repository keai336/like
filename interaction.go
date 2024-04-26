package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main1() {
	cmd := exec.Command("python", "-i", "btime.py")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	cmd.Start()
	time.Sleep(time.Second)
	go func() {

		// 创建一个读取器，用于逐行读取命令的标准输出

		for i := 0; i < 10; i++ {
			var inp string
			fmt.Scanln(&inp)
			if inp != "" {
				io.WriteString(in, fmt.Sprintf("%s\n", inp))

			}
			//fmt.Println("!")

			// 逐行读取命令的输出
			//time.Sleep(10 * time.Second)
		}
	}()
	go func() {
		reader := bufio.NewReader(out)
		for {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				fmt.Println("读取输出数据时出错:", err)
				return
			}
			fmt.Println(line)
			//fmt.Println("分割")
			if err == io.EOF {
				break
			}
		}
	}()
	cmd.Wait()
}
