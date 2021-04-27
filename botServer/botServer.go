package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
)

var (
	sertPath    = "/etc/letsencrypt/live/msbot.ddns.net/fullchain.pem"
	keyPath     = "/etc/letsencrypt/live/msbot.ddns.net/privkey.pem"
	accessKey   = "d804349249472af3567cff7f8ec5a179"
	accessToken = "PGxXex43ijl2VMKtAYY0CbeShmgSiLj/WLjSfZ0oiDOq7E+NyR8ty4dIhSLXE2VpfMOlI"
)

func main() {
	// init bot
	bot, err := linebot.New(accessKey, accessToken)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requestes from line platform
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Println(err)
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Println(err)
					}
				}
			}
		}
	})

	err = http.ListenAndServeTLS(":443", sertPath, keyPath, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// "d804349249472af3567cff7f8ec5a179",
// "PGxXex43ijl2VMKtAYY0CbeShmgSiLj/WLjSfZ0oiDOq7E+NyR8ty4dIhSLXE2VpfMOlI",
//w.Header().Set("Content-Type", "text/plain")
//w.Write([]byte("This is an example server.\n"))
// fmt.Fprintf(w, "This is an example server.\n")
// io.WriteString(w, "This is an example server.\n",))
