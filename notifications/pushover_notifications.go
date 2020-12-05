package pushover_notification

import (
	"log"

	"github.com/gregdel/pushover"
)

func NotifyPushover(title string, message string) {

	app := pushover.New("apq8s9bb2eq9ps58ktvhfihaxj2kzg")
	// Create a new recipient
	recipient := pushover.NewRecipient("ux7w4sef54hcfe7ufqw6u92k5mmfs6")
	message_p := pushover.NewMessageWithTitle(message, title)
	response, err := app.SendMessage(message_p, recipient)
	if err != nil {
		log.Panic(err)
	}
	log.Println(response)

}
