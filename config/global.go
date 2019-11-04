package config

// Config contains all information which are used by LEL
type Config struct {
	Server  ServerConfig
	Project ProjectConfig

	// name of the operator necessary for instructions
	// TODO currently not used
	Operator string
}
