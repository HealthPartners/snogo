package snogo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//DefaultServiceNowClient get the client from the environment vars
func DefaultServiceNowClient() *ServiceNowInstance {
	return NewInstance(
		os.Getenv("SERVICE_NOW_INSTANCE_NAME"),
		os.Getenv("SERVICE_NOW_USERNAME"),
		os.Getenv("SERVICE_NOW_PASSWORD"))
}

//IncidentCreationPayload Data to create an incident in service-now
type IncidentCreationPayload struct {
	AssignmentGroup  string      `json:"assignment_group"`
	CmdbCI           string      `json:"cmdb_ci"`
	ContactType      string      `json:"contact_type"` //Auto Ticket
	Customer         string      `json:"caller_id"`
	Description      string      `json:"description"`
	Impact           json.Number `json:"impact"`
	ShortDescription string      `json:"short_description"`
	State            json.Number `json:"state"`
	Urgency          json.Number `json:"urgency"`
}

//ServiceNowInstance should hold the necessary data for a client of the ServiceNow TableAPI
type ServiceNowInstance struct {
	name       string
	baseURL    string
	authHeader string

	client *http.Client
}

// NewInstance - Create new service-now instance
func NewInstance(name, user, pass string) *ServiceNowInstance {
	return &ServiceNowInstance{
		name:       name,
		baseURL:    fmt.Sprintf("https://%s.service-now.com", name),
		authHeader: fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte(user+":"+pass))),
		client:     http.DefaultClient,
	}
}

//Create an incident in service-now based on a post body
func (inst *ServiceNowInstance) Create(table string, body []byte) (string, error) {
	fmt.Printf("Creating a service now incident")
	req, _ := http.NewRequest("POST", buildPostURL(inst.baseURL, table), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", inst.authHeader)
	resp, reqErr := inst.client.Do(req)
	if reqErr != nil {
		fmt.Println(reqErr)
		return "", reqErr
	}
	// TODO: ignoring an error on this Close
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("responseBody: %s", string(responseBody))
	return string(responseBody), nil
}

func buildPostURL(baseURL, table string) string {
	return fmt.Sprintf("%s/api/now/table/%s", baseURL, table)
}
