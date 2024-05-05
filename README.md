### Flip API
#### Installation
```go
go get -u github.com/vannleonheart/flip-api-go
```
#### Config
```go
flipConfig := flip.Config{
    BaseV2Url: "{flip_api_v2_url}",
    BaseV3Url: "{flip_api_v3_url}",
    SecretKey: "{your_secret_key}",
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
#### Get Bank Info
```go
result, err := flipClient.GetBankInfo()

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

fmt.Println(result.Balance)
```
#### Bank Account Inquiry
```go
result, err := flipClient.BankAccountInquiry()

if err != nil {
    // handle error
}

fmt.Println(result.Status)
```
#### Create Disbursement
```go
idempotencyKey := "{your_idempotency_key}"

requestData := DisbursementRequest {
    AccountNumber: "",
    BankCode: "",
    Amount: 100000,
    Remark: "",
    RecipientCity: "",
    BeneficiaryEmail: "",
}

result, err := flipClient.CreateDisbursement(requestData, idempotencyKey)

if err != nil {
    // handle error
}

fmt.Println(result.Status)
```
#### Get Disbursement By Id
```go
transactionId := "{transaction_id}"

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