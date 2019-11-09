package misato

// Milestone is an information struct for storing milestones :D
type Milestone struct {
	Timestamp int    `json:"time,string"`
	User      string `json:"user"`
	Milestone string `json:"milestone"`
}

// Valid checks, if every necessary attribute is available
func (m *Milestone) Valid() (ok bool) {
	if ok = m.User != ""; !ok {
		return
	}
	if ok = m.Milestone != ""; !ok {
		return
	}
	if ok = m.Timestamp > 0; !ok {
		return
	}

	return
}

// Milestones is a collection of milestones
type Milestones []*Milestone

// Len returns the length of the Milestone, necessary for sorting
func (ms Milestones) Len() int      { return len(ms) }
// Swap swaps entries in the milestone slide
func (ms Milestones) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }

// MilestoneByTime is a Milestones list
type MilestoneByTime struct{ Milestones }

// Less checks if the timestamp of the  milestones are in the range
func (s MilestoneByTime) Less(i, j int) bool {
	return s.Milestones[i].Timestamp < s.Milestones[j].Timestamp
}
