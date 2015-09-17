package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type PhObject map[string]string

type ConduitError struct {
	Code string
	Info string
}

func (c ConduitError) Error() string {
	return fmt.Sprintf("%s %s", c.Code, c.Info)
}

type Phabricator struct {
	Host string
	User string
	Cert string

	conduitJson string
}

func (p *Phabricator) Connect() error {
	// Prepare connect parameters
	token, signature := p.genAuthTokenAndSignature()
	paramsJson, _ := json.Marshal(map[string]interface{}{
		"client":        "Bot",
		"clientVersion": 0,
		"user":          p.User,
		"host":          p.Host,
		"authToken":     token,
		"authSignature": signature,
	})

	resp, err := http.PostForm(p.Host+"/api/conduit.connect", url.Values{
		"params":      {string(paramsJson)},
		"output":      {"json"},
		"__conduit__": {"true"},
	})
	if err != nil {
		return err
	}

	var result struct {
		ErrorCode string `json:"error_code"`
		ErrorInfo string `json:"error_info"`
		Result    struct {
			ConnectionID float64 `json:"connectionID"`
			SessionKey   string  `json:"sessionKey"`
			UserPHID     string  `json:"userPHID"`
		} `json:"result"`
	}
	resultJson, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(resultJson, &result); err != nil {
		return err
	}
	if result.ErrorCode != "" {
		return ConduitError{Code: result.ErrorCode, Info: result.ErrorInfo}
	}

	conduitJson, _ := json.Marshal(map[string]interface{}{
		"sessionKey":   result.Result.SessionKey,
		"connectionID": result.Result.ConnectionID,
	})

	// Keep the conduit for later API calls
	p.conduitJson = string(conduitJson)
	return nil
}

func (p *Phabricator) genAuthTokenAndSignature() (int64, string) {
	token := time.Now().Unix()
	sum := sha1.Sum([]byte(fmt.Sprintf("%d%s", token, p.Cert)))
	return token, hex.EncodeToString(sum[:])
}

func (p *Phabricator) PhidQuery(phid string) (PhObject, error) {
	resp, err := http.PostForm(p.Host+"/api/phid.query", url.Values{
		"params[phids]":       {fmt.Sprintf(`["%s"]`, phid)},
		"params[__conduit__]": {p.conduitJson},
		"output":              {"json"},
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		ErrorCode string              `json:"error_code"`
		ErrorInfo string              `json:"error_info"`
		Result    map[string]PhObject `json:"result"`
	}
	resultJson, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(resultJson, &result); err != nil {
		return nil, err
	}
	if result.ErrorCode != "" {
		return nil, ConduitError{Code: result.ErrorCode, Info: result.ErrorInfo}
	}

	for _, obj := range result.Result {
		return obj, nil // expect only one
	}
	return nil, fmt.Errorf("Empty result from %s", resultJson)
}
