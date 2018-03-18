package ph2slack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidPostMessage(t *testing.T) {
	s := &Slack{
		Token:    "invalid",
		Username: "bot",
	}
	err := s.PostMessage("test", "ph2slack testing")
	assert.EqualError(t, err, "invalid_auth")
}

func TestNoTokenPostMessage(t *testing.T) {
	s := &Slack{
		Token:    "",
		Username: "bot",
	}
	err := s.PostMessage("test", "ph2slack testing")
	assert.EqualError(t, err, "not_authed")
}

func TestPostMessage(t *testing.T) {
	s := &Slack{
		Token:    os.Getenv("SLACK_TEST_TOKEN"),
		Username: "bot",
	}
	err := s.PostMessage("test", "ph2slack testing")
	assert.NoError(t, err)
}

func TestPostMessageContentError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `invalid JSON`)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	s := &Slack{
		APIHost:  ts.URL,
		Token:    os.Getenv("SLACK_TEST_TOKEN"),
		Username: "bot",
	}
	err := s.PostMessage("test", "ph2slack testing")
	assert.Error(t, err)
}

func TestPostMessageStatusError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	s := &Slack{
		APIHost:  ts.URL,
		Token:    os.Getenv("SLACK_TEST_TOKEN"),
		Username: "bot",
	}
	err := s.PostMessage("test", "ph2slack testing")
	assert.Error(t, err)
}
