package misato

// Milestone is an information struct for storing milestones :D
type Milestone struct {
	Timestamp int    `json:"time,string"`
	User      string `json:"user"`
	Milestone string `json:"milestone"`
}

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

type Milestones []*Milestone

func (ms Milestones) Len() int      { return len(ms) }
func (ms Milestones) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }

type MilestoneByTime struct{ Milestones }

func (s MilestoneByTime) Less(i, j int) bool {
	return s.Milestones[i].Timestamp < s.Milestones[j].Timestamp
}
