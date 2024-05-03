package dialog

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var OrderDic = map[string]Order{}

func (order Order) AddtoDic(dialog *Dialog) {
	OrderDic[order.name] = order
	dialog.Diaglog[order.parse] = order.name
}

type Order struct {
	name       string `json:"name"`
	parse      string `json:"parse"`
	describe   string `json:"describe"`
	paracheck  func(para string) bool
	paramodify func(para string) string
	run        func(para string) string
}

type Orderpre struct {
	name         string `json:"name"`
	parse        string `json:"parse"`
	describe     string `json:"describe"`
	paracheckstr string `json:"paracheck"`
	path         string
}

func NewOrder() *Order {
	order := new(Order)
	//order.paracheck = func(para string) bool {
	//	return true
	//
	//}
	order.paramodify = generalmodify
	return order

}

func NewOrderFromPre(orderpre *Orderpre) *Order {
	order := NewOrder()

	if orderpre.paracheckstr != "" {
		paracheck := func(para string) bool {
			return regexp.MustCompile(orderpre.paracheckstr).MatchString(para)
		}
		order.paracheck = paracheck
	}
	modify := generalmodify
	switch orderpre.parse {
	case "位置":
		modify = locationmodify

	}
	paramodify := func(para string) string {
		replacedStr := modify(para)
		return fmt.Sprintf("%s ", orderpre.path) + fmt.Sprintf("%s", replacedStr)
	}
	order.name = orderpre.name
	order.parse = orderpre.parse
	order.describe = orderpre.describe
	order.paramodify = paramodify
	order.run = GeneralRunPython
	return order

}

func translate(para string) string {
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone
	strslice := pinyin.LazyPinyin(para, a)
	strfi := strings.Join(strslice, "-")
	return strfi
}

func generalmodify(para string) string {
	replacedStr := strings.Replace(para, " ", "₹", -1)
	replacedStr = strings.Replace(replacedStr, "/n", "ℳ", -1)
	return replacedStr
}

func locationmodify(para string) string {
	paras := strings.Split(para, "`")
	url := paras[1]
	content := paras[0]
	r := regexp.MustCompile("coord=(-?[0-9.]+),(-?[0-9.]+)")
	fi := r.FindStringSubmatch(url)
	x, y := fi[1], fi[2]
	locationname := regexp.MustCompile("(.+?)[:/]").FindStringSubmatch(content)[1]
	return fmt.Sprintf("%s,%s,%s", locationname, x, y)
}

func Init(dialog *Dialog) {
	Order{name: "转换",
		parse:    "/ping",
		describe: "汉字 -> 拼音",
		run:      translate,
		paracheck: func(para string) bool {
			for _, char := range para {
				if unicode.Is(unicode.Han, char) {
					return true
				}
			}
			return false
		},
	}.AddtoDic(dialog)

	Order{name: "时间",
		parse:    "/time",
		describe: "当前时间",
		run: func(para string) string {
			if para == "" {
				return time.Now().Format("2006-01-02 15:04:05")
			}
			return time.Now().Format(para)
		},
	}.AddtoDic(dialog)

	Order{name: "pythontest",
		parse:    "/pytest",
		describe: "执行Python 脚本",
		paramodify: func(para string) string {
			//return "E:\\未完成\\mywebot\\test\\hello.py " + para
			return "../config/pyscrpits/hello.py " + para

		},
		run: GeneralRunPython,
	}.AddtoDic(dialog)

	Order{name: "彩云翻译",
		parse:    "/translate",
		describe: "翻译",
		paramodify: func(para string) string {
			replacedStr := generalmodify(para)
			//fmt.Println(replacedStr)

			//return "F:\\micloud\\coding\\py\\pythonProject1\\翻译.py " + fmt.Sprintf("%s", replacedStr)
			return "../config/pyscrpits/翻译.py " + fmt.Sprintf("%s", replacedStr)

		},
		run: GeneralRunPython,
	}.AddtoDic(dialog)

	Order{name: "emotionai",
		parse:    "/emoa",
		describe: "情绪分析",
		paramodify: func(para string) string {
			replacedStr := generalmodify(para)
			//fmt.Println(replacedStr)

			//return "F:\\micloud\\coding\\py\\pythonProject1\\翻译.py " + fmt.Sprintf("%s", replacedStr)
			return "../config/pyscrpits/aiapi.py " + fmt.Sprintf("%s", replacedStr)

		},
		run: GeneralRunPython,
	}.AddtoDic(dialog)

	Order{name: "街溜子ai",
		parse:    "/ai",
		describe: "回应ai",
		paramodify: func(para string) string {
			replacedStr := generalmodify(para)
			//fmt.Println(replacedStr)

			//return "F:\\micloud\\coding\\py\\pythonProject1\\翻译.py " + fmt.Sprintf("%s", replacedStr)
			return "../config/pyscrpits/chatg.py " + fmt.Sprintf("%s", replacedStr)

		},
		run: GeneralRunPython,
	}.AddtoDic(dialog)

	Order{name: "天气",
		parse:    "/weather",
		describe: "天气预报",
		paramodify: func(para string) string {
			replacedStr := strings.Replace(para, " ", "₹", -1)
			replacedStr = strings.Replace(replacedStr, "/n", "ℳ", -1)
			//fmt.Println(replacedStr)

			//return "F:\\micloud\\coding\\py\\pythonProject1\\翻译.py " + fmt.Sprintf("%s", replacedStr)
			return "../config/pyscrpits/weather.py " + fmt.Sprintf("%s", replacedStr)

		},
		run: GeneralRunPython,
	}.AddtoDic(dialog)

}

func Init2(dialog *Dialog) {
	dirPath := "../config/pyscrpits/"

	// 读取目录内容
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
	}

	// 遍历目录下的所有文件和子目录
	for _, file := range files {
		path := dirPath + file.Name()
		newa := NewOrderFromPre(PyscriptToOderpre(path))
		newa.AddtoDic(dialog)

	}
	println(dialog.Diaglog)
}
