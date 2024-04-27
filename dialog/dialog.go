package dialog

import (
	"bufio"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"os"
	"regexp"
)

type filepath string
type Dialog struct {
	Initpath filepath
	Diaglog  map[string]string
}

func (diaglog *Dialog) Init() {
	diaglog.Diaglog = make(map[string]string)
	file, err := os.Open(string(diaglog.Initpath))
	if err != nil {
		fmt.Printf("无法打开文件: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// 逐行读取并打印文件内容
	re := "([^:]+):(.+)"
	recompiled := regexp.MustCompile(re)
	for scanner.Scan() {
		if ok := recompiled.MatchString(scanner.Text()); ok {
			matchs := recompiled.FindStringSubmatch(scanner.Text())
			fmt.Println(matchs)
			if len(matchs) == 3 {
				diaglog.Diaglog[matchs[1]] = matchs[2]

			}

		}
	}
	fmt.Println(diaglog.Diaglog)

	// 检查Scanner是否在读取过程中出错
	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件出错: %v", err)
		return
	}
}
func (diaglog *Dialog) isorder(message *openwechat.MessageContext) (string, string, bool) {
	recompiled := regexp.MustCompile("^(/[a-z]+) ?((?s).*)$")
	if !recompiled.MatchString(message.Content) {
		return "", "", false
	} else {
		contend := recompiled.FindStringSubmatch(message.Content)
		name := contend[1]
		para := contend[2]
		return name, para, true
	}
}

func (diaglog *Dialog) Reply(message *openwechat.MessageContext) {
	//我靠写的啥啊看不懂了
	if name, para, ok := diaglog.isorder(message); !ok {
		if reply, ok := diaglog.Diaglog[message.Content]; ok {
			message.ReplyText(reply)
		}
	} else {
		if reply, ok := diaglog.Diaglog[name]; !ok {
			//fmt.Println("不ok")
			message.ReplyText("wrong order type /menu to get help hwhw")
		} else {
			if check := OrderDic[reply].paracheck; check == nil {
				//fmt.Println(check("asdfa"))
				modify := OrderDic[reply].paramodify
				para = modify(para)
				ctx := OrderDic[reply].run(para)
				message.ReplyText(ctx)
				fmt.Println("有修无检")
				//fmt.Println(para, "modified")
			} else {
				if check(para) {
					//message.ReplyText(fmt.Sprintf("执行%s,参数为%s", reply, para))
					ctx := OrderDic[reply].run(OrderDic[reply].paramodify(para))
					message.ReplyText(ctx)

				} else {
					message.ReplyText(fmt.Sprintf("order %s gets wrong parameters,type /help %s to get help hwhw", OrderDic[reply].parse, OrderDic[reply].parse))
				}
			}
		}

	}
}
