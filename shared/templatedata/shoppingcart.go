package templatedata

import "shared/models"

type ShoppingCart struct {
	PageTitle string
	Products  []*models.Product
}
