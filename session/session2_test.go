package session

import (
	"bufio"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"io"
	"mywebot/plus"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestS(t *testing.T) {
	cmd := exec.Command("python", "E:\\未完成\\mywebot\\learn\\test2.py")
	in, _ := cmd.StdinPipe()
	//cmd.Stdout = os.Stdout
	out, _ := cmd.StdoutPipe()
	cmd.Start()
	c := make(chan string)
	go func() {
		c <- "2"
		time.Sleep(time.Second)
		c <- "23"
		time.Sleep(time.Second)
		c <- "9"
		c <- "9"
		time.Sleep(time.Second)
		c <- "quit"
	}()
	reader := bufio.NewReader(out)
	for {
		line, err := reader.ReadString('§')
		if err != nil && err != io.EOF {
			fmt.Println("读取输出数据时出错:", err)
			return
		}
		//fmt.Println(strings.Replace(line, "\n", "", -1))
		fmt.Println(line)

		mes := <-c
		if mes != "quit" && mes != "" {
			io.WriteString(in, fmt.Sprintf("%s\n", mes))
			line, err := reader.ReadString('§')
			if err != nil && err != io.EOF {
				fmt.Println("读取输出数据时出错:", err)
				return
			}
			fmt.Println(line)
		} else {
			break
		}
	}
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err != nil && err != io.EOF {
	//		fmt.Println("读取输出数据时出错:", err)
	//		return
	//	}
	//	//fmt.Println(strings.Replace(line, "\n", "", -1))
	//	fmt.Println(line)
	//
	//	//fmt.Println("分割")
	//	if err == io.EOF {
	//		break
	//	}
	//}
	//cmd.Start()
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		mes := <-c
	//		if mes != "quit" && mes != "" {
	//			io.WriteString(in, fmt.Sprintf("%s\n", mes))
	//			//time.Sleep(time.Second)
	//		} else {
	//			break
	//		}
	//	}
	//	//io.WriteString(in, "exit()\n")
	//}()
	cmd.Process.Kill()
	cmd.Wait()
	t.Log("退出")

}

type OneSession struct {
	In      chan string
	Out     chan string
	Message *openwechat.Message
	Box     *map[string]*OneSession
}

func NewSession(message *openwechat.Message, box *map[string]*OneSession) *OneSession {
	session := new(OneSession)
	session.In = make(chan string)
	session.Out = make(chan string)
	session.Message = message
	session.Box = box
	return session
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

func (onesession *OneSession) Run() {
	cmd := exec.Command("python", "E:\\未完成\\mywebot\\learn\\test2.py")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	//cmd.Start()
	err := cmd.Start()
	if err != nil {
		onesession.Message.ReplyText(fmt.Sprintln(err))
		return
	}
	reader := bufio.NewReader(out)
	line, err := reader.ReadString('~')
	if err != nil && err != io.EOF {
		fmt.Println("读取输出数据时出错:", err)
		return
	}
	onesession.Message.ReplyText(strings.Replace(line, "~", "", 1))
	for {
		mes := <-onesession.In
		if mes != "quit" {
			io.WriteString(in, fmt.Sprintf("%s\n", mes))
			line, err := reader.ReadString('~')
			if err != nil && err != io.EOF {
				fmt.Println("读取输出数据时出错:", err)
				return
			}
			if strings.Contains(line, "结束") {
				//fmt.Println("结束")
				cmd.Process.Kill()
				delete(*onesession.Box, onesession.Message.FromUserName)
				onesession.Out <- strings.Replace(line, "~", "", 1)
				//onesession.Message.ReplyText(strings.Replace(line, "~", "", 1))
				return
			}
			onesession.Out <- strings.Replace(line, "~", "", 1)
		} else {
			break
		}
	}
	cmd.Process.Kill()
	delete(*onesession.Box, onesession.Message.FromUserName)
	onesession.Out <- "结束"
}

func TestOnbot(t *testing.T) {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	//bot.MessageHandler = func(msg *openwechat.Message) {
	//	println(msg.FromUserName)
	//	if msg.IsText() && msg.Content == "ping" && msg.FromUserName == "珂爱336" {
	//		msg.ReplyText("pong")
	//	}
	//}
	box := make(map[string]*OneSession)
	//ongoingsession := make(map[string]chan string)
	dip := openwechat.NewMessageMatchDispatcher()
	dip.SetAsync(true)
	dip.RegisterHandler(func(message *openwechat.Message) bool {
		num := runtime.NumGoroutine()
		fmt.Println("当前 goroutine 数量:", num)
		fmt.Println(box)
		if regexp.MustCompile("^/birth$").MatchString(message.Content) {
			if _, ok := box[message.FromUserName]; !ok {
				return true
			}
		}
		return false

	},
		func(ctx *openwechat.MessageContext) {
			box[ctx.FromUserName] = NewSession(ctx.Message, &box)
			box[ctx.FromUserName].Run()
		})
	dip.RegisterHandler(func(message *openwechat.Message) bool {
		_, ok := box[message.FromUserName]
		return ok
	},
		func(ctx *openwechat.MessageContext) {

			box[ctx.FromUserName].In <- ctx.Content
			plus.AutoReply(ctx, <-box[ctx.FromUserName].Out)
			//ctx.ReplyText(<-box[ctx.FromUserName].Out)
		})
	bot.MessageHandler = dip.AsMessageHandler()
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	reloadStorage := openwechat.NewFileHotReloadStorage("../config/storage.json")

	defer reloadStorage.Close()
	// 执行热登录
	bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	bot.Block()
}
