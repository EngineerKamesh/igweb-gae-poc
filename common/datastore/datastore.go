package datastore

import (
	"context"
	"errors"

	"shared/models"
)

type Datastore interface {
	CreateGopherTeam(team []*models.Gopher) error
	GetGopherTeam() []*models.Gopher
	CreateProduct(product *models.Product) error
	CreateProductRegistry(products []string) error
	GetProducts() []*models.Product
	GetProductDetail(productTitle string) *models.Product
	GetProductsInShoppingCart(cart *models.ShoppingCart) []*models.Product
	CreateContactRequest(contactRrequest *models.ContactRequest) error
	Close()
}

const (
	REDIS = iota
	GCD
)

func NewDatastore(datastoreType int, dbConnectionString string, ctx *context.Context) (Datastore, error) {

	switch datastoreType {

	case REDIS:
		return NewRedisDatastore(dbConnectionString)

	case GCD:
		return NewGoogleCloudDatastore(dbConnectionString, ctx)

	default:
		return nil, errors.New("Unrecognized Datastore!")

	}
}
