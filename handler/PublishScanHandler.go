package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/ddapi"
	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/ddapi/types"
	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/schemas"
	"github.com/gin-gonic/gin"
)

const (
	uploadDir    = "./uploads/"
	timeout      = 5 * time.Minute
	shortTimeout = 5 * time.Second
)

// PublishScanHandler handles the publishing of scan data
func PublishScanHandler(ctx *gin.Context) {
	request := PublishScan{}

	if err := ctx.ShouldBind(&request); err != nil {
		logger.Errorf("PublishScanHandler: multipart/form-data binding error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02d-%02d-%02d-%02d-", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	path := uploadDir + formatted + request.File.Filename

	if err := ctx.SaveUploadedFile(request.File, path); err != nil {
		logger.Errorf("PublishScanHandler: upload error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "File upload error: "+err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("PublishScanHandler: validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	scan := schemas.Scan{
		Product:    request.Product,
		Engagement: request.Engagement,
		Test:       request.Test,
		Branch:     request.Branch,
		RepoUrl:    request.RepoUrl,
		Origin:     request.Origin,
		FilePath:   path,
	}

	if err := db.Create(&scan).Error; err != nil {
		logger.Errorf("PublishScanHandler: error publishing scan: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}

	ddClient := ddapi.NewClient(shortTimeout, timeout, true)

	productID, err := validateOrCreateProduct(ddClient, request.Product)
	if err != nil {
		logger.Errorf("PublishScanHandler: product validation error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Product validation error: "+err.Error())
		return
	}

	fmt.Printf("[CREATING ENGAGEMENT | PRODUCT ID: %v]\n", productID)
	engagementID, err := validateOrCreateEngagement(ddClient, productID, request.Engagement, request.Branch, request.RepoUrl)
	if err != nil {
		logger.Errorf("PublishScanHandler: engagement validation error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Engagement validation error: "+err.Error())
		return
	}

	testID, err := validateOrImportTest(ddClient, engagementID, path, request.Test, request.Product, request.Engagement)
	if err != nil {
		logger.Errorf("PublishScanHandler: test validation error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Test validation error: "+err.Error())
		return
	}

	logger.Infof("PublishScanHandler: Test ID: %d", testID)
	sendCreationSuccess(ctx, "publish-scan", scan)
}

// validateOrCreateProduct validates if a product exists or creates a new one
func validateOrCreateProduct(client ddapi.Client, productName string) (int, error) {
	products, err := client.GetProducts(nil)
	if err != nil {
		return 0, fmt.Errorf("validateOrCreateProduct: error getting products from API: %v", err)
	}

	// Logging the products for debugging
	productsJSON, _ := json.Marshal(products)
	logger.Infof("validateOrCreateProduct: Retrieved products: %s", string(productsJSON))

	for _, product := range products.Results {
		fmt.Printf("Product.Name: %v\n", product.Name)
		if product.Name == productName {
			return product.ID, nil
		}
	}

	createdProduct, err := client.CreateProduct(types.Product{Name: productName})
	if err != nil {
		return 0, fmt.Errorf("validateOrCreateProduct: error creating product on API: %v", err)
	}

	return createdProduct.ID, nil
}

// validateOrCreateEngagement validates if an engagement exists or creates a new one
func validateOrCreateEngagement(client ddapi.Client, productID int, engagementName, branchTag string, repoUrl string) (int, error) {
	engagements, err := client.GetEngagements(nil)
	if err != nil {
		return 0, fmt.Errorf("validateOrCreateEngagement: error getting engagements from API: %v", err)
	}

	for _, engagement := range engagements.Results {
		if engagement.Product == productID && engagement.Name == engagementName && engagement.BranchTag == branchTag {
			return engagement.ID, nil
		}
	}

	createdEngagement, err := client.CreateEngagement(types.Engagement{Name: engagementName, BranchTag: branchTag, SourceCodeManagementURI: repoUrl, Product: productID})
	if err != nil {
		return 0, fmt.Errorf("validateOrCreateEngagement: error creating engagement on API: %v", err)
	}

	return createdEngagement.ID, nil
}

// validateOrImportTest validates if a test exists or imports a new one
func validateOrImportTest(client ddapi.Client, engagementID int, filePath, testName, productName, engagementName string) (int, error) {
	tests, err := client.GetTests(nil)
	if err != nil {
		return 0, fmt.Errorf("validateOrImportTest: error getting tests from API: %v", err)
	}

	for _, test := range tests.Results {
		if test.Engagement == engagementID && test.ScanType == testName {
			_, err := client.ReimportScan(test.ID, filePath, productName, engagementName, testName, true, true)
			if err != nil {
				return 0, fmt.Errorf("validateOrImportTest: error reimporting scan: %v", err)
			}
			return test.ID, nil
		}
	}

	_, err = client.ImportScan(engagementID, filePath, testName)
	if err != nil {
		return 0, fmt.Errorf("validateOrImportTest: error importing scan: %v", err)
	}

	return 0, nil
}
