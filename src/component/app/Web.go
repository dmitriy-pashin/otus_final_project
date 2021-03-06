package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"otus_final_project/src/component/config"
	"otus_final_project/src/component/web"
	"otus_final_project/src/controller"
	"time"

	"github.com/gorilla/mux"
)

const (
	DefaultWebAppPort = 9000
)

type Web struct {
	*Base
	port   int
	server *http.Server
	router *mux.Router
}

func NewWeb(conf *config.Config, port int) *Web {
	if port == 0 {
		port = DefaultWebAppPort
	}

	app := &Web{
		Base: NewBase(conf),
		port: port,
	}

	app.addRoutes()

	return app
}

func (app *Web) Run() {
	listenAddress := fmt.Sprintf(":%d", app.port)

	app.server = &http.Server{
		Addr:         listenAddress,
		Handler:      app.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Printf("\nStart listen addr %s\n", listenAddress)
	err := app.server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func (app *Web) Close() {
	defer app.Base.Close()

	err := app.server.Shutdown(context.TODO())
	if err != nil {
		log.Println(err)
	}
}

func (app *Web) Router() *mux.Router {
	return app.router
}

func (app *Web) addRoutes() {
	app.router = mux.NewRouter()
	r := app.router

	contBlacklist := app.Container().Get("app.controller.blacklist").(*controller.Blacklist)
	contWhitelist := app.Container().Get("app.controller.whitelist").(*controller.Whitelist)
	contLogin := app.Container().Get("app.controller.login").(*controller.Login)

	handleActionAddToBlacklist := func(writer http.ResponseWriter, request *http.Request) {
		web.HandleAPIAction(contBlacklist.ActionAdd, writer, request)
	}
	handleActionDeleteFromBlacklist := func(writer http.ResponseWriter, request *http.Request) {
		web.HandleAPIAction(contBlacklist.ActionDelete, writer, request)
	}
	handleActionAddToWhitelist := func(writer http.ResponseWriter, request *http.Request) {
		web.HandleAPIAction(contWhitelist.ActionAdd, writer, request)
	}
	handleActionDeleteFromWhitelist := func(writer http.ResponseWriter, request *http.Request) {
		web.HandleAPIAction(contWhitelist.ActionDelete, writer, request)
	}
	handleActionLogin := func(writer http.ResponseWriter, request *http.Request) {
		web.HandleAPIAction(contLogin.ActionLogin, writer, request)
	}

	r.HandleFunc("/blacklist", handleActionAddToBlacklist).
		Name("blacklist.add").
		Methods("POST")

	r.HandleFunc("/blacklist", handleActionDeleteFromBlacklist).
		Name("blacklist.delete").
		Methods("DELETE")

	r.HandleFunc("/whitelist", handleActionAddToWhitelist).
		Name("whitelist.add").
		Methods("POST")

	r.HandleFunc("/whitelist", handleActionDeleteFromWhitelist).
		Name("whitelist.delete").
		Methods("DELETE")

	r.HandleFunc("/login", handleActionLogin).
		Name("login").
		Methods("POST")
}
