package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joaomarcosbc/api-fc/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 10.00)
	productDb := NewProduct(db)

	err = productDb.Create(product)

	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.00)
	db.Create(product)

	productDb := NewProduct(db)

	productFound, err := productDb.FindByID(string(product.ID.String()))

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.00)
	db.Create(product)

	productDb := NewProduct(db)
	product.Name = "Product 2"
	product.Price = 15.00

	err = productDb.Update(product)
	assert.Nil(t, err)

	var productFound *entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, productFound.Name, "Product 2")
	assert.Equal(t, productFound.Price, 15.00)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.00)
	db.Create(product)

	productDb := NewProduct(db)

	var productFound *entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)

	err = productDb.Delete(product.ID.String())
	assert.Nil(t, err)

	err = productDb.Delete(product.ID.String())
	assert.NotNil(t, err)

	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.NotNil(t, err)
}
