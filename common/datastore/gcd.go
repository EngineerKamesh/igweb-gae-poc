package datastore

import (
	"context"
	"log"
	"sort"
	"strings"
	"time"

	"shared/models"

	gcd "cloud.google.com/go/datastore"
)

type GopherTeam struct {
	Gophers []models.Gopher
}

type ProductRegistry struct {
	Products []string
}

type ProductWrapper struct {
	Product models.Product
}

type ContactRequestWrapper struct {
	ContactRequest models.ContactRequest
}

type GoogleCloudDatastore struct {
	ctx   *context.Context
	store *gcd.Client
}

func NewGoogleCloudDatastore(address string, ctx *context.Context) (*GoogleCloudDatastore, error) {

	datastoreClient, err := gcd.NewClient(*ctx, address)
	if err != nil {
		return nil, err
	}
	return &GoogleCloudDatastore{
		ctx:   ctx,
		store: datastoreClient,
	}, nil
}

func (g *GoogleCloudDatastore) GetProducts() []*models.Product {

	registryKey := "product-registry"
	gv := ProductRegistry{}
	k := gcd.NameKey("ProductRegistry", registryKey, nil)
	err := g.store.Get(*g.ctx, k, &gv)

	if err == gcd.ErrNoSuchEntity {
		return nil
	} else if err != nil {
		log.Println("Encountered error: ", err)
		return nil
	}

	var productKeys []string = gv.Products

	products := make(models.Products, 0)

	for i := 0; i < len(productKeys); i++ {

		productTitle := strings.Replace(productKeys[i], "/product-detail/", "", -1)
		product := g.GetProductDetail(productTitle)
		products = append(products, product)

	}
	sort.Sort(products)
	return products
}

func (g *GoogleCloudDatastore) GetProductDetail(productTitle string) *models.Product {

	productKey := "/product-detail/" + productTitle
	gv := ProductWrapper{}
	k := gcd.NameKey("ProductWrapper", productKey, nil)

	err := g.store.Get(*g.ctx, k, &gv)
	if err == gcd.ErrNoSuchEntity {
		return nil
	} else if err != nil {
		log.Println("Encountered error: ", err)
		return nil
	}

	return &gv.Product
}

func (g *GoogleCloudDatastore) GenerateProductsMap(products []*models.Product) map[string]*models.Product {

	productsMap := make(map[string]*models.Product)
	for i := 0; i < len(products); i++ {
		productsMap[products[i].SKU] = products[i]
	}

	return productsMap
}

func (g *GoogleCloudDatastore) GetProductsInShoppingCart(cart *models.ShoppingCart) []*models.Product {

	products := g.GetProducts()
	productsMap := g.GenerateProductsMap(products)

	result := make(models.Products, 0)
	for _, v := range cart.Items {
		product := &models.Product{}
		product = productsMap[v.ProductSKU]
		product.Quantity = v.Quantity
		result = append(result, product)
	}
	sort.Sort(result)
	return result

}

func (g *GoogleCloudDatastore) CreateProduct(product *models.Product) error {

	theproduct := *product

	gv := ProductWrapper{Product: theproduct}

	k := gcd.NameKey("ProductWrapper", product.Route, nil)
	_, err1 := g.store.Put(*g.ctx, k, &gv)

	if err1 != nil {
		return err1
	}

	return nil
}

func (g *GoogleCloudDatastore) CreateProductRegistry(products []string) error {

	gv := ProductRegistry{Products: products}
	k := gcd.NameKey("ProductRegistry", "product-registry", nil)
	_, err1 := g.store.Put(*g.ctx, k, &gv)
	if err1 != nil {
		return err1
	}
	return nil
}

func (g *GoogleCloudDatastore) CreateGopherTeam(team []*models.Gopher) error {

	theteam := make([]models.Gopher, len(team))

	for i, v := range team {
		theteam[i] = *v
	}

	gv := GopherTeam{Gophers: theteam}

	k := gcd.NameKey("GopherTeam", "gopher-team", nil)
	_, err1 := g.store.Put(*g.ctx, k, &gv)
	//log.Printf("put result: ", p)
	//log.Printf("gv result: ", *team[0])

	if err1 != nil {
		log.Printf("got this error when attempting to put the gopher team: ", err1)
		return err1

	}

	return nil

}

func (g *GoogleCloudDatastore) GetGopherTeam() []*models.Gopher {

	log.Printf("reached get gopher team")
	team := GopherTeam{}
	k := gcd.NameKey("GopherTeam", "gopher-team", nil)

	err := g.store.Get(*g.ctx, k, &team)

	log.Printf("team: %+v", team.Gophers)

	if err != nil {
		log.Printf("got this error when attempting to get the gopher team: ", err)
		return nil
	}

	result := make([]*models.Gopher, len(team.Gophers))

	for i, _ := range team.Gophers {
		result[i] = &team.Gophers[i]
	}

	return result

}

func (g *GoogleCloudDatastore) CreateContactRequest(contactRequest *models.ContactRequest) error {

	now := time.Now()
	nowFormatted := now.Format(time.RFC822Z)

	contactRequestKey := "contact-request|" + contactRequest.Email + "|" + nowFormatted
	gv := ContactRequestWrapper{ContactRequest: *contactRequest}

	k := gcd.NameKey("ContactRequestWrapper", contactRequestKey, nil)
	_, err1 := g.store.Put(*g.ctx, k, &gv)

	if err1 != nil {
		return err1
	}
	return nil
}

func (g *GoogleCloudDatastore) Close() {

	g.store.Close()
}
