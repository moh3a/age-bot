/*
SIMPLE SLACK FILE UPLOAD BOT
*/

package upload

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func Upload(files []string) {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}

	for i := 0; i < len(files); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     files[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivateDownload)
	}
}
