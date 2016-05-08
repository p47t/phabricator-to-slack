package ph2slack

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhabricatorInvalidToken(t *testing.T) {
	ph := &Phabricator{
		Host:  "https://secure.phabricator.com",
		Token: "01234567890123456789012345678901",
	}
	_, err := ph.PhidQuery(os.Getenv("PHABRICATOR_TEST_PHID"))
	assert.Error(t, err)
}

func TestPhabricatorInvalidPhid(t *testing.T) {
	ph := &Phabricator{
		Host:  "https://secure.phabricator.com",
		Token: os.Getenv("PHABRICATOR_TEST_TOKEN"),
	}
	_, err := ph.PhidQuery("invalid")
	assert.Error(t, err)
}

func TestPhabricatorInvalidHost(t *testing.T) {
	ph := &Phabricator{
		Host:  "invalid",
		Token: os.Getenv("PHABRICATOR_TEST_TOKEN"),
	}
	_, err := ph.PhidQuery(os.Getenv("PHABRICATOR_TEST_PHID"))
	assert.Error(t, err)
}

func TestPhabricatorConnect(t *testing.T) {
	ph := &Phabricator{
		Host:  os.Getenv("PHABRICATOR_TEST_HOST"),
		Token: os.Getenv("PHABRICATOR_TEST_TOKEN"),
	}
	_, err := ph.PhidQuery(os.Getenv("PHABRICATOR_TEST_PHID"))
	assert.NoError(t, err)
}
