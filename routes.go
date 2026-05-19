package main

import (
	"net/http"
	handlers "serv-test/api"
	"serv-test/config"
	"serv-test/internal/models"
	"serv-test/middlewares"
	runServer "serv-test/server"

	"github.com/justinas/alice"
)

func RouteHandlers() http.Handler {

	mux := http.NewServeMux()

	SessionMan := alice.New(config.App.SessionManager.LoadAndSave)

	// initialize models with the configured DB so handlers receive ready model instances
	postModel := &models.PostModel{DB: config.App.DB}
	userModel := &models.UserModel{DB: config.App.DB}

	mux.Handle("GET /web/", http.StripPrefix("/web", runServer.FileServer()))

	mux.Handle("GET /{$}", SessionMan.ThenFunc(handlers.IndexHandler))
	mux.Handle("GET /signIn", SessionMan.ThenFunc(handlers.SignInHandler))
	mux.Handle("GET /about", SessionMan.ThenFunc(handlers.AboutHandler))
	mux.Handle("GET /create-account", SessionMan.ThenFunc(handlers.CreateAccountHandler))
	mux.Handle("GET /welcome", SessionMan.ThenFunc(handlers.WelcomeHandler))
	mux.Handle("GET /home", SessionMan.ThenFunc(handlers.HomeHandler))

	mux.Handle("POST /create-account", SessionMan.ThenFunc(handlers.CreateUser(userModel)))
	mux.Handle("POST /signIn", SessionMan.ThenFunc(handlers.PostSignInHandler(&models.Post{}, userModel)))
	mux.Handle("POST /create-post", SessionMan.ThenFunc(handlers.PostHandler(&models.Post{}, postModel)))

	PanicLogHeaders := alice.New(middlewares.PanicRecover, middlewares.LogRequest, middlewares.CommonHeaders)

	return PanicLogHeaders.Then(mux)

}
