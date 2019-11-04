package graph

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/c-f/lel/utils"
)

// MetaInformation contains raw information from the file
type MetaInfo struct {
	Names  []string `json:"names"`
	Entity string   `json:"entity"`

	// optional
	Tags      []string `json:"tags"`
	Relations []string `json:"relations"`
	Label     string   `json:"label"`

	References []string `json:"references"`
	// icon url
	Icon string `json:"icon"`

	// id is Origin
	Id string `json:"id"`

	// Todos
	Todos []Todo `json:"todos"`

	// Date
	Modified time.Time `json:"modified"`

	// state defines a condition where the match might
	state int    `json:"-"`
	tmp   string `json:"-"`
}

func (m *MetaInfo) AddRef(relation string) {
	m.References = append(m.References, relation)
}

const (
	UNKNOWN_TYPE int = iota
	// is node
	NODE_TYPE
	// is edge
	EDGE_TYPE
	// is info
	INFO_TYPE
	// is bulk
	BULK_TYPE

	// states
	NORMAL_STATE int = iota
	//
	BEGIN_TODO
)

// GetType returns the type information based on the entity
func (m *MetaInfo) GetType() int {
	switch m.Entity {
	case "info":
		return INFO_TYPE
	case "ref":
		return EDGE_TYPE
	case "bulk":
		return BULK_TYPE
	default:
		return NODE_TYPE
	}
}

// GetLabel returns the label or the first name of the entity
func (m *MetaInfo) GetLabel() (label string) {
	label = m.Label
	if label == "" && m.Valid() {
		return m.Names[0]
	}
	return
}

// Valid returns wether the file has enough information
// At least names and entity must be defined
func (m *MetaInfo) Valid() bool {

	if len(m.Names) <= 0 {
		return false
	}
	if m.Entity == "" {
		return false
	}
	return true
}

// Equal checks if the provided MetaInfo contains the exact same information
func (m *MetaInfo) Equal(other *MetaInfo) bool {
	if m.Entity != other.Entity {
		return false
	}
	if ok := utils.SliceAreEqual(m.Names, other.Names); !ok {
		return false
	}
	if ok := utils.SliceAreEqual(m.Tags, other.Tags); !ok {
		return false
	}
	if ok := utils.SliceAreEqual(m.Relations, other.Relations); !ok {
		return false
	}
	if m.Label != other.Label {
		return false
	}
	return true
}

// SearchAndCollect adds metainformation from the line into the object, if the metaprefix matches
func (m *MetaInfo) Collect(line string) (err error) {
	if !strings.HasPrefix(line, "@") && m.state == NORMAL_STATE {
		return nil
	}

	// states
	if m.state == BEGIN_TODO {
		if todo, ok := ConvertTodo(line, m.tmp); !ok {
			m.resetState()
		} else {
			m.Todos = append(m.Todos, *todo)
		}
		return
	}

	// add strings to Metainformation
	switch {
	case strings.HasPrefix(line, utils.TAGS_PATTERN):
		m.Tags = append(m.Tags, utils.SplitAndTrim(line, utils.TAGS_PATTERN)...)

	case strings.HasPrefix(line, utils.NAME_PATTERN):
		m.Names = append(m.Names, utils.SplitAndTrim(line, utils.NAME_PATTERN)...)

	case strings.HasPrefix(line, utils.REFERENCE_PATTERN):
		m.Relations = append(m.Relations, utils.SplitAndTrim(line, utils.REFERENCE_PATTERN)...)

	case strings.HasPrefix(line, utils.LABEL_PATTERN):
		m.Label = utils.SplitAndTrim(line, utils.LABEL_PATTERN)[0]

	case strings.HasPrefix(line, utils.ENTITY_PATTERN):
		m.Entity = strings.ToLower(utils.SplitAndTrim(line, utils.ENTITY_PATTERN)[0])

	case strings.HasPrefix(line, utils.ICON_PATTERN):
		m.Icon = utils.SplitAndTrim(line, utils.ICON_PATTERN)[0]

	case strings.HasPrefix(line, utils.TODO_PATTERN):
		m.tmp = utils.SplitAndTrim(line, utils.TODO_PATTERN)[0]
		m.state = BEGIN_TODO
	}

	return
}

// resetState reset the current state
func (m *MetaInfo) resetState() {
	m.state = NORMAL_STATE
	m.tmp = ""
}

// MetaInfoFromFile creates a new metainfo from a file
func MetaInfoFromFile(filename string, id string) (m *MetaInfo, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	finfo, err := os.Stat(filename)
	if err != nil {
		return
	}

	defer f.Close()
	m, err = NewMetaInfoFromReader(f, id)
	m.Modified = finfo.ModTime()

	return m, err
}

// NewMetaInformationFromReader creates a new metaInformation, through reading line by line
func NewMetaInfoFromReader(r io.Reader, id string) (m *MetaInfo, err error) {
	m = &MetaInfo{
		Id:         id,
		Tags:       []string{},
		References: []string{},
		Names:      []string{},
		Todos:      []Todo{},
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		if err = m.Collect(scanner.Text()); err != nil {
			return nil, err
		}
	}
	err = scanner.Err()
	return
}
