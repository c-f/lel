package controllers

type Stats struct{
	Notes int
	Videos int
 	Images int
	Misato int
	Milestones int
	Graphs int

	Nodes int
	Errors []string
}

func NewStats()*Stats{
	return &Stats{
		Errors: []string{},
	}
}