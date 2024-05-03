package test

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	main2 "mywebot/dialog"
	"regexp"
	"runtime"
	"testing"
)

// extractPingContent 函数用于提取字符串中以/ping开头后的内容
func extractPingContent(input string) string {
	// 定义正则表达式模式，使用捕获组提取/ping后的内容
	pattern := "^/ping(.*)"

	// 编译正则表达式
	reg := regexp.MustCompile(pattern)

	// 查找匹配的子串
	matches := reg.FindStringSubmatch(input)

	// 如果有匹配到内容
	if len(matches) > 1 {
		return matches[1] // 返回捕获组中的内容
	}
	return "" // 没有匹配到内容，返回空字符串
}

func main() {
	// 测试字符串
	testStrings := []string{"/ping", "/ping123", "pong/ping", "hello"}

	// 遍历测试字符串，提取/ping后的内容
	for _, str := range testStrings {
		content := extractPingContent(str)
		if content != "" {
			fmt.Printf("%s 的内容为：%s\n", str, content)
		} else {
			fmt.Printf("%s 不匹配 /ping 开头\n", str)
		}
	}
}

func match(message *openwechat.Message) bool {

	pattern := "^/ping(.*)"

	// 编译正则表达式
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(message.Content) {
		matches := reg.FindStringSubmatch(message.Content)
		message.Content = matches[0]
		return true
	} else {
		return false
	}
}

func TestOne(t *testing.T) {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	dip := openwechat.NewMessageMatchDispatcher()
	dip.SetAsync(true)
	a := main2.Dialog{Initpath: "../config/a.dialog"}
	a.Init()
	main2.Init2(&a)
	dip.RegisterHandler(func(message *openwechat.Message) bool {
		num := runtime.NumGoroutine()
		fmt.Println("当前 goroutine 数量:", num)
		return true
	},
		a.Reply)
	fmt.Println(a.Diaglog)
	bot.MessageHandler = dip.AsMessageHandler()
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	bot.Block()
}
