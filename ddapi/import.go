package ddapi

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/ddapi/types"
)

func (c *Client) ImportScan(engagementID int, filePath, testTypeName string) (types.TestResult, error) {
	url := BaseURL + "/import-scan/"
	authToken := c.token

	file, err := os.Open(filePath)
	if err != nil {
		return types.TestResult{}, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return types.TestResult{}, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return types.TestResult{}, err
	}

	_ = writer.WriteField("scan_type", testTypeName)

	_ = writer.WriteField("engagement", strconv.Itoa(engagementID))

	err = writer.Close()
	if err != nil {
		return types.TestResult{}, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return types.TestResult{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+authToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.TestResult{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.TestResult{}, err
	}

	var result types.TestResult
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return types.TestResult{}, err
	}

	c.cache.Reset()

	return result, nil
}
