/*
Copyright 2019 Doug Edgar.

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
package datastores

import (
	"fmt"
	"os"

	"github.com/rhdedgar/emailer/localfile"
	"github.com/rhdedgar/emailer/models"
)

var (
	AppSecrets = models.AppSecrets{}
	CharSet    = "UTF-8"
	Sender     = "no-reply@tacofreeze.com"
)

func init() {
	path := os.Getenv("SECRET_CONFIG_FILE")
	if path == "" {
		path = "/secrets/emailer_secrets.json"
	}
	localfile.GetJSON(path, &AppSecrets)

	if AppSecrets.GCPProjectID == "" {
		AppSecrets.GCPProjectID = os.Getenv("GCP_PROJECT_ID")
	}

	if AppSecrets.EmailTemplate == "" {
		AppSecrets.GCPProjectID = os.Getenv("EMAIL_TEMPLATE")
	}

	if AppSecrets.HTMLEmailTemplate == "" {
		AppSecrets.GCPProjectID = os.Getenv("HTML_EMAIL_TEMPLATE")
	}

	fmt.Println("printing AppSecrets")
	fmt.Printf("%+v\n", AppSecrets)
}
