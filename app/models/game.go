package models

type Game struct {
  TourneyId string `json:"tourney_id"`
  UserId string `json:"user_id"`
  GameId string `json:"game_id"`

  /// xlogfile fields follow
  Version string
  Points int
  Deathdnum int
  Deathlev int
  Maxlvl int
  Hp int
  Maxhp int
  Deaths int
  Deathdate int // yyyymmdd
  Birthdate int // yyyymmdd
  Uid int
  Role string
  Race string
  Gender string
  Align string
  Name string
  Death string
  Conduct int // bitfield
  Turns int
  Achieve int // bitfield
  Realtime int // seconds
  Starttime int // seconds since epoch
  Endtime int // seconds since epoch
  Gender0 string
  Align0 string
}

