package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Payload struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func Send(incomingURL string, payload Payload) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.WithStack(err)
	}

	resp, err := http.PostForm(
		incomingURL,
		url.Values{"payload": {string(payloadBytes)}},
	)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("bad response status: %d, resp=%s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
