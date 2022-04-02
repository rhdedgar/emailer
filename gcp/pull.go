/*
Copyright 2022 Doug Edgar.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gcp

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/rhdedgar/emailer/channels"
	"github.com/rhdedgar/emailer/models"
)

func PullMsgs(projectID, subID string) error {
	fmt.Println("Starting PullMsgs")
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
