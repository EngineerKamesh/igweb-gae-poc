package handlers

import (
	"net/http"

	"common"

	"github.com/isomorphicgo/isokit"
)

func LocalStorageDemoHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env.TemplateSet.Render("localstorage_example_page", &isokit.RenderParams{Writer: w, Data: nil})
	})
}
