package ddapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/ddapi/types"
	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/utils"
)

// GetProducts retrieves all products, handling pagination
func (c *Client) GetProducts(pageURL *string) (types.RespProducts, error) {
	url := BaseURL + "/products"
	authToken := c.token
	allProducts := types.RespProducts{}

	for {
		if pageURL != nil {
			url = *pageURL
		}

		if val, ok := c.cache.Get(url); ok {
			products := types.RespProducts{}
			err := json.Unmarshal(val, &products)
			if err != nil {
				return types.RespProducts{}, err
			}
			allProducts.Results = append(allProducts.Results, products.Results...)

			if products.Next == "" {
				break
			}

			pageURL = &products.Next
			continue
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return types.RespProducts{}, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Token "+authToken)

		// fmt.Printf("[ HEADER: %v ]", req.Header)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return types.RespProducts{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return types.RespProducts{}, err
		}

		if resp.StatusCode != http.StatusOK {
			return types.RespProducts{}, fmt.Errorf("failed to get products: %s", utils.GetErrorMessageFromBody(body))
		}

		products := types.RespProducts{}
		err = json.Unmarshal(body, &products)
		if err != nil {
			return types.RespProducts{}, err
		}

		c.cache.Add(url, body)
		allProducts.Results = append(allProducts.Results, products.Results...)

		if products.Next == "" {
			break
		}

		pageURL = &products.Next
	}

	return allProducts, nil
}

// GetProduct retrieves a specific product by ID
func (c *Client) GetProduct(id int) (types.Product, error) {
	url := BaseURL + "/products/" + strconv.Itoa(id)
	authToken := c.token

	if val, ok := c.cache.Get(url); ok {
		product := types.Product{}
		err := json.Unmarshal(val, &product)
		if err != nil {
			return types.Product{}, err
		}

		return product, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Product{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.Product{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Product{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return types.Product{}, fmt.Errorf("failed to get product: %s", utils.GetErrorMessageFromBody(body))
	}

	product := types.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		return types.Product{}, errors.New("failed to unmarshal product with ID " + strconv.Itoa(id))
	}

	c.cache.Add(url, body)

	return product, nil
}

// CreateProduct creates a new product
func (c *Client) CreateProduct(prod types.Product) (types.Product, error) {
	url := BaseURL + "/products/"
	authToken := c.token

	payload := types.Product{
		Tags:             []string{"DDCONN"},
		Name:             prod.Name,
		Description:      "Default - Projeto não avaliado",
		ProductManager:   1,
		TechnicalContact: 1,
		TeamManager:      1,
		ProdType:         16,
		SLAConfiguration: 1,
		Regulations:      []int{},
	}

	product, err := json.Marshal(payload)
	if err != nil {
		return types.Product{}, err
	}

	// fmt.Printf("[PAYLOAD PROJECT]: %v\n", payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(product))
	if err != nil {
		return types.Product{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// fmt.Printf("[PRODUCT REQUEST]\n%v\n", req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.Product{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Product{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return types.Product{}, fmt.Errorf("failed to create product: %s", utils.GetErrorMessageFromBody(body))
	}

	respProduct := types.Product{}
	err = json.Unmarshal(body, &respProduct)
	if err != nil {
		return types.Product{}, err
	}

	// fmt.Printf("[RESPONSE BODY PROJECT]: %v\n", respProduct)

	c.cache.Delete(BaseURL + "/products")
	c.cache.Delete(BaseURL + "/products/" + strconv.Itoa(respProduct.ID))

	return respProduct, nil
}

// ProductExists checks if a product already exists
func (c *Client) ProductExists(product types.Product) (bool, error) {
	products, err := c.GetProducts(nil)
	if err != nil {
		return false, err
	}

	for _, p := range products.Results {
		if p.Name == product.Name {
			return true, nil
		}
	}

	return false, nil
}
