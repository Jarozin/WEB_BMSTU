package testutils

import (
	"project/internal/pkg/models"
	"time"
)

type QualResBuilder struct {
	qualRes     models.QualResult
	qualResView models.QualResultView
}

func NewQualResBuilder() *QualResBuilder {
	return &QualResBuilder{}
}

func (b *QualResBuilder) WithID(id int) *QualResBuilder {
	b.qualRes.ID = id
	b.qualResView.ID = id
	return b
}

func (b *QualResBuilder) WithDriverPlace(driverPlace int) *QualResBuilder {
	b.qualRes.DriverPlace = driverPlace
	b.qualResView.DriverPlace = driverPlace
	return b
}

func (b *QualResBuilder) WithDriverId(driverId int) *QualResBuilder {
	b.qualRes.DriverId = driverId
	return b
}

func (b *QualResBuilder) WithDriverName(driverName string) *QualResBuilder {
	b.qualResView.DriverName = driverName
	return b
}

func (b *QualResBuilder) WithTeamId(teamId int) *QualResBuilder {
	b.qualRes.TeamId = teamId
	return b
}

func (b *QualResBuilder) WithTeamName(teamName string) *QualResBuilder {
	b.qualResView.TeamName = teamName
	return b
}

func (b *QualResBuilder) WithQ1time(q1time time.Time) *QualResBuilder {
	b.qualRes.Q1time = q1time
	b.qualResView.Q1time = q1time
	return b
}

func (b *QualResBuilder) WithQ2time(q2time time.Time) *QualResBuilder {
	b.qualRes.Q2time = q2time
	b.qualResView.Q2time = q2time
	return b
}

func (b *QualResBuilder) WithQ3time(q3time time.Time) *QualResBuilder {
	b.qualRes.Q3time = q3time
	b.qualResView.Q3time = q3time
	return b
}

func (b *QualResBuilder) WithGPId(gpId int) *QualResBuilder {
	b.qualRes.GPId = gpId
	return b
}

func (b *QualResBuilder) WithGPName(gpName string) *QualResBuilder {
	b.qualResView.GPName = gpName
	return b
}

func (b *QualResBuilder) Build() (models.QualResult, models.QualResultView) {
	return b.qualRes, b.qualResView
}
