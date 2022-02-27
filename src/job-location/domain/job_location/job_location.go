package job_location

type JobLocation struct {
	JobID int
	JobCity
	JobState
}

func NewJobLocation(jobID int, jobCity *JobCity, jobState *JobState) *JobLocation {
	if jobCity == nil {
		panic("nil JobCity")
	}

	if jobState == nil {
		panic("nil JobState")
	}
	return &JobLocation{JobID: jobID, JobCity: *jobCity, JobState: *jobState}
}
