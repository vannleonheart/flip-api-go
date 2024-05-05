package flip

import (
	"encoding/base64"
	"fmt"
	"github.com/vannleonheart/goutil"
	"time"
)

func New(config *Config) *Client {
	return &Client{Config: config}
}

func (c *Client) generateBasicAuthorizationString() string {
	str := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", c.Config.ApiSecretKey)))

	return fmt.Sprintf("Basic %s", str)
}

func (c *Client) sendGetRequest(targetUrl string, requestParams interface{}, requestHeaders *map[string]string, result interface{}) (*[]byte, error) {
	useRequestHeaders := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": c.generateBasicAuthorizationString(),
		"X-TIMESTAMP":   time.Now().Format("2006-01-02T15:04:05-0700"),
	}

	if requestHeaders != nil {
		for k, v := range *requestHeaders {
			useRequestHeaders[k] = v
		}
	}

	byteResponseBody, err := goutil.SendHttpGet(targetUrl, requestParams, &useRequestHeaders, result)

	if err != nil {
		return nil, err
	}

	return byteResponseBody, nil
}

func (c *Client) sendPostRequest(targetUrl string, requestData interface{}, requestHeaders *map[string]string, result interface{}) (*[]byte, error) {
	useRequestHeaders := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": c.generateBasicAuthorizationString(),
		"X-TIMESTAMP":   time.Now().Format("2006-01-02T15:04:05-0700"),
	}

	if requestHeaders != nil {
		for k, v := range *requestHeaders {
			useRequestHeaders[k] = v
		}
	}

	byteResponseBody, err := goutil.SendHttpPost(targetUrl, requestData, &useRequestHeaders, result)

	if err != nil {
		return nil, err
	}

	return byteResponseBody, nil
}

func (c *Client) GetBalance() (*Balance, error) {
	var result Balance

	targetUrl := fmt.Sprintf("%s%s", c.Config.BaseUrl, UrlGetBalance)

	if _, err := c.sendGetRequest(targetUrl, nil, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBankInfo(codeBank string) (*[]BankInfo, error) {
	var result []BankInfo

	url := fmt.Sprintf("%s%s", c.Config.BaseUrl, UrlGetBankInfo)

	var params *map[string]string

	if len(codeBank) > 0 {
		params = &map[string]string{
			"code": codeBank,
		}
	}

	if _, err := c.sendGetRequest(url, params, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) IsMaintenance() (*Maintenance, error) {
	var result Maintenance

	url := fmt.Sprintf("%s%s", c.Config.BaseUrl, UrlGetIsMaintenance)

	if _, err := c.sendGetRequest(url, nil, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BankAccountInquiry(bankCode, accountNumber string, inquiryKey *string) (*BankAccountInquiry, error) {
	var result BankAccountInquiry

	url := fmt.Sprintf("%s%s", c.Config.BaseUrl, UrlGetBankAccountInquiry)

	requestData := BankAccountInquiryRequest{
		AccountNumber: accountNumber,
		BankCode:      bankCode,
		InquiryKey:    inquiryKey,
	}

	if _, err := c.sendPostRequest(url, &requestData, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
