package models

var GamesMap = make(map[string]Game)

type Game struct {
  TourneyId string `json:"tourney_id"`
  UserId string `json:"user_id"`
  GameId string `json:"game_id"`

  /// xlogfile fields follow
  Version string `json:"version"`
  Points int `json:"points"`
  Deathdnum int `json:"deathdnum"`
  Deathlev int `json:deathlev"`
  Maxlvl int `json:"maxlvl"`
  Hp int  `json:"hp"`
  Maxhp int  `json:"maxhp"`
  Deaths int  `json:"deaths"`
  Deathdate int `json:"deathdate"` // yyyymmdd
  Birthdate int `json:"birthdate"`
  Uid int  `json:"uid"`
  Role string  `json:"role"`
  Race string  `json:"race"`
  Gender string  `json:"gender"`
  Align string  `json:"align"`
  Name string  `json:"name"`
  Death string  `json:"death"`
  Conduct int `json:"conduct"` //bitfield, e.g. 0x0f0
  Turns int  `json:"turns"`
  Achieve int `json:"achieve"` //bitfield
  Realtime int  `json:"realtime"` //seconds
  Starttime int `json:"starttime"` //unixtime
  Endtime int  `json:"endtime"` //unixtime
  Gender0 string  `json:"gender0"`
  Align0 string  `json:"align0"`
}

func GameById(id string) (game Game) {
  game = GamesMap[id]
  return
}

func NewGameFromXlogLine(line string) (game Game) {
  
}

func NewGameFromMap(m map[string]interface{}) (game Game) {

}
