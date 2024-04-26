package learn

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"testing"
	"time"
)

//	func TestInteraction(t *testing.T) {
//		cmd := exec.Command("python", "-i")
//		in, _ := cmd.StdinPipe()
//		//cmd.Stdout = os.Stdout
//		out, _ := cmd.StdoutPipe()
//		cmd.Start()
//		go func() {
//			for i := 0; i < 10; i++ {
//				var inp string
//				inp = "print(1+1)"
//				io.WriteString(in, fmt.Sprintf("%s\n", inp))
//				//time.Sleep(time.Second)
//				fmt.Println("!")
//				time.Sleep(10 * time.Second)
//				contend, err := ioutil.ReadAll(out)
//				fmt.Println(1)
//				if err != nil {
//					fmt.Println(err)
//				}
//				fmt.Println(string(contend))
//			}
//			//io.WriteString(in, "exit()\n")
//		}()
//
//		cmd.Wait()
//	}
func TestInteraction(t *testing.T) {
	cmd := exec.Command("python", "-i", "btime.py")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	cmd.Start()

	go func() {

		// 创建一个读取器，用于逐行读取命令的标准输出

		for i := 0; i < 10; i++ {
			var inp string
			//fmt.Scanln(&inp)
			//inp = "1"
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
			time.Sleep(time.Second)
		}
	}()
	cmd.Wait()
}
