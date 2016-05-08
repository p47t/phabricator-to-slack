package ph2slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PhObject represents attributes of a Phabricator object
type PhObject map[string]string

// PhObjects represents collection of PhObject
type PhObjects map[string]PhObject

// ConduitError represents an error happened when communicating with Phabricator
type ConduitError struct {
	Code string
	Info string
}

func (c ConduitError) Error() string {
	return fmt.Sprintf("%s %s", c.Code, c.Info)
}

// Phabricator provides API access to a Phabricator site specified by Host, User, and Cert.
type Phabricator struct {
	Host  string
	Token string
}

// PhidQuery queries attributes of a Phabricator object by its ID
func (p *Phabricator) PhidQuery(phid string) (PhObject, error) {
	conduitJSON, _ := json.Marshal(map[string]interface{}{
		"token": p.Token,
	})

	resp, err := http.PostForm(p.Host+"/api/phid.query", url.Values{
		"params[phids]":       {fmt.Sprintf(`["%s"]`, phid)},
		"params[__conduit__]": {string(conduitJSON)},
		"output":              {"json"},
	})
	if err != nil {
		return nil, err
	}

	result, err := p.resultOrError(resp)
	if err != nil {
		return nil, err
	}

	for _, obj := range result {
		return obj, nil // expect only one
	}
	return nil, fmt.Errorf("Empty result")
}

func (p *Phabricator) resultOrError(resp *http.Response) (PhObjects, error) {
	var result struct {
		ErrorCode string              `json:"error_code"`
		ErrorInfo string              `json:"error_info"`
		Result    map[string]PhObject `json:"result"`
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ErrorCode != "" {
		return nil, ConduitError{Code: result.ErrorCode, Info: result.ErrorInfo}
	}
	return result.Result, nil
}
