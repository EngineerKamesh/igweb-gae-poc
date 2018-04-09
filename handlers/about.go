package handlers

import (
	"net/http"

	"common"
	"shared/templatedata"

	"github.com/isomorphicgo/isokit"
)

func AboutHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gophers := env.DB.GetGopherTeam()
		templateData := templatedata.About{PageTitle: "About", Gophers: gophers}
		env.TemplateSet.Render("about_page", &isokit.RenderParams{Writer: w, Data: templateData})
	})
}
