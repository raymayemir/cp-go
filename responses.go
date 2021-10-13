package cp_go

import "time"

type BaseResponse struct {
	Success bool
	Message string
}
type PaymentModel struct {
	ID                int
	Amount            float64
	Currency          string
	CurrencyCode      int
	InvoiceID         string
	AccountID         string
	Email             string
	Description       string
	Data              []byte
	CreatedAt         time.Time
	AuthorizedAt      time.Time
	ConfirmedAt       time.Time
	AuthCode          string
	TestMode          bool
	IpAddress         string
	IpCountry         string
	IpCity            string
	IpRegion          string
	IpDistrict        string
	IpLatitude        float64
	IpLongitude       float64
	CardFirstSix      string
	CardLastFour      string
	CardExpiredMonth  int
	CardExpiredYear   int
	CardType          string
	CardTypeCode      int
	Issuer            string
	IssuerBankCountry string
	Status            string
	StatusCode        int
	Reason            string
	ReasonCode        int
	CardHolderMessage string
	CardHolderName    string
	Token             string
	JsonData          string
	Name              string
}

type SubscriptionModel struct {
	Id                           string
	CurrencyCode                 int
	StartDateIso                 string
	IntervalCode                 int
	StatusCode                   int
	Status                       string
	SuccessfulTransactionsNumber int
	FailedTransactionsNumber     int
	LastTransactionDate          string
	LastTransactionDateIso       string
	NextTransactionDate          string
	NextTransactionDateIso       string
}

type Payment3DSModel struct {
	TransactionId int
	PaReq         string
	AcsUrl        string
}

type PaymentFailedResponse struct {
	BaseResponse
	Model PaymentModel
}

type PaymentSuccessResponse struct {
	BaseResponse
	Model struct {
		PaymentModel
		AuthDate       string
		AuthDateIso    string
		AuthCode       string
		ConfirmDate    string
		ConfirmDateIso string
		Token          string
	}
}

type Payment3DSResponse struct {
	BaseResponse
	Model Payment3DSModel
}

type SubscriptionResponse struct {
	BaseResponse
	Model SubscriptionModel
}

type SubscriptionsListResponse struct {
	BaseResponse
	Model []SubscriptionModel
}
