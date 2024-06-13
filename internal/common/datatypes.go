package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "Other"
)

func (r *Gender) Scan(value interface{}) error {
	*r = Gender(value.([]byte))
	return nil
}

func (r Gender) Value() (driver.Value, error) {
	return string(r), nil
}

type Status string

const (
	StatusActive   Status = "Active"
	StatusInactive Status = "Inactive"
	StatusPending  Status = "Pending"
)

func (r *Status) Scan(value interface{}) error {
	*r = Status(value.([]byte))
	return nil
}

func (r Status) Value() (driver.Value, error) {
	return string(r), nil
}

type JSON json.RawMessage

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
