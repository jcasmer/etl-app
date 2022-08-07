package entity

import (
	"strings"
	"time"
)

type Incidents struct {
	Incident []Incident
}
type Incident struct {
	ID          int          `json:"id" validate:"required"`
	Name        string       `json:"name" validate:"required"`
	Discovered  IncidentTime `json:"discovered" validate:"required"`
	Description string       `json:"description" validate:"required"`
	Status      string       `json:"status" validate:"required"`
}

type IncidentTime time.Time

func (c *IncidentTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = IncidentTime(t) //set result using the pointer
	return nil
}

func (c IncidentTime) Format() string {
	return time.Time(c).Format("2006-01-02")
}

func (c IncidentTime) ParseTime() time.Time {
	t := time.Time(c)
	return t
}
