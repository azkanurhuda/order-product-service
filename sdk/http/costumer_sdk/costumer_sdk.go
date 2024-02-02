package costumer_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
)

type CostumerResponse struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CostumerSDK struct {
	Viper *viper.Viper
	Log   *logrus.Logger
}

func (c CostumerSDK) GetCustomerByID(ctx context.Context, ID int) (*CostumerResponse, error) {
	consumerBaseUrl := c.Viper.GetString("sdk.costumer-service-url")
	url := fmt.Sprintf("%s/api/costumer/%d", consumerBaseUrl, ID)

	// Create a request with the provided context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		c.Log.Errorf("Error creating HTTP request: %v", err)
		return nil, err
	}

	// Send the request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		c.Log.Errorf("Error making the request: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.Log.Errorf("Error reading the response body: %v", err)
		return nil, err
	}

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		c.Log.Errorf("Error: Status Code %d\n", response.StatusCode)
		err = fmt.Errorf("Error: Status Code %d\n", response.StatusCode)
		return nil, err
	}

	// Unmarshal the JSON response
	var responseData map[string]CostumerResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		c.Log.Errorf("Error unmarshalling JSON: %v", err)
		return nil, err
	}

	// Extract the CostumerResponse from the "data" field
	data, ok := responseData["data"]
	if !ok {
		c.Log.Errorf("Missing 'data' field in the JSON response")
		return nil, fmt.Errorf("Missing 'data' field in the JSON response")
	}

	spew.Dump(data)

	return &data, nil
}

func NewCostumerSDK(viper *viper.Viper, log *logrus.Logger) SDK {
	return &CostumerSDK{
		Viper: viper,
		Log:   log,
	}
}
