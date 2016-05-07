package ph2slack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInvalidPhabricatorConnect(t *testing.T) {
    ph := &Phabricator{
        Host: "invalid",
        User: "invalid",
        Cert: "invalid",
    }
    err := ph.Connect()
    assert.Error(t, err)
}
