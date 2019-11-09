package graph

// Meta contains all metaannotation information from a input source.
//

import (
	"errors"

	"regexp"
	"strings"

	"github.com/c-f/lel/utils"
)

var (
	// ErrNoValidRelation indicates, that a syntax error happens in the relation
	ErrNoValidRelation = errors.New("No Valid Relation")

	referenceReg = regexp.MustCompile(utils.REFERENCE_REG_PATTERN)
)

// GetRefInfo returns if the raw input is ref or not
// @ref hello
// this is only valid for @entity info types
func GetRefInfo(raw string) (result string) {
	if isRelation := strings.HasPrefix(raw, "(") || strings.HasSuffix(raw, ")"); isRelation {
		return ""
	}

	return raw
}

// GetRefRelation Splits the raw input into the parts, the origin is the key 
// (subject)->[relation]->(object)
func GetRefRelation(raw, origin string) (subject, relation, object string, err error) {
	results := referenceReg.FindAllStringSubmatch(raw, -1)
	if len(results) <= 0 {
		err = ErrNoValidRelation
		return
	}
	if subject = results[0][1]; subject == utils.THIS_REFERENCE {
		subject = origin
	}
	if relation = results[0][2]; relation == utils.THIS_REFERENCE {
		relation = origin
	}
	if object = results[0][3]; object == utils.THIS_REFERENCE {
		object = origin
	}

	return
}

// IsRelationAlias checks if a relation string is an alias or not
// this is done by checking if the first char is a collon
// (subject)->[:hello]->(object)
func IsRelationAlias(relation string) bool {
	if strings.HasPrefix(relation, ":") {
		return true
	}
	return false
}

// GetAlias returns the alias, without the collon prefix
func GetAlias(relation string) string {
	return strings.TrimPrefix(":", relation)
}
