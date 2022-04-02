package main

import (
	"fmt"

	"github.com/rhdedgar/email-confirm/datastores"
	"github.com/rhdedgar/email-confirm/gcp"
	"github.com/rhdedgar/email-confirm/mailer"
	"github.com/rhdedgar/email-confirm/models"
)

func main() {
	fmt.Println("Starting emailer v0.0.1")

	go mailer.ConfirmEmail(models.MailChan)

	err := gcp.PullMsgs(datastores.AppSecrets.GCPProjectID, datastores.AppSecrets.UserEmailSubscription)
	if err != nil {
		fmt.Printf("Error returned from PullMsgs: %v\n", err)
	}
}
