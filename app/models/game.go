package models

import (
	"strconv"
	"strings"
)

var GamesMap = make(map[string]Game)

type Game struct {
	TourneyId string `json:"tourney_id"`
	UserId    string `json:"user_id"`
	GameId    string `json:"game_id"`

	/// xlogfile fields follow
	Version   string `json:"version"`
	Points    int64  `json:"points"`
	Deathdnum int64  `json:"deathdnum"`
	Deathlev  int64  `json:deathlev"`
	Maxlvl    int64  `json:"maxlvl"`
	Hp        int64  `json:"hp"`
	Maxhp     int64  `json:"maxhp"`
	Deaths    int64  `json:"deaths"`
	Deathdate int64  `json:"deathdate"` // yyyymmdd
	Birthdate int64  `json:"birthdate"`
	Uid       int64  `json:"uid"`
	Role      string `json:"role"`
	Race      string `json:"race"`
	Gender    string `json:"gender"`
	Align     string `json:"align"`
	Name      string `json:"name"`
	Death     string `json:"death"`
	Conduct   int64  `json:"conduct"` //bitfield, e.g. 0x0f0
	Turns     int64  `json:"turns"`
	Achieve   int64  `json:"achieve"`   //bitfield
	Realtime  int64  `json:"realtime"`  //seconds
	Starttime int64  `json:"starttime"` //unixtime
	Endtime   int64  `json:"endtime"`   //unixtime
	Gender0   string `json:"gender0"`
	Align0    string `json:"align0"`
}

func GameById(id string) (game Game) {
	game = GamesMap[id]
	return
}

func NewGameFromXlogLine(line string) (game Game) {
	gameId_fields := strings.SplitN(line, " ", 2)
	if len(gameId_fields) < 2 {
		return
	}
	gameId := gameId_fields[0]
	fields := gameId_fields[1]

	// build up map of game attributes
	fields_map := make(map[string]string)
	fields_map["GameId"] = gameId
	kvpairs := strings.Split(fields, ":")
	for _, kvpair := range kvpairs {
		key_value := strings.SplitN(kvpair, "=", 2)
		if len(key_value) < 2 {
			break
		}
		key := key_value[0]
		value := key_value[1]

		fields_map[key] = value
	}

	game = NewGameFromStringMap(fields_map)
	return
}

func NewGameFromStringMap(m map[string]string) (game Game) {
	// regular base-10 ints
	points, _ := strconv.ParseInt(m["points"], 32, 10)
	deathdnum, _ := strconv.ParseInt(m["deathdnum"], 32, 10)
	deathlev, _ := strconv.ParseInt(m["deathlev"], 32, 10)
	maxlvl, _ := strconv.ParseInt(m["maxlvl"], 32, 10)
	hp, _ := strconv.ParseInt(m["hp"], 32, 10)
	maxhp, _ := strconv.ParseInt(m["maxhp"], 32, 10)
	deaths, _ := strconv.ParseInt(m["deaths"], 32, 10)
	deathdate, _ := strconv.ParseInt(m["deathdate"], 32, 10)
	birthdate, _ := strconv.ParseInt(m["birthdate"], 32, 10)
	uid, _ := strconv.ParseInt(m["uid"], 32, 10)
	turns, _ := strconv.ParseInt(m["turns"], 32, 10)
	realtime, _ := strconv.ParseInt(m["realtime"], 32, 10)
	starttime, _ := strconv.ParseInt(m["starttime"], 32, 10)
	endtime, _ := strconv.ParseInt(m["endtime"], 32, 10)

	// hex bitfields
	achieve, _ := strconv.ParseInt(m["achieve"], 32, 16)
	conduct, _ := strconv.ParseInt(m["conduct"], 32, 16)

	game = Game{
		TourneyId: m["tourneyId"],
		UserId:    m["userId"],
		GameId:    m["GameId"],
		Version:   m["version"],
		Points:    points,
		Deathdnum: deathdnum,
		Deathlev:  deathlev,
		Maxlvl:    maxlvl,
		Hp:        hp,
		Maxhp:     maxhp,
		Deaths:    deaths,
		Deathdate: deathdate,
		Birthdate: birthdate,
		Uid:       uid,
		Role:      m["role"],
		Race:      m["race"],
		Gender:    m["gender"],
		Align:     m["align"],
		Name:      m["name"],
		Death:     m["death"],
		Conduct:   conduct,
		Turns:     turns,
		Achieve:   achieve,
		Realtime:  realtime,
		Starttime: starttime,
		Endtime:   endtime,
		Gender0:   m["gender0"],
		Align0:    m["align0"]}
	return
}

func (game Game) Save() {
	GamesMap[game.GameId] = game
	return
}
