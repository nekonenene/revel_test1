package controllers

import (
	"strconv"
	"github.com/revel/revel"
	"revel-test1/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	db := models.Gorm
	println("Gorm works? : " + strconv.FormatBool(db != nil))

	newMemo := models.Memo{Text: "メモをどんどん作るよ！"}
	db.NewRecord(newMemo)
	db.Create(&newMemo)

	var memos []models.Memo
	db.Find(&memos)

	userName := "夏目さん"
	age := 14

	return c.Render(userName, age, memos)
}
