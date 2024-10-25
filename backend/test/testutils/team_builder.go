package testutils

import "project/internal/pkg/models"

type TeamBuilder struct {
	team models.Team
}

func NewTeamBuilder() *TeamBuilder {
	return &TeamBuilder{}
}

func (b *TeamBuilder) WithID(id int) *TeamBuilder {
	b.team.ID = id
	return b
}

func (b *TeamBuilder) WithName(name string) *TeamBuilder {
	b.team.Name = name
	return b
}

func (b *TeamBuilder) WithCountry(country string) *TeamBuilder {
	b.team.Country = country
	return b
}

func (b *TeamBuilder) WithBase(base string) *TeamBuilder {
	b.team.Base = base
	return b
}

func (b *TeamBuilder) Build() models.Team {
	return b.team
}
