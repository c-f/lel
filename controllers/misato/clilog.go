package misato

import (
	"encoding/json"
	"errors"
)

var (
	ErrMissingFields = errors.New("Missing Fields")
)

// CLILOg is a struct containing all relevant information about alll the things
type CLILog struct {
	User      string `json:"user"`
	Command   string `json:"command"`
	TermUUID  string `json:"termuuid"`
	Timestamp int    `json:"time,string"`
}

// CLILogFromJson returns a CLILog from
func CLILogFromJson(raw string) (*CLILog, error) {
	var clilog CLILog
	err := json.Unmarshal([]byte(raw), &clilog)
	return &clilog, err
}

// Valid checks if all necessary arguments are available
func (cli *CLILog) Valid() (ok bool) {
	if ok = cli.User != ""; !ok {
		return
	}
	if ok = cli.Command != ""; !ok {
		return
	}

	return
}
