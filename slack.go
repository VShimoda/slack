package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// URLPrefix is a endpoint for slack api
	URLPrefix = "https://slack.com/api/"
)

// Slack is for all clients
type Slack struct {
	Token   string
	Request http.Client
}

// New initial Slack client for api
func New(token string) *Slack {
	client := http.Client{Timeout: time.Duration(15) * time.Second}
	return &Slack{
		Token:   token,
		Request: client,
	}
}

func requestPostForm(client *http.Client, endpoint string, values *url.Values) ([]byte, error) {
	resp, err := client.PostForm(endpoint, *values)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
