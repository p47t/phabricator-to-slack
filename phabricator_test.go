package ph2slack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestPhabricatorQuery(t *testing.T) {
	ph := &Phabricator{
		Host:  os.Getenv("PHABRICATOR_TEST_HOST"),
		Token: os.Getenv("PHABRICATOR_TEST_TOKEN"),
	}
	_, err := ph.PhidQuery(os.Getenv("PHABRICATOR_TEST_PHID"))
	assert.NoError(t, err)
}

func TestPhabricatorQueryError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error_code": "500", "error_info": "something wrong"}`)
	}))
	defer ts.Close()

	ph := &Phabricator{
		Host: ts.URL,
	}
	_, err := ph.PhidQuery("token")
	assert.EqualError(t, err, "500 something wrong")
}
