package data

type Race struct {
	ID   int
	Name string
}

type Timers struct {
	ID        int
	RacerName string
	Duration  int
	RaceID    int
}

type AuthKey struct {
	ID   int
	AuthKey  string
}