package graph

import (
	"regexp"
	"strings"
)


// Todo is the structor of the todo.txt
// https://github.com/todotxt/todo.txt#todotxt-format-rules
type Todo struct {
	Complete       bool     `json:"done"`
	Priority       string   `json:"priority"`
	CompletionDate string   `json:"completion_date"`
	CreationDate   string   `json:"creation_date"`
	Description    string   `json:"description"`
	Context        string   `json:"context"`
	Projects       []string `json:"projects"`

	// Special tokens
	RepeatToken string `json:"repeat"`
	ActionToken string `json:"action"`
	DueToken    string `json:"due"`

	Raw string `json:"_raw"`
}

// special fields
// todo
const (
	CONTEXT_TAG = "@"
	PROJECT_TAG = "+"

	// Special Fields inside the description field
	DUE_PATTERN    = "due:"
	REPEAT_PATTERN = "repeat:"
	ACTION_PATTERN = "lel:"
)

// NewTodo returns an empty todo
func NewTodo() *Todo {
	return &Todo{
		Projects: []string{},
	}
}

// ConvertTodo parses the raw string and convert it into a todo
func ConvertTodo(raw string, context string) (todo *Todo, noerror bool) {
	if isEmpty := emptyReg.MatchString(raw); isEmpty {
		return
	}

	parts := strings.Split(raw, " ")
	todo = NewTodo()

	current, parts, ok := pop(parts)

	// --[done?]--
	if strings.HasPrefix(current, "x") || strings.HasPrefix(current, "[x]") {
		if current, parts, ok = pop(parts); !ok {
			return
		}
		todo.Complete = true

	} else if current == "[" {
		if len(parts) < 2 {
			return
		}
		if current, parts, ok = pop(parts); current != "]" {
			return
		}
		current, parts, ok = pop(parts)
		todo.Complete = false
	}

	// --[priority]--
	if token := prioReg.FindString(current); token != "" {
		if current, parts, ok = pop(parts); !ok {
			return
		}
		todo.Priority = token
	}

	// --[date]--
	if token := dateReg.FindString(current); token != "" {
		if current, parts, ok = pop(parts); !ok {
			return
		}
		if token2 := dateReg.FindString(current); token != "" {
			todo.CreationDate = token
			todo.CompletionDate = token2
		} else {
			todo.CreationDate = token

		}
	}
	// onStack
	parts = append([]string{current}, parts...)
	todo.Description = strings.Join(parts, " ")
	todo.Context = context
	todo.Raw = raw

	// Description flags
	for _, part := range parts {
		switch {
		case strings.HasPrefix(part, CONTEXT_TAG):
			todo.Context = part

		case strings.HasPrefix(part, PROJECT_TAG):
			todo.Projects = append(todo.Projects, part)

		case strings.HasPrefix(part, DUE_PATTERN):
			todo.DueToken = strings.TrimPrefix(part, DUE_PATTERN)

		case strings.HasPrefix(part, REPEAT_PATTERN):
			todo.RepeatToken = strings.TrimPrefix(part, REPEAT_PATTERN)

		case strings.HasPrefix(part, ACTION_PATTERN):
			todo.ActionToken = strings.TrimPrefix(part, ACTION_PATTERN)
		}
	}

	noerror = true
	return
}

// pop returns the first item of the slice
func pop(in []string) (item string, out []string, ok bool) {
	if ok = len(in) != 0; !ok {
		return
	}

	item, out = in[0], in[1:]
	return
}

// Regular expressions for specific inputs
var (
	prioReg  = regexp.MustCompile(`^(\([A-Z]\))$`)
	dateReg  = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	emptyReg = regexp.MustCompile(`^\s*$`)
)

// full example first part
// x (A) 2019-08-01 2019-08-01 ........
// ...........
// (A) ...........
// (B) 2019-08-01 ...........
// (A) 2019-08-01 2019-08-01 ........
// x (A) 2019-08-01 2019-08-01 ........
// 2019-08-01 .........
