package plus

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"os"
	"strings"
)

func startsWithText(s string) bool {
	return strings.HasPrefix(s, "text:")

}

func startsWithFile(s string) bool {
	return strings.HasPrefix(s, "file:")
}

func startsWithVideo(s string) bool {
	return strings.HasPrefix(s, "vide:")
}

func startsWithImage(s string) bool {
	return strings.HasPrefix(s, "imag:")
}

func ProtocolReplyFile(ctx *openwechat.MessageContext, message string, mark string) {
	path := message[5:]
	path = strings.TrimRight(path, "\n\r\t")
	file, err := os.Open(path)
	if err != nil {
		ctx.ReplyText(fmt.Sprintln(err))
		return
	}
	defer file.Close()
	switch mark {
	case "file":
		ctx.ReplyFile(file)
	case "video":
		ctx.ReplyVideo(file)
	case "image":
		ctx.ReplyImage(file)
	}
}

func ProtocolReplyImage(ctx *openwechat.MessageContext, message string) {
	path := message[5:]
	img, err := os.Open(path)
	if err != nil {
		ctx.ReplyText(fmt.Sprintln(err))
		return
	}
	defer img.Close()
	ctx.ReplyImage(img)
}

func AutoReply(ctx *openwechat.MessageContext, message string) {
	messages := strings.Split(message, "~")
	for _, v := range messages {
		switch {
		case startsWithText(v):
			ctx.ReplyText(v[5:])
		case startsWithImage(v):
			ProtocolReplyFile(ctx, v, "image")
		case startsWithVideo(v):
			ProtocolReplyFile(ctx, v, "video")
		case startsWithFile(v):
			ProtocolReplyFile(ctx, v, "file")
		default:
			ctx.ReplyText(v)
		}

	}

}
