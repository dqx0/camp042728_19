package main

import (
	"github.com/dqx0/camp042728_19/db"
	"github.com/dqx0/camp042728_19/handler"
	"github.com/dqx0/camp042728_19/repository"
	"github.com/dqx0/camp042728_19/router"
	"github.com/dqx0/camp042728_19/usecase"
)

func main() {
	//
	db := db.NewDB()
	repository := repository.NewBaseRepository(db)
	h := handler.NewBaseHandler(usecase.NewBaseUsecase(repository))
	r := router.NewRouter(h)
	r.Run(":8080")
}
