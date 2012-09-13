package controllers

import (
	"eit_go/app/models"
	"fmt"
	"github.com/robfig/revel"
	"io/ioutil"
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
		g.Save()
	}
	fmt.Printf("imported %d lines. ", len(models.GamesMap))
}

func init() {
	rev.RegisterPlugin(XlogPlugin{})
}
