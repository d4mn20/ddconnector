package ddapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/ddapi/types"
	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/utils"
)

func (c *Client) GetTests(pageURL *string) (types.RespTests, error) {
	url := BaseURL + "/tests"
	authToken := c.token
	allTests := types.RespTests{}

	for {
		if pageURL != nil {
			url = *pageURL
		}

		if val, ok := c.cache.Get(url); ok {
			tests := types.RespTests{}
			err := json.Unmarshal(val, &tests)
			if err != nil {
				return types.RespTests{}, err
			}
			allTests.Results = append(allTests.Results, tests.Results...)

			if tests.Next == "" {
				break
			}

			pageURL = &tests.Next
			continue
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return types.RespTests{}, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Token "+authToken)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return types.RespTests{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return types.RespTests{}, err
		}

		if resp.StatusCode != http.StatusOK {
			return types.RespTests{}, fmt.Errorf("failed to get tests: %s", utils.GetErrorMessageFromBody(body))
		}

		tests := types.RespTests{}
		err = json.Unmarshal(body, &tests)
		if err != nil {
			return types.RespTests{}, err
		}

		c.cache.Add(url, body)
		allTests.Results = append(allTests.Results, tests.Results...)

		if tests.Next == "" {
			break
		}

		pageURL = &tests.Next
	}

	return allTests, nil
}

func (c *Client) GetTest(id int) (types.Test, error) {
	url := BaseURL + "/tests/" + fmt.Sprintf("%v", id)
	authToken := c.token

	if val, ok := c.cache.Get(url); ok {
		test := types.Test{}
		err := json.Unmarshal(val, &test)
		if err != nil {
			return types.Test{}, err
		}

		return test, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Test{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.Test{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Test{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return types.Test{}, fmt.Errorf("failed to get test: %s", utils.GetErrorMessageFromBody(body))
	}

	test := types.Test{}
	err = json.Unmarshal(body, &test)
	if err != nil {
		return types.Test{}, fmt.Errorf("failed to unmarshal test with ID %d: %w", id, err)
	}

	c.cache.Add(url, body)

	return test, nil
}

// func (c *Client) CreateTest(test types.Test) (types.Test, error) {
// 	url := BaseURL + "/tests/"
// 	authToken := c.token

// 	payload, err := json.Marshal(test)
// 	if err != nil {
// 		return types.Test{}, err
// 	}

// 	fmt.Printf("[PAYLOAD TEST]: %v\n", payload)

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
// 	if err != nil {
// 		return types.Test{}, err
// 	}
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("Authorization", "Token "+authToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return types.Test{}, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return types.Test{}, err
// 	}

// 	if resp.StatusCode != http.StatusCreated {
// 		return types.Test{}, fmt.Errorf("failed to create test: %s", utils.GetErrorMessageFromBody(body))
// 	}

// 	respTest := types.Test{}
// 	err = json.Unmarshal(body, &respTest)
// 	if err != nil {
// 		return types.Test{}, err
// 	}

// 	fmt.Printf("[RESPONSE BODY TEST]: %v\n", respTest)

// 	c.cache.Delete(BaseURL + "/tests")
// 	c.cache.Delete(BaseURL + "/tests/" + strconv.Itoa(test.ID))

// 	return respTest, nil
// }

func (c *Client) TestExists(test types.Test, engagement types.Engagement) (bool, error) {
	tests, err := c.GetTests(nil)
	if err != nil {
		return false, err
	}

	for _, t := range tests.Results {
		if t.ScanType == test.ScanType && t.Engagement == engagement.ID {
			return true, nil
		}
	}

	return false, nil
}
