package testutils

import (
	"project/internal/pkg/models"
	"time"
)

func BuildDrivers(driverBuilder *DriverBuilder, ids []int, names []string, countries []string, birthdates []string) []*models.Driver {
	drivers := make([]*models.Driver, 0)
	for i := range ids {
		driver := driverBuilder.
			WithID(ids[i]).
			WithName(names[i]).
			WithCountry(countries[i]).
			WithBirthDate(birthdates[i]).
			Build()
		drivers = append(drivers, &driver)
	}

	return drivers
}

func BuildGrandPrixes(gpBuilder *GPBuilder, ids []int, seasons []int, names []string, dateNums []int, months []string, places []string, trackIds []int) []*models.GrandPrix {
	gps := make([]*models.GrandPrix, 0)
	for i := range ids {
		gp := gpBuilder.
			WithID(ids[i]).
			WithSeason(seasons[i]).
			WithName(names[i]).
			WithDateNum(dateNums[i]).
			WithMonth(months[i]).
			WithPlace(places[i]).
			WithTrackId(trackIds[i]).
			Build()
		gps = append(gps, &gp)
	}

	return gps
}

func BuildQualResults(qBuilder *QualResBuilder, ids []int, driverPlaces []int, driverIds []int, driverNames []string, teamIds []int, teamNames []string, q1Times []time.Time, q2Times []time.Time, q3Times []time.Time, gpIds []int, gpNames []string) ([]*models.QualResult, []*models.QualResultView) {
	quals := make([]*models.QualResult, 0)
	qualViews := make([]*models.QualResultView, 0)
	for i := range ids {
		qual, qualView := qBuilder.
			WithID(ids[i]).
			WithDriverPlace(driverPlaces[i]).
			WithDriverId(driverIds[i]).
			WithDriverName(driverNames[i]).
			WithTeamId(teamIds[i]).
			WithTeamName(teamNames[i]).
			WithQ1time(q1Times[i]).
			WithQ2time(q2Times[i]).
			WithQ3time(q3Times[i]).
			WithGPId(gpIds[i]).
			WithGPName(gpNames[i]).
			Build()
		quals = append(quals, &qual)
		qualViews = append(qualViews, &qualView)
	}

	return quals, qualViews
}

func BuildRaceResults(rBuilder *RaceResBuilder, ids []int, driverPlaces []int, driverIds []int, driverNames []string, teamIds []int, teamNames []string, gpIds []int, gpNames []string) ([]*models.RaceResult, []*models.RaceResultView) {
	races := make([]*models.RaceResult, 0)
	raceViews := make([]*models.RaceResultView, 0)
	for i := range ids {
		race, raceView := rBuilder.
			WithID(ids[i]).
			WithDriverPlace(driverPlaces[i]).
			WithDriverId(driverIds[i]).
			WithDriverName(driverNames[i]).
			WithTeamId(teamIds[i]).
			WithTeamName(teamNames[i]).
			WithGPId(gpIds[i]).
			WithGPName(gpNames[i]).
			Build()
		races = append(races, &race)
		raceViews = append(raceViews, &raceView)
	}

	return races, raceViews
}

func BuildTeams(teamBuilder *TeamBuilder, ids []int, names []string, countries []string, bases []string) []*models.Team {
	teams := make([]*models.Team, 0)
	for i := range ids {
		team := teamBuilder.
			WithID(ids[i]).
			WithName(names[i]).
			WithCountry(countries[i]).
			WithBase(bases[i]).
			Build()
		teams = append(teams, &team)
	}

	return teams
}

func BuildTracks(trackBuilder *TrackBuilder, ids []int, names []string, countries []string, towns []string) []*models.Track {
	tracks := make([]*models.Track, 0)
	for i := range ids {
		track := trackBuilder.
			WithID(ids[i]).
			WithName(names[i]).
			WithCountry(countries[i]).
			WithTown(towns[i]).
			Build()
		tracks = append(tracks, &track)
	}

	return tracks
}

func BuildUsers(userBuilder *UserBuilder, ids []int, logins []string, password []string, roles []string) []*models.User {
	users := make([]*models.User, 0)
	for i := range ids {
		user := userBuilder.
			WithID(ids[i]).
			WithLogin(logins[i]).
			WithPassword(password[i]).
			WithRole(roles[i]).
			Build()
		users = append(users, &user)
	}

	return users
}

func MakePointerSlice[T any](src []T) []*T {
	resSlice := make([]*T, len(src))
	for idx := range src {
		val := new(T)
		*val = src[idx]
		resSlice[idx] = val
	}

	return resSlice
}
