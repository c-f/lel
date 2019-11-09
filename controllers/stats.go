package controllers

// Stats contains all relevant information, which is neccessary for statistical analysis
type Stats struct{
	Notes int `json:"notes"`
	Videos int `json:"videos"`
 	Images int `json:"images"`
	Misato int `json:"misato"`
	Milestones int `json:"milestones"` 
	Graphs int `json:"graphs"`

	Nodes int `json:"nodes"`
	Errors []string  `json:"errors"`
}

// NewStats creates a new Stats objects
func NewStats()*Stats{
	return &Stats{
		Errors: []string{},
	}
}