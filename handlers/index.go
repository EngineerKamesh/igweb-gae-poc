package handlers

import (
	"net/http"

	"common"
	"shared/templatedata"

	"github.com/isomorphicgo/isokit"
)

func IndexHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templateData := templatedata.Index{PageTitle: "IGWEB"}
		env.TemplateSet.Render("index_page", &isokit.RenderParams{Writer: w, Data: templateData})
	})
}
