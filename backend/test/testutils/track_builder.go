package testutils

import "project/internal/pkg/models"

type TrackBuilder struct {
	track models.Track
}

func NewTrackBuilder() *TrackBuilder {
	return &TrackBuilder{}
}

func (b *TrackBuilder) WithID(id int) *TrackBuilder {
	b.track.ID = id
	return b
}

func (b *TrackBuilder) WithName(name string) *TrackBuilder {
	b.track.Name = name
	return b
}

func (b *TrackBuilder) WithCountry(country string) *TrackBuilder {
	b.track.Country = country
	return b
}

func (b *TrackBuilder) WithTown(town string) *TrackBuilder {
	b.track.Town = town
	return b
}

func (b *TrackBuilder) Build() models.Track {
	return b.track
}
