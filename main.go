package main

import (
	"github.com/rhdedgar/email-confim/mailer"
	"github.com/rhdedgar/email-confim/models"
	"github.com/rhdedgar/email-confirm/datastores"
	"github.com/rhdedgar/email-confirm/gcp"
)

func main() {
	go mailer.ConfirmEmail(models.MailChan)
	gcp.PullMsgs(datastores.AppSecrets.GCPProjectID, datastores.AppSecrets.UserEmailTopic)
}
