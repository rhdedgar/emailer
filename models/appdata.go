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
	CookieSecret    string          `json:"cookie_secret"`
	GoogleAuthID    string          `json:"google_auth_id"`
	GoogleAuthKey   string          `json:"google_auth_key"`
	PsqlPassword    string          `json:"psql_password"`
	PsqlUser        string          `json:"psql_user"`
	PsqlServicePort string          `json:"psql_service_port"`
	PsqlDatabase    string          `json:"psql_database"`
	PsqlServiceHost string          `json:"psql_service_host"`
	Subject         string          `json:"subject"`
	CharSet         string          `json:"charset"`
	SESSender       string          `json:"ses_sender"`
	Recipient       string          `json:"recipient"`
	DynamoDBKeyID   string          `json:"dynamodb_key_id"`
	DynamoDBKey     string          `json:"dynamodb_key"`
	DynamoDBRegion  string          `json:"dynamodb_region"`
	DynamoDBTable   string          `json:"dynamodb_table"`
	SESRegion       string          `json:"ses_region"`
	SESKeyID        string          `json:"ses_key_id"`
	SESKey          string          `json:"ses_key"`
	AuthMap         map[string]bool `json:"auth_map"`
	SiteName        string          `json:"site_name"`
	GCPProjectID    string          `json:"gcp_project"`
	UserEmailTopic  string          `json:"user_email_topic"`
}
