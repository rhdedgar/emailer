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

package models

type AppSecrets struct {
	Subject               string `json:"subject"`
	CharSet               string `json:"charset"`
	SESSender             string `json:"ses_sender"`
	SESRegion             string `json:"ses_region"`
	SESKeyID              string `json:"ses_key_id"`
	SESKey                string `json:"ses_key"`
	SiteName              string `json:"site_name"`
	AppName               string `json:"app_name"`
	GCPProjectID          string `json:"gcp_project"`
	UserEmailTopic        string `json:"user_email_topic"`
	UserEmailSubscription string `json:"user_email_subscription"`
	CompanyInfo           string `json:"company_info"`
	EmailTemplate         string `json:"email_template"`
	HTMLEmailTemplate     string `json:"html_email_template"`
}
