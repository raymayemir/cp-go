package cp_go

type BaseRequest struct {
	CultureName string
}

type PaymentRequest struct {
	BaseRequest
	Amount      float64
	Currency    string
	IpAddress   string
	Name        string
	InvoiceId   string
	Description string
	Email       string
	JsonData    string
}

type CryptogramPaymentRequest struct {
	PaymentRequest
	IpAddress            string
	CardCryptogramPacket string
	AccountId            *string
}

type TokenPaymentRequest struct {
	PaymentRequest
	AccountId string
	Token     string
}

type Confirm3DSRequest struct {
	BaseRequest
	TransactionId string
	PaRes         string
}

type ConfirmPaymentRequest struct {
	BaseRequest
	TransactionId int
	Amount        float64
	JsonData      *string
}

type RefundPaymentRequest struct {
	BaseRequest
	TransactionId int
	Amount        float64
	JsonData      *string
}

type VoidPaymentRequest struct {
	BaseRequest
	TransactionId int
}

type GetPaymentRequest struct {
	BaseRequest
	TransactionId int
}

type LinkPaymentRequest struct {
	BaseRequest
	Amount      int
	Currency    string // ValidCurrency
	JsonData    *string
	Description string
	email       string
	phone       string
}

type AccountRequest struct {
	AccountId string
	Email     string
}

type SubscriptionBase struct {
	AccountRequest
	Description         string
	Amount              float64
	Currency            string
	RequireConfirmation bool
	StartDate           string
	Interval            string
	Period              int
	MaxPeriods          int
}

type SubscriptionCreateRequest struct {
	SubscriptionBase
	Token string
}

type SubscriptionUpdateRequest struct {
	SubscriptionBase
	Id string
}

type SubscriptionGetRequest struct {
	BaseRequest
	Id string
}

type SubscriptionListRequest struct {
	BaseRequest
	AccountId string
}
