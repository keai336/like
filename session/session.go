package session

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
)

func A(mc chan string, wait *sync.WaitGroup) {
	cmd := exec.Command("python", "E:\\未完成\\mywebot\\learn\\btime.py")
	in, _ := cmd.StdinPipe()
	//cmd.Stdout = os.Stdout
	out, _ := cmd.StdoutPipe()
	go func() {
		reader := bufio.NewReader(out)
		for {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				fmt.Println("读取输出数据时出错:", err)
				return
			}
			fmt.Println(strings.Replace(line, "\n", "", -1))

			//fmt.Println("分割")
			if err == io.EOF {
				break
			}
		}
	}()
	cmd.Start()
	go func() {
		for i := 0; i < 10; i++ {
			mes := <-mc
			if mes != "quit" && mes != "" {
				io.WriteString(in, fmt.Sprintf("%s\n", mes))
				//time.Sleep(time.Second)
			} else {
				wait.Done()
				break
			}
		}
		//io.WriteString(in, "exit()\n")
	}()
	cmd.Wait()
}
