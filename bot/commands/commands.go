package commands

import (
	"bot/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// contains method
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


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

	if res.Blacklisted == true {
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

func Status(b *gotgbot.Bot , ext *ext.Context) error {
	caption_text := fmt.Sprintf(`
	<b>Welcome Master %v</b>

	<i>The Vampire Of BlueMoon</i>
	
	<b>Status</b> : Clan Of BlueMoon
	
	<b>Use me to slay the CrimsonMoon</b>
` , ext.Message.From.FirstName)

	devs := []int{5014715207, 825664681, 1091139479, 2076788242, 1719660492, 1700123830, 1707848873, 2079472115, 2107137268, 870471128, 1742378072, 2108196300, 1802324609, 645739169, 1483971607, 1757774449, 1306525471, 769830161, 5086015489, 5244167973}
	if contains(devs , int(ext.Message.From.Id)) {
		b.SendPhoto(ext.Message.Chat.Id , "https://telegra.ph/file/6369995ee354ab21c6f06.jpg" , &gotgbot.SendPhotoOpts{
			Caption: caption_text,
			ParseMode: "html",
		})
	} else {
		b.SendMessage(ext.Message.Chat.Id , fmt.Sprintf("<b>Welcome %v</b>\n<b>Status</b> : %v" , ext.Message.From.FirstName , "Human") , &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
	}

}