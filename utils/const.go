package utils

const (
	// FKKK represents something went bad ¯\_(ツ)_/¯
	FKKK = "¯\\_(ツ)_/¯"
)

const (

	// TODO_PATTERN define the todo pattern for actions, which should be
	// should be in form of
	// @todo: tags-related
	// [ ] Something
	// later some form of todo.txt might come in handy.
	//
	TODO_PATTERN = "@todo "

	// Tags define tags for a specific document
	TAGS_PATTERN = "@tags "

	// name defines aliases for later references
	// "this" can't be chosen, must be unique
	NAME_PATTERN = "@name "

	// the first name will be used, if @display is not defined
	LABEL_PATTERN = "@label "

	// entity defines the different types (also defines custom icons)
	ENTITY_PATTERN = "@entity "

	// icon defines the src of custom svg images
	ICON_PATTERN = "@icon "

	// References contain information for specific objects
	// which can be link by providing the name or an alias of the object
	// and the relation it might have
	// Example:
	// @ref example
	// 		(diese resource is refereced for example (only works with entity: info))
	// @ref (this)->[:dns]->(CW)
	// 		(diese resource hat einen dns channel zum object CW)
	// @ref (fireye)->[installed]->(this)
	// 		other resources can be linked to this
	// @ref (CW)->[this]->(fireeye)
	// 		references can also define the relation itself
	REFERENCE_PATTERN = "@ref "

	VAL_DELIMITER = ","

	REFERENCE_REG_PATTERN = `\(([^)]*)\)->\[([^\]]*)\]->\(([^)]*)\)`
	THIS_REFERENCE        = "this"

	// Supported Icons
	COMPUTER_SHAPE = ""

	CLIENT   = ""
	SERVER   = ""
	DC       = ""
	ACCOUNT  = ""
	PROXY    = ""
	DATABASE = ""
	FILE     = ""
	CLOUD    = ""
	CW       = ""
)
