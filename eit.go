package main

import (
  "code.google.com/p/gorest"
  "net/http"
  "encoding/json"
  "reflect"
  "fmt"
  "strings"
)


func main() {
  // gotta register the marshaller before registering the service
  marshaller := ourMarshaller()
  gorest.RegisterMarshaller("application/json", marshaller)
  if gorest.GetMarshallerByMime("application/json") != marshaller {
    fmt.Printf("OH DIP, RegisterMarshaller() didn't work\n")
  }

  gorest.RegisterService(new(EitService))
  http.Handle("/", gorest.Handle())
  http.ListenAndServe(":8099", nil)
}


func ourMarshal(v interface{}) ([]byte, error) {
  fmt.Printf("entered ourMarshal\n")

  // if v is a struct, we want the output's keys to be non-capitalized.
  // make that happen, by way of converting to a map[string]
  typ := reflect.TypeOf(v)
  val := reflect.ValueOf(v)
  if typ.Kind() == reflect.Struct {
    m := make(map[string]interface{})
    for i := 0; i < typ.NumField(); i++ {
      fieldname := typ.Field(i).Name
      lowercase_fieldname := strings.Join([]string{strings.ToLower(string(fieldname[0])),
        fieldname[1:]}, "")
      m[lowercase_fieldname] = reflect.ValueOf(val.FieldByName(fieldname))
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
  TourneyId string "tourney_id"
  UserId string "user_id"
  GameId string "game_id"

  version string
  points int
  deathdnum int
  deathlev int
  maxlvl int
  hp int
  maxhp int
  deaths int
  deathdate int // yyyymmdd
  birthdate int // yyyymmdd
  uid int
  role string
  race string
  gender string
  align string
  name string
  death string
  conduct int // bitfield
  turns int
  achieve int // bitfield
  realtime int // seconds
  starttime int // seconds since epoch
  endtime int // seconds since epoch
  gender0 string
  align0 string
}



func(serv EitService) GamesShow(gameId string) (game Game) {
  game = Game{
    GameId: gameId, name: "test", turns: 1, points: 22,
    deaths: 1, maxlvl: 1, death: "choked on dicks"}
  return
}

func(serv EitService) GamesCreate(game Game) {
  serv.ResponseBuilder().SetResponseCode(201)
  return
}
