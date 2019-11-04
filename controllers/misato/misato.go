package misato

import (
	"encoding/json"
	"errors"
	"github.com/c-f/lel/utils"
	"os"
	"time"

	"fmt"
	"log"
)

// TODO refactor Put commands --> in util

var (
	ErrNoFiles            = errors.New("No MisatoFiles available in this location")
	ErrMissingInformation = errors.New("Invalid entries - fields are missing")
)

// Misato handles the search for commands
type Misato struct {
	Dir string
}

// New Creates a new misato instance
func New(misatoDir string) (*Misato, error) {
	err := os.MkdirAll(misatoDir, 0755)
	misato := &Misato{
		Dir: misatoDir,
	}
	return misato, err
}

// SearchCommands searches for all json files based on the searchstring
func (m *Misato) SearchCommands(search string) (result map[string][]string, err error) {
	result = make(map[string][]string)
	files := utils.FindMisatoLogs(m.Dir)
	if len(files) == 0 {
		err = ErrNoFiles
		return
	}

	for _, file := range files {
		var matches []string
		matches, err = utils.SearchWithPreviewInFile(file, search)
		if err != nil {
			return
		}
		if len(matches) == 0 {
			continue
		}
		origin := utils.TrimPart(file, m.Dir)
		result[origin] = matches
	}
	return result, nil
}

func (m *Misato) GetByTime(start, stop int) (result []CLILog, err error) {
	result = []CLILog{}

	files := utils.FindMisatoLogs(m.Dir)
	if len(files) == 0 {
		err = ErrNoFiles
		return
	}

	for _, fullpath := range files {
		lines, err := utils.ReadLinesFromFiles(fullpath)
		if err != nil {
			log.Printf("\t[%s] %s %s %s", "api:milestone", "get", "error", err.Error())
			continue
		}
		for _, entry := range lines {
			var log CLILog
			err := json.Unmarshal([]byte(entry), &log)
			if err != nil {
				continue
			}

			if log.Timestamp > start && log.Timestamp < stop {
				result = append(result, log)
			}
		}
	}
	return
}

func (m *Misato) PutActivity(activity Activity, id string) error {
	if !activity.Valid() {
		return ErrMissingInformation
	}

	prefix := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s-%s%s", prefix, id, ".activities.json")
	path, err := utils.GetAbsPath(m.Dir, filename)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer f.Close()

	bts, err := json.Marshal(activity)
	if err != nil {
		return err
	}

	_, err = f.Write(bts)
	f.Write([]byte("\n"))
	return err
}

// PutCommands saves the Command
func (m *Misato) PutCommand(clilog CLILog, id string) error {
	if !clilog.Valid() {
		return ErrMissingInformation
	}

	prefix := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s-%s%s", prefix, id, ".misato.json")
	path, err := utils.GetAbsPath(m.Dir, filename)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer f.Close()

	bts, err := json.Marshal(clilog)
	if err != nil {
		return err
	}

	_, err = f.Write(bts)
	f.Write([]byte("\n"))
	return err
}

// getActivities
func (m *Misato) GetActivities() []*Activity {
	files := utils.FindActivities(m.Dir)

	activities := []*Activity{}
	for _, fullpath := range files {

		lines, err := utils.ReadLinesFromFiles(fullpath)
		if err != nil {
			log.Printf("\t[%s] %s %s %s", "api:activities", "get", "error", err.Error())
			continue
		}
		for _, entry := range lines {
			var activity Activity
			err := json.Unmarshal([]byte(entry), &activity)
			if err != nil {

				continue
			}
			activities = append(activities, &activity)
		}
	}
	return activities
}

// aaa
// GetMileStonesViewer
func (m *Misato) GetMileStones() []*Milestone {
	files := utils.FindMilestones(m.Dir)

	milestones := []*Milestone{}
	for _, fullpath := range files {

		lines, err := utils.ReadLinesFromFiles(fullpath)
		if err != nil {
			log.Printf("\t[%s] %s %s %s", "api:milestone", "get", "error", err.Error())
			continue
		}
		for _, entry := range lines {
			var milestone Milestone
			err := json.Unmarshal([]byte(entry), &milestone)
			if err != nil {

				continue
			}
			milestones = append(milestones, &milestone)
		}
	}
	return milestones
}

// PutMileStone saves
func (m *Misato) PutMileStone(milestone Milestone, id string) error {
	if !milestone.Valid() {
		return ErrMissingInformation
	}

	prefix := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s-%s%s", prefix, id, ".milestones.json")
	path, err := utils.GetAbsPath(m.Dir, filename)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	bts, err := json.Marshal(milestone)
	if err != nil {
		return err
	}

	_, err = f.Write(bts)
	f.Write([]byte("\n"))
	return err
}
