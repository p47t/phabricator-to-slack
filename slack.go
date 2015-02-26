package main

import (
	"net/http"
	"net/url"
)

type Slack struct {
	Token    string
	Username string
}

func (s *Slack) postMessage(channel string, text string) {
	http.PostForm("https://slack.com/api/chat.postMessage", url.Values{
		"token":    {s.Token},
		"username": {s.Username},
		"channel":  {channel},
		"text":     {text},
	})
}
