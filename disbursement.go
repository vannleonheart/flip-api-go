package flip

import (
	"fmt"
)

func (c *Client) CreateDisbursement(data DisbursementRequest, key string) (*Disbursement, error) {
	var result Disbursement

	url := fmt.Sprintf("%s%s", c.Config.BaseUrl, UrlCreateDisbursement)

	requestHeaders := map[string]string{
		"idempotency-key": key,
	}

	if _, err := c.sendPostRequest(url, data, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDisbursementById(id string) (*Disbursement, error) {
	var result Disbursement

	url := fmt.Sprintf("%s%s", c.Config.BaseUrl, UrlGetDisbursement)

	params := &map[string]string{
		"id": id,
	}

	if _, err := c.sendGetRequest(url, params, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDisbursementByIdempotencyKey(key string) (*Disbursement, error) {
	var result Disbursement

	url := fmt.Sprintf("%s%s", c.Config.BaseUrl, UrlGetDisbursement)

	params := &map[string]string{
		"idempotency-key": key,
	}

	if _, err := c.sendGetRequest(url, params, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
