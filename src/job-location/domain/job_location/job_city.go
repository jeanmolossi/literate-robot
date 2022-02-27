package job_location

type JobCity struct {
	cityID int
	city   string
}

func NewJobCity(cityID int, city string) *JobCity {
	if cityID == 0 {
		panic("cityID cannot be 0")
	}

	if city == "" {
		panic("city cannot be empty")
	}

	return &JobCity{
		cityID: cityID,
		city:   city,
	}
}

func (j *JobCity) CityID() int {
	return j.cityID
}

func (j *JobCity) City() string {
	return j.city
}
