package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func getProductById(t *testing.T, id int32) Products {
	availableProduct, err := testQueries.GetProduct(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, availableProduct)

	require.Equal(t, availableProduct.ProductID, id)

	return availableProduct
}

func TestGetProduct(t *testing.T) {
	testProductId := int32(3)
	getProductById(t, testProductId)
}

func TestUpdateProduct(t *testing.T) {

	testProductId := int32(3)
	testAmount := int32(5)
	product := getProductById(t, testProductId)

	updatedParams := UpdateProductParams{
		ProductID: product.ProductID,
		Amount:    product.Amount + testAmount,
	}

	updatedProduct, err := testQueries.UpdateProduct(context.Background(), updatedParams)

	require.NoError(t, err)
	require.NotEmpty(t, updatedProduct)

	require.Equal(t, updatedProduct.ProductID, testProductId)
	require.Equal(t, updatedProduct.Amount, product.Amount+testAmount)
	require.Equal(t, updatedProduct.Price, product.Price)
}
