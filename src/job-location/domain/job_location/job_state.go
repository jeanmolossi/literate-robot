package job_location

type JobState struct {
	stateID int
	state   string
}

func NewJobState(stateID int, state string) *JobState {
	if stateID == 0 {
		panic("stateID cannot be 0")
	}

	if state == "" {
		panic("state cannot be empty")
	}

	return &JobState{
		stateID: stateID,
		state:   state,
	}
}

func (j *JobState) StateID() int {
	return j.stateID
}

func (j *JobState) State() string {
	return j.state
}
