package models

import "time"

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
