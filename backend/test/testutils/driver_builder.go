package testutils

import "project/internal/pkg/models"

type DriverBuilder struct {
	driver models.Driver
}

func NewDriverBuilder() *DriverBuilder {
	return &DriverBuilder{}
}

func (b *DriverBuilder) WithID(id int) *DriverBuilder {
	b.driver.ID = id
	return b
}

func (b *DriverBuilder) WithName(name string) *DriverBuilder {
	b.driver.Name = name
	return b
}

func (b *DriverBuilder) WithCountry(country string) *DriverBuilder {
	b.driver.Country = country
	return b
}

func (b *DriverBuilder) WithBirthDate(birthdate string) *DriverBuilder {
	b.driver.BirthDate = birthdate
	return b
}

func (b *DriverBuilder) Build() models.Driver {
	return b.driver
}
