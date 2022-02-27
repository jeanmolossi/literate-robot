package query

type JobLocation struct {
	JobID int
	JobCity
	JobState
}

type JobCity struct {
	CityID int
	City   string
}

type JobState struct {
	StateID int
	State   string
}
