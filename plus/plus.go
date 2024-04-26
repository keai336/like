package plus

import (
	"fmt"
	oc "github.com/eatmoreapple/openwechat"
	"io"
	"os"
)

func Doloadfile(ctx *oc.MessageContext) error {
	file_name := ctx.FileName
	if res, err := ctx.GetFile(); err != nil {
		fmt.Println(err)
		return err

	} else {
		rc := res.Body

		// 创建一个新文件用于写入
		file, err := os.Create(fmt.Sprintf("../download/%smodified", file_name))
		if err != nil {
			fmt.Println("无法创建文件:", err)
			return err

		}
		defer file.Close()

		// 将 rc 的内容复制到文件中
		_, err = io.Copy(file, rc)
		if err != nil {
			fmt.Println("复制数据时出错:", err)
			return err
		}

		fmt.Println("数据已成功写入文件！")
	}
	return nil
}

func DoloadPicture(ctx *oc.MessageContext) error {
	if res, err := ctx.GetPicture(); err != nil {
		fmt.Println(err)
		return err

	} else {
		rc := res.Body

		// 创建一个新文件用于写入
		file, err := os.Create(fmt.Sprintf("../download/picture/%d.png", 1))
		if err != nil {
			fmt.Println("无法创建文件:", err)
			return err

		}
		defer file.Close()

		// 将 rc 的内容复制到文件中
		_, err = io.Copy(file, rc)
		if err != nil {
			fmt.Println("复制数据时出错:", err)
			return err
		}

		fmt.Println("数据已成功写入文件！")
	}
	return nil

}
