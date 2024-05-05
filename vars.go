package flip

import "encoding/json"

type Client struct {
	Config *Config
}

type Config struct {
	BaseUrl         string `json:"base_url"`
	ApiSecretKey    string `json:"api_secret_key"`
	ValidationToken string `json:"validation_token"`
}

type Maintenance struct {
	Maintenance bool `json:"maintenance"`
}

type Balance struct {
	Balance json.Number `json:"balance"`
}

type BankInfo struct {
	BankCode string `json:"bank_code"`
	Name     string `json:"name"`
	Fee      int64  `json:"fee"`
	Queue    int64  `json:"queue"`
	Status   string `json:"status"`
}

type BankAccountInquiryRequest struct {
	AccountNumber string  `json:"account_number"`
	BankCode      string  `json:"bank_code"`
	InquiryKey    *string `json:"inquiry_key"`
}

type BankAccountInquiry struct {
	BankCode      string `json:"bank_code"`
	InquiryKey    string `json:"inquiry_key"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
	Status        string `json:"status"`
	*ErrorResponse
}

type DisbursementRequest struct {
	AccountNumber    string  `json:"account_number"`
	BankCode         string  `json:"bank_code"`
	Amount           int64   `json:"amount"`
	Remark           *string `json:"remark,omitempty"`
	RecipientCity    *int64  `json:"recipient_city,omitempty"`
	BeneficiaryEmail *string `json:"beneficiary_email,omitempty"`
}

type Disbursement struct {
	Id               json.Number `json:"id"`
	UserId           json.Number `json:"user_id"`
	Amount           int64       `json:"amount"`
	Status           string      `json:"status"`
	Reason           string      `json:"reason"`
	Timestamp        string      `json:"timestamp"`
	BankCode         string      `json:"bank_code"`
	AccountNumber    string      `json:"account_number"`
	RecipientName    string      `json:"recipient_name"`
	SenderBank       *string     `json:"sender_bank"`
	Remark           string      `json:"remark"`
	Receipt          string      `json:"receipt"`
	TimeServed       string      `json:"time_served"`
	BundleId         int64       `json:"bundle_id"`
	CompanyId        int64       `json:"company_id"`
	RecipientCity    int64       `json:"recipient_city"`
	CreatedFrom      string      `json:"created_from"`
	Direction        string      `json:"direction"`
	Sender           *Sender     `json:"sender,omitempty"`
	Fee              int64       `json:"fee"`
	BeneficiaryEmail string      `json:"beneficiary_email"`
	IdempotencyKey   string      `json:"idempotency_key"`
}

type Sender struct {
	SenderName           string `json:"sender_name"`
	PlaceOfBirth         int64  `json:"place_of_birth"`
	DateOfBirth          string `json:"date_of_birth"`
	Address              string `json:"address"`
	SenderIdentityType   string `json:"sender_identity_type"`
	SenderIdentityNumber string `json:"sender_identity_number"`
	SenderCountry        int64  `json:"sender_country"`
	Job                  string `json:"job"`
}

type ErrorResponse struct {
	Code   string              `json:"code"`
	Errors []ErrorResponseItem `json:"errors"`
}

type ErrorResponseItem struct {
	Attribute string `json:"attribute"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}
