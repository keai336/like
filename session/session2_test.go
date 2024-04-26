package session

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestS(t *testing.T) {
	cmd := exec.Command("python", "E:\\未完成\\mywebot\\learn\\btime.py")
	in, _ := cmd.StdinPipe()
	//cmd.Stdout = os.Stdout
	out, _ := cmd.StdoutPipe()
	c := make(chan string)
	go func() {
		reader := bufio.NewReader(out)
		for {
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				fmt.Println("读取输出数据时出错:", err)
				return
			}
			//fmt.Println(strings.Replace(line, "\n", "", -1))
			fmt.Println(line)

			//fmt.Println("分割")
			if err == io.EOF {
				break
			}
		}
	}()
	cmd.Start()
	go func() {
		for i := 0; i < 10; i++ {
			mes := <-c
			if mes != "quit" && mes != "" {
				io.WriteString(in, fmt.Sprintf("%s\n", mes))
				//time.Sleep(time.Second)
			} else {
				break
			}
		}
		//io.WriteString(in, "exit()\n")
	}()
	c <- "1"
	c <- "2"
	c <- "q"
	c <- "quit"
	cmd.Wait()
}

func TestWriter(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)

	defer file.Close()

	// 获取bufio.Writer实例
	writer := bufio.NewWriter(file)

	// 写入字符串
	count, _ := writer.Write([]byte("hello go\n"))

	fmt.Printf("wrote %d bytes\n", count)

	// 写入字符串
	count, _ = writer.WriteString("hello world\n")

	fmt.Printf("wrote %d bytes\n", count)

	// 清空缓存 确保写入磁盘
	writer.Flush()
	files, _ := os.Open("test.txt")
	me := make([]byte, 1000)

	n, _ := files.Read(me)
	fmt.Println(string(me[:n]))
}
