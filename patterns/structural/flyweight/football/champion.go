package football

import "time"

const (
	TeamA = iota
	TeamB
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

func getTeamFactory(team int) Team {
	switch team {
	case TeamB:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

type TeamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

func (t *TeamFlyweightFactory) GetTeam(teamName int) *Team {
	if t.createdTeams[teamName] != nil {
		return t.createdTeams[teamName]
	}

	team := getTeamFactory(teamName)
	t.createdTeams[teamName] = &team

	return t.createdTeams[teamName]
}

func (t *TeamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
