package controllers

import (
	"assignment/views"
	"net/http"
)

// var store = sessions.NewCookieStore([]byte("mysession"))

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(tpl.HTMLtpl.Name())
		// fmt.Println(r.URL.Path)
		tpl.Execute(w, nil)
	}
}
