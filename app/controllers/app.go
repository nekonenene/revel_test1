package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	userName := "夏目さん"
	age := 14

	return c.Render(userName, age)
}
