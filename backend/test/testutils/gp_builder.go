package testutils

import "project/internal/pkg/models"

type GPBuilder struct {
	gp models.GrandPrix
}

func NewGPBuilder() *GPBuilder {
	return &GPBuilder{}
}

func (b *GPBuilder) WithID(id int) *GPBuilder {
	b.gp.ID = id
	return b
}

func (b *GPBuilder) WithSeason(season int) *GPBuilder {
	b.gp.Season = season
	return b
}

func (b *GPBuilder) WithName(name string) *GPBuilder {
	b.gp.Name = name
	return b
}

func (b *GPBuilder) WithDateNum(dateNum int) *GPBuilder {
	b.gp.DateNum = dateNum
	return b
}

func (b *GPBuilder) WithMonth(month string) *GPBuilder {
	b.gp.Month = month
	return b
}

func (b *GPBuilder) WithPlace(place string) *GPBuilder {
	b.gp.Place = place
	return b
}

func (b *GPBuilder) WithTrackId(trackId int) *GPBuilder {
	b.gp.TrackId = trackId
	return b
}

func (b *GPBuilder) Build() models.GrandPrix {
	return b.gp
}
