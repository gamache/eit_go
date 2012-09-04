package main

import (
  "./gorest"
  "net/http"
  "encoding/json"
  "reflect"
  "fmt"
  "strings"
)


var games map[string]Game


func main() {
  games = make(map[string]Game)

  // gotta register the marshaller before registering the service
  marshaller := ourMarshaller()
  gorest.RegisterMarshaller("application/json", marshaller)
  if gorest.GetMarshallerByMime("application/json") != marshaller {
    fmt.Printf("OH DIP, RegisterMarshaller() didn't work. We aren't using a patched\n")
    fmt.Printf("version of gorest, are we?\n")
    return
  }

  gorest.RegisterService(new(EitService))
  http.Handle("/", gorest.Handle())
  http.ListenAndServe(":8099", nil)
}


func ourMarshal(v interface{}) ([]byte, error) {
  // if v is a struct, we want the output's keys to be non-capitalized.
  // make that happen, by way of converting to a map[string]
  typ := reflect.TypeOf(v)
  val := reflect.ValueOf(v)
  if typ.Kind() == reflect.Struct {
    m := make(map[string]interface{})
    for i := 0; i < typ.NumField(); i++ {
      fieldname := typ.Field(i).Name
      lowercase_fieldname := strings.Join(
        []string{strings.ToLower(string(fieldname[0])), fieldname[1:]},
        "")
      if fieldname != lowercase_fieldname {
        fieldval := val.Field(i).Interface()
        m[lowercase_fieldname] = fieldval
      }
    }
    return json.Marshal(m)
  }

  return json.Marshal(v)
}
func ourUnmarshal(data []byte, v interface{}) error {
  return json.Unmarshal(data, v)
}
func ourMarshaller() *gorest.Marshaller {
  return &gorest.Marshaller{ourMarshal, ourUnmarshal}
}


type EitService struct {
  gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`

  gamesShow gorest.EndPoint `method:"GET" path:"/games/{gameId:string}" output:"Game"`
  gamesCreate gorest.EndPoint `method:"POST" path:"/games" postdata:"Game"`
}


type Game struct {
  TourneyId string
  UserId string
  GameId string

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



func(serv EitService) GamesShow(gameId string) (game Game) {
  game, found := games[gameId]
  if !found {
    game = Game{
      GameId: gameId, UserId: "eit_krog", Name: "test", Turns: 1, Points: 22,
      Deaths: 1, Maxlvl: 1, Death: "choked on dicks"}
  }
  return
}

func(serv EitService) GamesCreate(game Game) {
  ingest_game(game)
  serv.ResponseBuilder().SetResponseCode(201)
  return
}

func ingest_game(game Game) {
  games[game.GameId] = game
  return
}
