package payment

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentMethod string

const (
	BankBNI     PaymentMethod = "bni"
	BankMandiri PaymentMethod = "mandiri"
	BankCIMB    PaymentMethod = "cimb"
	BankBCA     PaymentMethod = "bca"
	BankBRI     PaymentMethod = "bri"
	BankMaybank PaymentMethod = "maybank"
	BankPermata PaymentMethod = "permata"
	BankMega    PaymentMethod = "mega"
)

type ChargeRequest struct {
	PaymentType     string
	Invoice         string
	Total           int
	ItemsDetails    *[]midtrans.ItemDetails
	CustomerDetails *midtrans.CustomerDetails
}

type ChargeResponse struct {
	TransactionID          string             `json:"transaction_id"`
	OrderID                string             `json:"order_id"`
	GrossAmount            string             `json:"gross_amount"`
	PaymentType            string             `json:"payment_type"`
	TransactionTime        string             `json:"transaction_time"`
	TransactionStatus      string             `json:"transaction_status"`
	FraudStatus            string             `json:"fraud_status"`
	StatusCode             string             `json:"status_code"`
	Bank                   string             `json:"bank"`
	StatusMessage          string             `json:"status_message"`
	ChannelResponseCode    string             `json:"channel_response_code"`
	ChannelResponseMessage string             `json:"channel_response_message"`
	Currency               string             `json:"currency"`
	ValidationMessages     []string           `json:"validation_messages"`
	PermataVaNumber        string             `json:"permata_va_number"`
	VaNumbers              []coreapi.VANumber `json:"va_numbers"`
	BillKey                string             `json:"bill_key"`
	BillerCode             string             `json:"biller_code"`
	Actions                []coreapi.Action   `json:"actions"`
	PaymentCode            string             `json:"payment_code"`
	QRString               string             `json:"qr_string"`
	Expire                 string             `json:"expiry_time"`
}

type Midtrans struct {
	Client         coreapi.Client
	Request        *coreapi.ChargeReq
	OrderTime      string
	ExpiryDuration int
	ExpiryUnit     string
}

// func NewMidtrans(cfg Config) *Midtrans {
// 	return &Midtrans{
// 		Midtrans: coreapi.Client{
// 			ServerKey:  cfg.Midtrans.ServerKey,
// 			ClientKey:  cfg.Midtrans.ClientKey,
// 			Env:        midtrans.EnvironmentType(cfg.Midtrans.Env),
// 			HttpClient: midtrans.GetHttpClient(midtrans.EnvironmentType(cfg.Midtrans.Env)),
// 			Options: &midtrans.ConfigOptions{
// 				PaymentOverrideNotification: &cfg.Midtrans.URLHandler,
// 				PaymentAppendNotification:   &cfg.Midtrans.URLHandler,
// 			},
// 		},
// 		ExpDuration: cfg.Midtrans.ExpiryDuration,
// 		ExpUnit:     cfg.Midtrans.Unit,
// 	}
// }

func (m *Midtrans) CreateCharge(r ChargeRequest) (*ChargeResponse, error) {
	request := &coreapi.ChargeReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  r.Invoice,
			GrossAmt: int64(r.Total),
		},
		Items:           r.ItemsDetails,
		CustomerDetails: r.CustomerDetails,
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: m.ExpiryDuration,
			Unit:           m.ExpiryUnit,
		},
	}
	m.Request = request
	switch r.PaymentType {
	case "bca":
		return m.WithBank(BankBCA)
	case "mandiri", "cimb":
		return m.WithBank(BankMandiri)
	case "bni":
		return m.WithBank(BankBNI)
	}
	return nil, errors.New("payment type not available")
}

func (m *Midtrans) WithBank(pm PaymentMethod) (*ChargeResponse, error) {
	if pm != BankMandiri {
		m.Request.PaymentType = "bank_transfer"
		m.Request.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(pm),
		}
	} else {
		m.Request.PaymentType = coreapi.PaymentTypeEChannel
		m.Request.EChannel = &coreapi.EChannelDetail{
			BillInfo1: "pembayaran",
			BillInfo2: "pembayaran",
		}
	}
	res, err := m.ChargeCustom(m.Request)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res, nil
}

func (m *Midtrans) ChargeCustom(r *coreapi.ChargeReq) (*ChargeResponse, error) {
	response := ChargeResponse{}
	jsonRequest, _ := json.Marshal(r)
	err := m.Client.HttpClient.Call(http.MethodPost,
		fmt.Sprintf("%s/v2/charge", m.Client.Env.BaseUrl()),
		&m.Client.ServerKey,
		m.Client.Options,
		bytes.NewBuffer(jsonRequest),
		&response,
	)
	switch response.PaymentType {
	case "bank_transfer", "echannel":
		if response.PermataVaNumber != "" {
			response.PaymentCode = response.PermataVaNumber
		} else if response.BillerCode != "" || response.BillKey != "" {
			response.PaymentCode = fmt.Sprintf("BillCode:%s-BillKey:%s", response.BillerCode, response.BillKey)
		} else {
			response.PaymentCode = response.VaNumbers[0].VANumber
		}
	case "gopay", "qris":
		response.PaymentCode = response.Actions[0].URL
	}
	if err != nil {
		return nil, err
	}
	return &response, nil
}
