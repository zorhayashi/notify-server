package supapp

import (
	"net/http"
	"strings"

	"github.com/zorhayashi/notify-server/config"
	"github.com/zorhayashi/notify-server/util"
)

//DiscordPost discord webhooks
func DiscordPost(msg string) {
	resp, err := http.Post(config.Global.Discord.WebhookLink,
		"application/x-www-form-urlencoded",
		strings.NewReader("content="+msg))

	if err != nil {
		// handle error
		util.Error(err.Error())
	}
	util.Success("[Discord].Post " + msg)
	defer resp.Body.Close()
}
