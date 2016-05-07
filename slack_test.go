package ph2slack

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestInvalidPostMessage(t *testing.T) {
    s := &Slack{
        Token: "invalid",
        Username: "bot",
    }
    err := s.PostMessage("test", "ph2slack testing")
    assert.EqualError(t, err, "invalid_auth")
}

func TestNoTokenPostMessage(t *testing.T) {
    s := &Slack{
        Token: "",
        Username: "bot",
    }
    err := s.PostMessage("test", "ph2slack testing")
    assert.EqualError(t, err, "not_authed")
}

func TestPostMessage(t *testing.T) {
    s := &Slack{
        Token: os.Getenv("SLACK_TEST_TOKEN"),
        Username: "bot",
    }
    err := s.PostMessage("test", "ph2slack testing")
    assert.NoError(t, err)
}