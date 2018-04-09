package templatedata

import "shared/models"

type Products struct {
	PageTitle string
	Products  []*models.Product
}
