package main

import (
	"html/template"
	"net/http"
	"strconv"

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
	type getuserdata struct {
		username string
	}

	idstr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		a.templates.ExecuteTemplate(w, "badrequest.html", nil)
	}
	request := a.userclient.DefaultApi.UsersIdGet(r.Context(), int32(id))
	u, _, err := request.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.templates.ExecuteTemplate(w, "servererror.html", nil)
	}

	a.templates.ExecuteTemplate(w, "getuserid.html", getuserdata{
		username: u.GetName(),
	})
}
