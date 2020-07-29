package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"test_task/src/calculator/delivery"
	ucase "test_task/src/calculator/usecase"
	"test_task/src/utills"
)

func main() {
	r := mux.NewRouter()
	usecase := ucase.NewCalculatorUsecase(utills.Timeouts.ContextTimeout)
	delivery.NewCalculatorHandler(r, usecase)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         utills.ServerUrl,
		WriteTimeout: utills.Timeouts.WriteTimeout,
		ReadTimeout:  utills.Timeouts.ReadTimeout,
	}
	log.Info().Msgf("Server started at " + utills.ServerUrl)
	log.Error().Msgf(srv.ListenAndServe().Error())

}
