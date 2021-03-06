package handlers

import (
	"net/http"

	"common"
	"shared/templatedata"

	"github.com/isomorphicgo/isokit"
)

func ProductsHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products := env.DB.GetProducts()
		templateData := &templatedata.Products{PageTitle: "Products", Products: products}
		env.TemplateSet.Render("products_page", &isokit.RenderParams{Writer: w, Data: templateData})
	})
}
