package ph2slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SlackResult represents result returned from Slack API
type SlackResult struct {
	OK       bool   `json:"ok"`
	ErrorStr string `json:"error"`
}

func (r SlackResult) Error() string {
	return r.ErrorStr
}

// Slack provides API access to a Slack site.
type Slack struct {
	APIHost  string
	Token    string
	Username string
}

func (s *Slack) getAPIHost() string {
	if s.APIHost != "" {
		return s.APIHost
	}
	return "https://slack.com/api/"
}

// PostMessage posts a message to specified channel
func (s *Slack) PostMessage(channel string, text string) error {
	resp, err := http.PostForm(s.getAPIHost()+"chat.postMessage", url.Values{
		"token":    {s.Token},
		"username": {s.Username},
		"channel":  {channel},
		"text":     {text},
	})

	var result SlackResult
	resultJSON, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err = json.Unmarshal(resultJSON, &result); err != nil {
		return err
	}
	if !result.OK {
		return result
	}
	return nil
}
