### Flip API
#### Installation
```go
go get -u github.com/vannleonheart/flip-api-go
```
#### Config
```go
flipConfig := flip.Config{
    BaseUrl:         "{flip_api_base_url}",
    ApiSecretKey:    "{your_flip_api_secret_key}",
    ValidationToken: "{your_flip_validation_token}",
}
```
#### Create Client
```go
flipClient := flip.New(&flipConfig)
```
#### Get Balance
```go
result, err := flipClient.GetBalance()

if err != nil {
    // handle error
}

fmt.Println(result.Balance)
```
#### Get All Bank Info
```go
result, err := flipClient.GetBankInfo("")

if err != nil {
    // handle error
}

fmt.Println(len(result))
```
#### Get Single Bank Info
```go
bankCode := "bca"

result, err := flipClient.GetBankInfo(bankCode)

if err != nil {
    // handle error
}

fmt.Println(len(result))
```
#### Is Maintenance
```go
result, err := flipClient.IsMaintenance()

if err != nil {
    // handle error
}

fmt.Println(result.Maintenance)
```
#### Bank Account Inquiry
```go
bankCode := "{flip_bank_code}"
bankAccountNo := "{target_bank_account_number}"

result, err := flipClient.BankAccountInquiry(bankCode, bankAccountNo, nil)

if err != nil {
    // handle error
}

fmt.Println(result.Status)
```
#### Create Disbursement
```go
idempotencyKey := "{your_idempotency_key}"

testAmount := 100000

requestData := DisbursementRequest {
    AccountNumber: "{target_bank_account_number}",
    BankCode: "{flip_bank_code}",
    Amount: testAmount,
    Remark: "",
    RecipientCity: "",
    BeneficiaryEmail: "",
}

result, err := flipClient.CreateDisbursement(requestData, idempotencyKey)

if err != nil {
    // handle error
}

fmt.Println(result.Id)
```
#### Get Disbursement By Id
```go
transactionId := "{flip_transaction_id}"

result, err := flipClient.GetDisbursementById(transactionId)

if err != nil {
    // handle error
}

fmt.Println(result.Status)
```
#### Get Disbursement By Idempotency Key
```go
idempotencyKey := "{your_idempotency_key}"

result, err := flipClient.GetDisbursementByIdempotencyKey(idempotencyKey)

if err != nil {
    // handle error
}

fmt.Println(result.Status)
```