package main

import (
	"html/template"
	"net/http"

	todoclient "github.com/craftlab-demo-tests/todoapi/todo"
	userclient "github.com/craftlab-demo-tests/todoapi/user"
	"github.com/gorilla/mux"
)

type app struct {
	templates  *template.Template
	todoclient *todoclient.APIClient
	userclient *userclient.APIClient
}

func newDefaultApp() (*app, error) {
	tmpl, err := template.New("all").ParseGlob("./templates/*.html")
	if err != nil {
		return nil, err
	}

	todoConfig := todoclient.NewConfiguration()
	userConfig := userclient.NewConfiguration()
	return &app{
		templates:  tmpl,
		todoclient: todoclient.NewAPIClient(todoConfig),
		userclient: userclient.NewAPIClient(userConfig),
	}, nil
}

func (a app) registerRoute(router *mux.Router) {
	router.HandleFunc("/user/{id}", a.HandleGetUser).Methods("GET")
}

func (a app) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	a.templates.ExecuteTemplate(w, "getuserid.html", nil)
}
