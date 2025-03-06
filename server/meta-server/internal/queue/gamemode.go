package queue

// GameMode holds the rules for a match type.
type GameMode struct {
	Name string 		// e.g., "FFA10"
	MaxPlayers int 		// e.v., 10
	IsTeamBased bool 	// For future expansions
}

// FFA10 is the default 10-player free-for-all mode.
var FFA10 = GameMode {
	Name: "FFA10",
	MaxPlayers: 10,
	IsTeamBased: false,
}