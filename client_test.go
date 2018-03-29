package snogo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestClientNewInstance(t *testing.T) {
	instance := NewInstance("testname", "user", "pass")
	assert := assertions.New(t)
	assert.So(instance.baseURL, should.Equal, "https://testname.service-now.com")
	assert.So(instance.authHeader, should.ContainSubstring, "Basic")
	assert.So(instance.client, should.HaveSameTypeAs, http.DefaultClient)
}
func TestInstanceCreate(t *testing.T) {
	assert := assertions.New(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, _ := ioutil.ReadAll(r.Body)
		assert.So(string(body), should.ContainSubstring, "description")
		fmt.Fprint(w, `{"sys_id":"abcdefg","number":"INC88888"}`)
	}))
	defer ts.Close()

	instance := &ServiceNowInstance{
		client:     http.DefaultClient,
		baseURL:    ts.URL,
		authHeader: "Basic: abc123",
	}
	incident := &IncidentCreationPayload{
		Description: "The missiles, they are coming!",
	}
	json, _ := json.Marshal(incident)
	output, _ := instance.Create("incident", json)
	assert.So(output, should.Equal, `{"sys_id":"abcdefg","number":"INC88888"}`)
}
