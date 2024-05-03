package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	main2 "mywebot/dialog"
	"runtime"
)

func main() {
	//bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	//
	//// 注册消息处理函数
	////bot.MessageHandler = func(msg *openwechat.Message) {
	////	println(msg.FromUserName)
	////	if msg.IsText() && msg.Content == "ping" && msg.FromUserName == "珂爱336" {
	////		msg.ReplyText("pong")
	////	}
	////}
	//dip := openwechat.NewMessageMatchDispatcher()
	//dip.SetAsync(true)
	////dip.RegisterHandler(match, translate) /* /ping 汉字 转拼音 */
	////a := main2.Dialog{Initpath: "E:\\未完成\\mywebot\\learn\\a.dialog"}
	//a := main2.Dialog{Initpath: "../config/a.dialog"}
	//a.Init()
	//main2.Init(&a)
	//dip.OnText(a.Reply) /* 字典回复 */
	////位置
	//dip.RegisterHandler(
	//	func(message *openwechat.Message) bool {
	//		return message.IsLocation()
	//	},
	//	func(ctx *openwechat.MessageContext) {
	//		r := regexp.MustCompile("coord=(-?[0-9.]+),(-?[0-9.]+)")
	//		fi := r.FindStringSubmatch(ctx.Url)
	//		x, y := fi[1], fi[2]
	//		//ctx.ReplyText(fmt.Sprintf("经度:%s\n纬度:%s", x, y))
	//		//println(ctx.Content)
	//		locationname := regexp.MustCompile("(.+?)[:/]").FindStringSubmatch(ctx.Content)[1]
	//		a := fmt.Sprintf("../config/pyscrpits/weather2.py %s,%s,%s", locationname, y, x)
	//		fmt.Println(a)
	//		rpy := main2.GeneralRunPython(fmt.Sprintf("../config/pyscrpits/weather2.py %s,%s,%s", locationname, y, x))
	//		ctx.ReplyText(rpy)
	//	})
	//dip.OnRecalled(func(ctx *openwechat.MessageContext) {
	//	ctx.ReplyText("嘿嘿,看到了.学长,你也不想别人知道吧")
	//	img, _ := os.Open("../config/hwhw.gif")
	//	defer img.Close()
	//	if _, err := ctx.ReplyImage(img); err != nil {
	//		fmt.Println(err)
	//	}
	//})
	//
	////音乐
	//
	//dip.RegisterHandler(
	//	func(message *openwechat.Message) bool {
	//		msre := regexp.MustCompile("^music ?.+")
	//
	//		return msre.MatchString(message.Content)
	//
	//	},
	//	func(ctx *openwechat.MessageContext) {
	//		r := regexp.MustCompile("^music ?(.+)")
	//		fi := r.FindStringSubmatch(ctx.Content)
	//		name := fi[1]
	//		//ctx.ReplyText(name)
	//		path := fmt.Sprintf("./test/%s.mp3", name)
	//		file, err := os.Open(path)
	//		if err != nil {
	//			main2.GeneralOrder(name)
	//			file, _ = os.Open(path)
	//
	//		}
	//		ctx.ReplyFile(file)
	//		defer file.Close()
	//
	//		//fmt.Println(path)
	//		//ctx.ReplyFile()
	//
	//		//ctx.ReplyText(fmt.Sprintf("经度:%s\n纬度:%s", x, y))
	//	})
	////拍一拍
	//dip.RegisterHandler(
	//	func(message *openwechat.Message) bool {
	//		return message.IsPaiYiPai() && message.IsSendByFriend()
	//
	//	},
	//	func(ctx *openwechat.MessageContext) {
	//		//fmt.Println("payipa")
	//		reply := main2.GeneralRunPython("../config/pyscrpits/whenpad.py")
	//		ctx.ReplyText(reply)
	//		//fmt.Println(path)
	//		//ctx.ReplyFile()
	//
	//		//ctx.ReplyText(fmt.Sprintf("经度:%s\n纬度:%s", x, y))
	//	})
	////
	//bot.MessageHandler = dip.AsMessageHandler()
	//// 注册登陆二维码回调
	//bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	//reloadStorage := openwechat.NewFileHotReloadStorage("../config/storage.json")
	//
	//defer reloadStorage.Close()
	//
	//// 执行热登录
	////bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	//// 登陆
	//if err := bot.Login(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// 获取登陆的用户
	//self, err := bot.GetCurrentUser()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////println(self.Friends(true))
	////println(self.Groups(true))
	//
	//// 获取所有的好友
	//friends, err := self.Friends()
	//for _, v := range friends {
	//	println(v.UserName, v.ID(), v.Sex, v.City)
	//}
	//
	//// 获取所有的群组
	//groups, err := self.Groups()
	//fmt.Println(groups, err)
	//
	//// 阻塞主goroutine, 直到发生异常或者用户主动退出
	//bot.Block()
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
