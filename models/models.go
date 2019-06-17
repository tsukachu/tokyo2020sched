package models

import (
	"encoding/json"
	"time"

	"gopkg.in/guregu/null.v3"
)

type Competition struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Classification struct {
	Id            int64     `db:"id"`
	Name          string    `db:"name"`
	CompetitionId int64     `db:"competition_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type ClassificationWithCompetition struct {
	Id          int64
	Name        string
	Competition Competition
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Place struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type OlympicSchedule struct {
	Id             int64
	Competition    Competition
	Classification Classification
	Title          string
	Begin          time.Time
	End            time.Time
	Place          Place
	Content        null.String
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (c Classification) MarshalJSON() ([]byte, error) {
	if c.Id == 0 {
		return []byte("null"), nil
	}

	type __classification Classification
	return json.Marshal(__classification(c))
}
