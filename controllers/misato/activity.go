package misato

import (
	"encoding/json"
)

// Activities
type Activity struct {
	User      string `json:"user"`
	Timestamp int    `json:"time,string"`

	Type    string `json:"type"`
	Subject string `json:"subject"`
	Data    string `json:"data"`
}

// CLILogFromJson returns a CLILog from
func ActivityFromJson(raw string) (*Activity, error) {
	var activity Activity
	err := json.Unmarshal([]byte(raw), &activity)
	return &activity, err
}

// Valid checks if all necessary arguments are available
func (ac *Activity) Valid() (ok bool) {
	if ok = ac.User != ""; !ok {
		return
	}
	if ok = ac.Type != ""; !ok {
		return
	}
	return
}
