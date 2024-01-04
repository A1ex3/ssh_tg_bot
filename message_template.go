package main

import (
	"fmt"
	"time"
)

type MessageTemplate struct{}

type template struct {
	DateTime time.Time
	HostName string
	UserName string
	RemoteIp string
}

func (messageTemplate *MessageTemplate) Create(
	host string,
	username string,
	dateTime time.Time,
	remoteIp string,
) string {
	templ := &template{
		UserName: username,
		DateTime: dateTime,
		RemoteIp: remoteIp,
		HostName: host,
	}

	var formatedDateTime string = templ.DateTime.UTC().Format(time.RFC3339)

	return fmt.Sprintf(
		"Host: %s\nUser Name: %s\nDate Time UTC: %s\nIp: %s",
		templ.HostName,
		templ.UserName,
		formatedDateTime,
		templ.RemoteIp,
	)
}