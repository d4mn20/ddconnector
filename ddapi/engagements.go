package ddapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/ddapi/types"
	"dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/utils"
)

// GetEngagements retrieves all engagements, handling pagination
func (c *Client) GetEngagements(pageURL *string) (types.RespEngagements, error) {
	url := BaseURL + "/engagements"
	authToken := c.token
	allEngagements := types.RespEngagements{}

	for {
		if pageURL != nil {
			url = *pageURL
		}

		if val, ok := c.cache.Get(url); ok {
			engagements := types.RespEngagements{}
			err := json.Unmarshal(val, &engagements)
			if err != nil {
				return types.RespEngagements{}, err
			}
			allEngagements.Results = append(allEngagements.Results, engagements.Results...)

			if engagements.Next == "" {
				break
			}

			pageURL = &engagements.Next
			continue
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return types.RespEngagements{}, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Token "+authToken)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return types.RespEngagements{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return types.RespEngagements{}, err
		}

		if resp.StatusCode != http.StatusOK {
			return types.RespEngagements{}, fmt.Errorf("failed to get engagements: %s", utils.GetErrorMessageFromBody(body))
		}

		engagements := types.RespEngagements{}
		err = json.Unmarshal(body, &engagements)
		if err != nil {
			return types.RespEngagements{}, err
		}

		c.cache.Add(url, body)
		allEngagements.Results = append(allEngagements.Results, engagements.Results...)

		if engagements.Next == "" {
			break
		}

		pageURL = &engagements.Next
	}

	return allEngagements, nil
}

// GetEngagement retrieves a specific engagement by ID
func (c *Client) GetEngagement(id int) (types.Engagement, error) {
	url := BaseURL + "/engagements/" + strconv.Itoa(id)
	authToken := c.token

	if val, ok := c.cache.Get(url); ok {
		engagement := types.Engagement{}
		err := json.Unmarshal(val, &engagement)
		if err != nil {
			return types.Engagement{}, err
		}

		return engagement, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Engagement{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.Engagement{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Engagement{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return types.Engagement{}, fmt.Errorf("failed to get engagement: %s", utils.GetErrorMessageFromBody(body))
	}

	engagement := types.Engagement{}
	err = json.Unmarshal(body, &engagement)
	if err != nil {
		return types.Engagement{}, fmt.Errorf("failed to unmarshal engagement with ID %d: %w", id, err)
	}

	c.cache.Add(url, body)

	return engagement, nil
}

// CreateEngagement creates a new engagement
func (c *Client) CreateEngagement(engage types.Engagement) (types.Engagement, error) {
	url := BaseURL + "/engagements/"
	authToken := c.token

	interval := 30 * 24 * time.Hour
	targetEnd := time.Now().Add(interval)

	payload := types.Engagement{
		Name:                      engage.Name,
		EngagementType:            "CI/CD",
		Product:                   engage.Product,
		FirstContacted:            time.Now().Format("2006-01-02"),
		TargetStart:               time.Now().Format("2006-01-02"),
		TargetEnd:                 targetEnd.Format("2006-01-02"),
		BranchTag:                 engage.BranchTag,
		SourceCodeManagementURI:   engage.SourceCodeManagementURI,
		DeduplicationOnEngagement: true,
	}

	engagement, err := json.Marshal(payload)
	if err != nil {
		return types.Engagement{}, err
	}

	fmt.Printf("[PAYLOAD ENGAGEMENT]: %v\n", payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(engagement))
	if err != nil {
		return types.Engagement{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+authToken)
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("[ENGAGEMENT REQUEST]\n%v\n", req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.Engagement{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Engagement{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return types.Engagement{}, fmt.Errorf("failed to create engagement: %s", utils.GetErrorMessageFromBody(body))
	}

	respEngagement := types.Engagement{}
	err = json.Unmarshal(body, &respEngagement)
	if err != nil {
		return types.Engagement{}, err
	}

	fmt.Printf("[RESPONSE BODY ENGAGEMENT]: %v\n", respEngagement)

	c.cache.Delete(BaseURL + "/engagements")
	c.cache.Delete(BaseURL + "/engagements/" + strconv.Itoa(respEngagement.ID))

	return respEngagement, nil
}

// EngagementExists checks if an engagement already exists
func (c *Client) EngagementExists(engagement types.Engagement, product types.Product) (bool, error) {
	url := BaseURL + "/engagements"
	authToken := c.token

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to check if engagement exists: %s", utils.GetErrorMessageFromBody(body))
	}

	engagements := types.RespEngagements{}
	err = json.Unmarshal(body, &engagements)
	if err != nil {
		return false, err
	}

	for _, e := range engagements.Results {
		if e.Product == product.ID && e.Name == engagement.Name {
			return true, nil
		}
	}

	return false, nil
}
