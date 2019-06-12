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
	// guregu/null/zero で初期化はするが API 的には不正なので null にする
	if c.Id == 0 {
		return []byte("null"), nil
	}

	// Classification をそのまま使うと無限ループするので Alias で回避
	type Alias Classification

	return json.Marshal(&struct{ Alias }{Alias: (Alias)(c)})
}
