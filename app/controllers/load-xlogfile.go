package controllers

import (
	"eit_go/app/models"
	"fmt"
	"github.com/robfig/revel"
	"io/ioutil"
	"reflect"
	"strings"
)

//var GamesMap = make(map[string]models.Game)

type XlogPlugin struct {
	rev.EmptyPlugin
}

func (p XlogPlugin) OnAppStart() {
	fmt.Printf("Reading the Xlogfile... ")
	importXlog("scores.xlogfile.2010")
	fmt.Printf("done.\n")
}

func importXlog(xlogfile string) {
	fmt.Printf("importXlog ")
	bytes, _ := ioutil.ReadFile(xlogfile)
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		g := models.NewGameFromXlogLine(line)
		//g.Save()
		models.GamesMap[g.GameId] = g
	}
	fmt.Printf("imported %d lines. ", len(models.GamesMap))
}

func importXlogLine(line string) {
	gameId_fields := strings.SplitN(line, " ", 2)
	if len(gameId_fields) < 2 {
		return
	}
	gameId := gameId_fields[0]
	fields := gameId_fields[1]

	// build up map of game attributes
	fields_map := make(map[string]interface{})
	fields_map["GameId"] = gameId
	kvpairs := strings.Split(fields, ":")
	for _, kvpair := range kvpairs {
		key_value := strings.SplitN(kvpair, "=", 2)
		if len(key_value) < 2 {
			break
		}
		key := key_value[0]
		value := key_value[1]
		//fmt.Printf("setting %s to %s\n", key, value)

		fields_map[key] = value
	}

	// convert the map into a Game
	game := mapToGame(fields_map)
	importGame(game)
}

func mapToGame(m map[string]interface{}) (game models.Game) {
	gameval := reflect.Indirect(reflect.ValueOf(&game).Elem())

	for i := 0; i < gameval.NumField(); i++ {
		key := gameval.Type().Field(i).Name
		if m[key] == nil {
			gameval.Field(i).Set(reflect.Zero(gameval.Field(i).Type()))
		} else {
			gameval.Field(i).Set(reflect.ValueOf(m[key]))
		}
	}

	// game = gameval.Interface().(models.Game)
	return
}

/*
// this is pulled from https://github.com/sdegutis/go.mapstruct
func mapToStruct(m map[string]interface{}, s interface{}) {
    v := reflect.Indirect(reflect.ValueOf(s))

    for i := 0; i < v.NumField(); i++ {
        key := v.Type().Field(i).Name
        v.Field(i).Set(reflect.ValueOf(m[key]))
    }
}
*/

func importGame(game models.Game) {
	models.GamesMap[game.GameId] = game
	return
}

func init() {
	models.GamesMap = make(map[string]models.Game)
	rev.RegisterPlugin(XlogPlugin{})
}
