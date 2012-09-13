package controllers

import (
	"eit_go/app/models"
	"github.com/robfig/revel"
)

type Games struct {
	*rev.Controller
}

func (c Games) Show(id string) rev.Result {
	return c.RenderJson(models.GameById(id))
}

func (c Games) Create(game models.Game) rev.Result {
	return c.RenderJson(models.Game{})
}

func (c Games) Update(game models.Game) rev.Result {
	return c.RenderJson(models.Game{})
}
