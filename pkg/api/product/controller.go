package product

import (
	"encoding/json"
	"gfg/pkg/api/seller"
	"gfg/pkg/api/urlutil"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	LIST_PAGE_SIZE = 10
)

func NewController(repository *repository, sellerRepository *seller.Repository, provider seller.Provider) *controller {
	return &controller{
		repository:       repository,
		sellerRepository: sellerRepository,
		sellerProvider:   provider,
	}
}

type controller struct {
	repository       *repository
	sellerRepository *seller.Repository
	sellerProvider   seller.Provider
}

func (pc *controller) List(c *gin.Context) {
	request := &struct {
		Page int `form:"page,default=1"`
	}{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var responseValue interface{}
	var err error
	if c.FullPath() == "/api/v1/products" {
		products := make([]*Product, 0)
		err = pc.repository.list((request.Page-1)*LIST_PAGE_SIZE, LIST_PAGE_SIZE, &products)
		responseValue = products
	} else if c.FullPath() == "/api/v2/products" {
		products := make([]*ProductV2, 0)
		err = pc.repository.list((request.Page-1)*LIST_PAGE_SIZE, LIST_PAGE_SIZE, &products, location.Get(c))
		responseValue = products
	}

	if err != nil {
		log.Error().Err(err).Msg("Fail to query product list")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to query product list"})
		return
	}

	productsJson, err := json.Marshal(responseValue)

	if err != nil {
		log.Error().Err(err).Msg("Fail to marshal products")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to marshal products"})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", productsJson)
}

func (pc *controller) Get(c *gin.Context) {

	var product interface{}
	if c.FullPath() == "/api/v1/product" {
		product = new(Product)
	} else if c.FullPath() == "/api/v2/product" {
		product = new(ProductV2)
	}

	request := &struct {
		UUID string `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := pc.repository.findByUUID(request.UUID, product)
	if err != nil {
		log.Error().Err(err).Msg("Fail to query product by uuid")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to query product by uuid"})
		return
	}

	if productV2, ok := product.(*ProductV2); ok {
		productV2.Seller.Links.Self.Href = urlutil.BuildSelfReferenceURL(location.Get(c), "/sellers", productV2.Seller.UUID)
	}

	productJson, err := json.Marshal(product)

	if err != nil {
		log.Error().Err(err).Msg("Fail to marshal product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to marshal product"})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", productJson)
}

func (pc *controller) Post(c *gin.Context) {
	request := &struct {
		Name   string `form:"name"`
		Brand  string `form:"brand"`
		Stock  int    `form:"stock"`
		Seller string `form:"seller"`
	}{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seller, err := pc.sellerRepository.FindByUUID(request.Seller)

	if err != nil {
		log.Error().Err(err).Msg("Fail to query seller by UUID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to query seller by UUID"})
		return
	}

	if seller == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Seller is not found"})
		return
	}

	product := &Product{
		BaseProduct: BaseProduct{
			UUID:  uuid.New().String(),
			Name:  request.Name,
			Brand: request.Brand,
			Stock: request.Stock,
		},
		SellerUUID: seller.UUID,
	}

	err = pc.repository.insert(product)

	if err != nil {
		log.Error().Err(err).Msg("Fail to insert product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to insert product"})
		return
	}

	jsonData, err := json.Marshal(product)

	if err != nil {
		log.Error().Err(err).Msg("Fail to marshal product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to marshal product"})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}

func (pc *controller) Put(c *gin.Context) {
	queryRequest := &struct {
		UUID string `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(queryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := new(Product)
	err := pc.repository.findByUUID(queryRequest.UUID, product)
	if err != nil {
		log.Error().Err(err).Msg("Fail to query product by uuid")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to query product by uuid"})
		return
	}

	request := &struct {
		Name  string `form:"name"`
		Brand string `form:"brand"`
		Stock int    `form:"stock"`
	}{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldStock := product.Stock

	product.Name = request.Name
	product.Brand = request.Brand
	product.Stock = request.Stock

	err = pc.repository.update(product)

	if err != nil {
		log.Error().Err(err).Msg("Fail to insert product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to insert product"})
		return
	}

	if oldStock != product.Stock {
		s, err := pc.sellerRepository.FindByUUID(product.SellerUUID)

		if err != nil {
			log.Error().Err(err).Msg("Fail to query seller by UUID")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to query seller by UUID"})
			return
		}

		stockFeed := &seller.StockFeed{
			Email:       s.Email,
			Phone:       s.Phone,
			OldStock:    oldStock,
			NewStock:    product.Stock,
			ProductName: product.Name,
			SellerUUID:  s.UUID,
		}

		pc.sellerProvider.StockChanged(stockFeed)
	}

	jsonData, err := json.Marshal(product)

	if err != nil {
		log.Error().Err(err).Msg("Fail to marshal product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to marshal product"})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}

func (pc *controller) Delete(c *gin.Context) {
	request := &struct {
		UUID string `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// todo(marcel) don't fetch all to make a exits check
	product := new(Product)
	err := pc.repository.findByUUID(request.UUID, product)
	if err != nil {
		if err == ErrNotFound {
			// todo(marcel) If resource is not available then please return http.StatusNotFound
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product is not found"})
			return
		}
		log.Error().Err(err).Msg("Fail to query product by uuid")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to query product by uuid"})
		return
	}

	err = pc.repository.delete(product)
	if err != nil {
		log.Error().Err(err).Msg("Fail to delete product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
