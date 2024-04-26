package test

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"regexp"
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

	// 注册消息处理函数
	//bot.MessageHandler = func(msg *openwechat.Message) {
	//	println(msg.FromUserName)
	//	if msg.IsText() && msg.Content == "ping" && msg.FromUserName == "珂爱336" {
	//		msg.ReplyText("pong")
	//	}
	//}
	//ongoingsession := make(map[string]chan string)
	dip := openwechat.NewMessageMatchDispatcher()

	dip.OnText(func(ctx *openwechat.MessageContext) {
		fmt.Println(ctx.FromUserName)
		//a, _ := ctx.Sender()
		//openwechat.Sech
		//fmt.Println()
	})
	bot.MessageHandler = dip.AsMessageHandler()
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	defer reloadStorage.Close()

	// 执行热登录
	bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	// 登陆
	//if err := bot.Login(); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	//println(self.Friends(true))
	//println(self.Groups(true))

	// 获取所有的好友
	friends, err := self.Friends()
	friendss := friends.SearchByUserName(10, "@a3315313dca96f30df42b0436efe2a03aa17dc2bf304413a122f2fad729e016a")
	self.SendTextToFriends("hello", 0, friendss...)
	//for _, v := range friends {
	//	if v.UserName == "@a3315313dca96f30df42b0436efe2a03aa17dc2bf304413a122f2fad729e016a" {
	//		println(v.UserName, v.ID(), v.Sex, v.City)
	//		v.SendText("nihc")
	//
	//	}
	//}

	// 获取所有的群组
	groups, err := self.Groups()
	t.Log(groups, err)
	t.Log("aaa")

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
