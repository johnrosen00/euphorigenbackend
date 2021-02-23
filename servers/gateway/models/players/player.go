package players

//Player contains player info
type Player struct {
	PlayerID     int64 `json:"playerid"`
	LastPuzzleID int64 `json:"lastpuzzleid"`
}
