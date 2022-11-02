package main

import (
	"projects-monitor/src/database"
	"projects-monitor/src/handler"
	"projects-monitor/src/repository"
	"projects-monitor/src/service"
)

func main() {
	db := database.New()
	dbCon := db.Open()

	r := repository.New(dbCon)
	s := service.New(r)
	h := handler.New(s)

	gr := h.InitGinRouter()

	if err := gr.Run(); err != nil {
		panic(err)
	}
}
