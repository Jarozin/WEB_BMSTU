package testutils

import (
	"project/internal/pkg/models"
)

type RaceResBuilder struct {
	raceRes     models.RaceResult
	raceResView models.RaceResultView
}

func NewRaceResBuilder() *RaceResBuilder {
	return &RaceResBuilder{}
}

func (b *RaceResBuilder) WithID(id int) *RaceResBuilder {
	b.raceRes.ID = id
	b.raceResView.ID = id
	return b
}

func (b *RaceResBuilder) WithDriverPlace(driverPlace int) *RaceResBuilder {
	b.raceRes.DriverPlace = driverPlace
	b.raceResView.DriverPlace = driverPlace
	return b
}

func (b *RaceResBuilder) WithDriverId(driverId int) *RaceResBuilder {
	b.raceRes.DriverId = driverId
	return b
}

func (b *RaceResBuilder) WithDriverName(driverName string) *RaceResBuilder {
	b.raceResView.DriverName = driverName
	return b
}

func (b *RaceResBuilder) WithTeamId(teamId int) *RaceResBuilder {
	b.raceRes.TeamId = teamId
	return b
}

func (b *RaceResBuilder) WithTeamName(teamName string) *RaceResBuilder {
	b.raceResView.TeamName = teamName
	return b
}

func (b *RaceResBuilder) WithGPId(gpId int) *RaceResBuilder {
	b.raceRes.GPId = gpId
	return b
}

func (b *RaceResBuilder) WithGPName(gpName string) *RaceResBuilder {
	b.raceResView.GPName = gpName
	return b
}

func (b *RaceResBuilder) Build() (models.RaceResult, models.RaceResultView) {
	return b.raceRes, b.raceResView
}
