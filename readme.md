### CloudPayments Golang Client 

Client for the [Cloud Payments](https://cloudpayments.ru) payment service allows access the [Cloud Payments API](https://developers.cloudpayments.ru/#api) from Golang code.

### Install

`go get github.com/raymayemir/cp-go`

### Config
| Property | Type | Description | Default | Required
| :--- | :--- | :--- | :--- | :--- |
| ApiSecret | `string` | Secret key (you can get it in your personal account) | `empty string` | `true` |
| PublicId | `string` |  Public key (you can get it in your personal account) | `empty string` | `true` |
| Timeout | `time.Duration` | Request timeout | `30 * time.Second` |  `false` |

### Usage 
```go

import cp "github.com/raymayemir/cp-go"

client := cp.NewClient(cp.Config{
    ApiSecret: "api_secret",
    PublicId:  "public_id",
    Timeout:   30 * time.Second,
})

response, err := client.Ping()
...
```

### Test method [description](https://developers.cloudpayments.ru/#testovyy-metod)


```go
response, err := client.Ping()
```
If successful, it returns a `map[string]interface{}` else returns `error`.

### Client API
| Method | Description | Documentation Link |
| :--- | :--- | :--- | 
| ChargeCryptogramPayment | Payment by cryptogram | https://developers.cloudpayments.ru/#oplata-po-kriptogramme |
| AuthorizeCryptogramPayment | Payment by cryptogram (authorization) |  https://developers.cloudpayments.ru/#oplata-po-kriptogramme |
| ChargeTokenPayment | Payment by token | https://developers.cloudpayments.ru/#oplata-po-tokenu-rekarring |
| AuthorizeTokenPayment | Payment by token (authorization) | https://developers.cloudpayments.ru/#oplata-po-tokenu-rekarring |
| Confirm3DSPayment | 3-D Secure Processing | https://developers.cloudpayments.ru/#obrabotka-3-d-secure |
| ConfirmPayment | Payment confirmation | https://developers.cloudpayments.ru/#podtverzhdenie-oplaty |
| RefundPayment | Refund money | https://developers.cloudpayments.ru/#vozvrat-deneg |
| VoidPayment | Cancellation of payment | https://developers.cloudpayments.ru/#otmena-oplaty |
| GetPayment | Viewing operation information(todo) | https://developers.cloudpayments.ru/#prosmotr-tranzaktsii |
| FindPaymentByInvoiceId | Checking the payment status(todo) | https://developers.cloudpayments.ru/#proverka-statusa-platezha |
| GetPaymentsList | List of transactions(todo) | https://developers.cloudpayments.ru/#vygruzka-spiska-tranzaktsiy |
| CreateOrder | Creating an invoice to send by mail | https://developers.cloudpayments.ru/#sozdanie-scheta-dlya-otpravki-po-pochte |
| CreateSubscription | Creating a subscription  | https://developers.cloudpayments.ru/#sozdanie-podpiski-na-rekurrentnye-platezhi |
| UpdateSubscription | Changing a subscription | https://developers.cloudpayments.ru/#izmenenie-podpiski-na-rekurrentnye-platezhi |
| CancelSubscription | Cancellation of subscription | https://developers.cloudpayments.ru/#izmenenie-podpiski-na-rekurrentnye-platezhi |
| GetSubscription | Request for subscription information | https://developers.cloudpayments.ru/#zapros-informatsii-o-podpiske |
| GetSubscriptionsList | Search subscriptions | https://developers.cloudpayments.ru/#poisk-podpisok |
| ChargeCryptogramPayout | Payout by cryptogram(todo) | https://developers.cloudpayments.ru/#vyplata-po-kriptogramme |
| ChargeTokenPayout | Payout by token(todo) | https://developers.cloudpayments.ru/#vyplata-po-tokenu |

### TODO

* Receipt
* Notification handlers?
* Finish (todo) marked methods

### Discussion

* API methods multiple return values or just no need to `PaymentFailedResponse`


### Problems and improvements

* Create issue
* Create pull-request


### License

MIT

