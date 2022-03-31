package gcp

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/rhdedgar/email-confim/channels"
	"github.com/rhdedgar/email-confim/models"
)

func PullMsgs(projectID, subID string) error {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		retMsg := string(msg.Data)
		fmt.Printf("Got message: %q\n", retMsg)
		channels.SetStringChan(models.MailChan, retMsg)
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}

	return nil
}
