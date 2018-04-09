package templatedata

import "shared/models"

type About struct {
	PageTitle string
	Gophers   []*models.Gopher
}
