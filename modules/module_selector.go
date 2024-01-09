package modules

import (
	"openwechat-with-ChatGPT/openwechat"
)

func RUN() (msg *openwechat.Message) {
	// 选择模块
	if msg.IsText() && msg.Content == "ping" {
		msg.ReplyText("pong")
	}
	return
}
