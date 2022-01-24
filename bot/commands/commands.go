package commands

import (
	"bot/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ext *ext.Context) error {
	b.SendMessage(ext.Message.Chat.Id, "Hello There\nI'mSylviorus Bot", &gotgbot.SendMessageOpts{
		ParseMode: "markdown",
	})

	return nil
}

func CheckCmd(b *gotgbot.Bot, ext *ext.Context) error {
	x := ext.Args()[1]
	user, err := strconv.Atoi(x)

	if err != nil {
		ext.Message.Reply(b, err.Error(), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})

		return nil
	}

	res, _ := utils.Check(user)

	if res.Blacklisted != false {
		var text string = fmt.Sprintf("<b>USER</b> : %v\n<b>REASON</b> : %v\n<b>ENFORCER</b> : %v\n<b>Message</b> : %v", res.User, res.Reason, res.Enforcer, res.Message)
		b.SendMessage(ext.Message.Chat.Id, text, &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
	} else {
		b.SendMessage(ext.Message.Chat.Id, "This User Is Not Blacklisted", &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
	}

	return nil

}

func BanCmd(b *gotgbot.Bot, ext *ext.Context) error {
	x := ext.Args()[1]
	user, _ := strconv.Atoi(x)
	str := strings.Replace(ext.Message.Text, ext.Args()[0], "", -1)
	reason := strings.Replace(str, x, "", -1)
	utils.Ban(user, int(ext.Message.From.Id), reason, "")

	ext.Message.Reply(b, "User Banned", &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	return nil
}

func UnBanCmd(b *gotgbot.Bot, ext *ext.Context) error {
	x := ext.Args()[1]
	user, _ := strconv.Atoi(x)
	utils.Unban(user)

	return nil
}
