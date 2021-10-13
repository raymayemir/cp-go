package cp_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://api.cloudpayments.ru/"

func (c *Client) sendRequest(endpoint string, params []byte, requestId *int) ([]byte, error) {
	req, err := http.NewRequest("POST", url+endpoint, bytes.NewBuffer(params))

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.config.PublicId, c.config.ApiSecret)
	req.Header.Set("Content-Type", "application/json")

	if requestId != nil {
		req.Header.Set("X-Request-ID", fmt.Sprint(requestId))
	}

	client := &http.Client{Timeout: c.config.Timeout}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (c *Client) Ping() (map[string]interface{}, error) {
	var data map[string]interface{}
	response, err := c.sendRequest("test", nil, nil)
	err = json.Unmarshal(response, &data)
	return data, err
}

// ChargeCryptogramPayment Payment by cryptogram
func (c *Client) ChargeCryptogramPayment(cpr *CryptogramPaymentRequest) (*PaymentSuccessResponse, *PaymentFailedResponse, error) {
	success := &PaymentSuccessResponse{}
	failed := &PaymentFailedResponse{}

	params, _ := json.Marshal(cpr)

	response, err := c.sendRequest("payments/cards/charge", params, nil)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(response, &success)

	if err != nil {
		err = json.Unmarshal(response, &failed)
		if err != nil {
			return nil, nil, err
		}
		return nil, failed, nil
	}

	return success, nil, nil
}

// AuthorizeCryptogramPayment two-stage payment
func (c *Client) AuthorizeCryptogramPayment(cpr CryptogramPaymentRequest) (*Payment3DSResponse, *PaymentFailedResponse, error) {
	success := &Payment3DSResponse{}
	failed := &PaymentFailedResponse{}

	params, _ := json.Marshal(cpr)

	response, err := c.sendRequest("payments/cards/auth", params, nil)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(response, &success)

	if err != nil {
		err = json.Unmarshal(response, &failed)
		if err != nil {
			return nil, nil, err
		}
		return nil, failed, nil
	}

	return success, nil, nil
}

// ChargeTokenPayment payment by token
func (c *Client) ChargeTokenPayment(tpr TokenPaymentRequest) (*PaymentSuccessResponse, *PaymentFailedResponse, error) {
	success := &PaymentSuccessResponse{}
	failed := &PaymentFailedResponse{}

	params, _ := json.Marshal(tpr)

	response, err := c.sendRequest("payments/tokens/charge", params, nil)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(response, &success)

	if err != nil {
		err = json.Unmarshal(response, &failed)
		if err != nil {
			return nil, nil, err
		}
		return nil, failed, nil
	}

	return success, nil, nil
}

// AuthorizeTokenPayment authorize token payment
func (c *Client) AuthorizeTokenPayment(tpr TokenPaymentRequest) (*PaymentSuccessResponse, *PaymentFailedResponse, error) {
	success := &PaymentSuccessResponse{}
	failed := &PaymentFailedResponse{}

	params, _ := json.Marshal(tpr)

	response, err := c.sendRequest("payments/tokens/auth", params, nil)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(response, &success)

	if err != nil {
		err = json.Unmarshal(response, &failed)
		if err != nil {
			return nil, nil, err
		}
		return nil, failed, nil
	}

	return success, nil, nil
}

// Confirm3DSPayment Confirm a 3DS payment
func (c *Client) Confirm3DSPayment(confirm3DS Confirm3DSRequest) (*PaymentSuccessResponse, *PaymentFailedResponse, error) {
	success := &PaymentSuccessResponse{}
	failed := &PaymentFailedResponse{}

	params, _ := json.Marshal(confirm3DS)

	response, err := c.sendRequest("payments/cards/post3ds", params, nil)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(response, &success)

	if err != nil {
		err = json.Unmarshal(response, &failed)
		if err != nil {
			return nil, nil, err
		}
		return nil, failed, nil
	}

	return success, nil, nil
}

// ConfirmPayment confirm an authorized payment
func (c *Client) ConfirmPayment(confirm ConfirmPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, _ := json.Marshal(confirm)
	log.Println(string(params))
	response, err := c.sendRequest("payments/confirm", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)

	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// RefundPayment refund
func (c *Client) RefundPayment(rpp RefundPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, _ := json.Marshal(rpp)
	response, err := c.sendRequest("payments/refund", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)

	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// VoidPayment cancellation of payment
func (c *Client) VoidPayment(vpr VoidPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, _ := json.Marshal(vpr)
	response, err := c.sendRequest("payments/void", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)

	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// GetPayment
func (c *Client) GetPayment(gpr GetPaymentRequest) {
	// TODO implement
}

// FindPaymentByInvoiceId
func (c *Client) FindPaymentByInvoiceId() {
	// TODO implement
}

// GetPaymentsList
func (c *Client) GetPaymentsList() {
	// TODO implement
}

// CreateOrder to send by email
func (c *Client) CreateOrder(lpr LinkPaymentRequest) (*BaseResponse, error) {

	baseResponse := &BaseResponse{}

	params, _ := json.Marshal(lpr)
	response, err := c.sendRequest("orders/create", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)

	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// CreateSubscription create subscription
func (c *Client) CreateSubscription(scr SubscriptionCreateRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, _ := json.Marshal(scr)
	response, err := c.sendRequest("subscriptions/create", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &subscriptionResponse)

	if err != nil {
		return nil, err
	}

	return subscriptionResponse, nil
}

// UpdateSubscription update subscription
func (c *Client) UpdateSubscription(sur SubscriptionUpdateRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, _ := json.Marshal(sur)
	response, err := c.sendRequest("subscriptions/update", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &subscriptionResponse)

	if err != nil {
		return nil, err
	}

	return subscriptionResponse, nil
}

// CancelSubscription cancel subscription
func (c *Client) CancelSubscription(sur SubscriptionUpdateRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, _ := json.Marshal(sur)
	response, err := c.sendRequest("subscriptions/cancel", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &baseResponse)

	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

// GetSubscription get subscription by ID of transaction
func (c *Client) GetSubscription(sgr SubscriptionGetRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, _ := json.Marshal(sgr)
	response, err := c.sendRequest("subscriptions/get", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &subscriptionResponse)

	if err != nil {
		return nil, err
	}

	return subscriptionResponse, nil
}

// GetSubscriptionsList get subscription list by AccountId
func (c *Client) GetSubscriptionsList(slr SubscriptionListRequest) (*SubscriptionsListResponse, error) {
	list := &SubscriptionsListResponse{}

	params, _ := json.Marshal(slr)
	response, err := c.sendRequest("subscriptions/find", params, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c *Client) ChargeCryptogramPayout() {
	// TODO implement
}

func (c *Client) ChargeTokenPayout() {
	// TODO implement
}
