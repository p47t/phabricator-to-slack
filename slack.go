package ph2slack

import (
	"net/http"
	"net/url"
)

// Slack provides API access to a Slack site.
type Slack struct {
	Token    string
	Username string
}

// PostMessage posts a message to specified channel
func (s *Slack) PostMessage(channel string, text string) {
	http.PostForm("https://slack.com/api/chat.postMessage", url.Values{
		"token":    {s.Token},
		"username": {s.Username},
		"channel":  {channel},
		"text":     {text},
	})
}
