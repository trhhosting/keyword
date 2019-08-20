package config

import (
	"github.com/jinzhu/configor"
)

/*
Sid	A 34 character string that uniquely identifies this API Key. You will use this as the basic-auth user when authenticating to the API.
FriendlyName	A descriptive string for this resource, chosen by your application, up to 64 characters long.
Secret	The secret your application uses to sign Access Tokens and to authenticate to the REST API (you will use this as the basic-auth password).
Note that for security reasons, this field is ONLY returned when the API Key is first created.
Uri	The URI for this resource, relative to https://api.twilio.com
DateCreated	The date-time this API Key was created, given as a RFC 2822 Timestamp.
DateUpdated	The date-time this API Key was most recently updated, given as a RFC 2822 Timestamp.
*/
var TLMConfig = struct {
	Cockroach struct {
		Name        string
		Adapter     string
		Host        string
		Port        string
		User        string
		Password    string
		SSLMode     string
		SSLCert     string
		SSLKey      string
		SSLRootCert string
		Root        string
		RootCert    string
		RootKey     string
	}
}{}

func CockroachDB() {
	configor.Load(&TLMConfig,  "settings/cockroach.yml")
}
