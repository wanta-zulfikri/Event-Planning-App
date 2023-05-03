package helper

import (
	"errors"
	"log"

	pg "github.com/pandudpn/go-payment-gateway"
	"github.com/pandudpn/go-payment-gateway/gateway/midtrans"
)

const SandBoxServerKey string = "SB-Mid-server-OvkShVQ_Lno2P2-jMFR9m1Hq"
const SandBoxClientKey string = "SB-Mid-client-dWlmfZziZUWwJy87"

type TransactionDetail struct {
	OrderID     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
}

func SetTransactionDetails(bt *midtrans.BankTransferCreateParams, details []TransactionDetail) {
	bt.TransactionDetails = &midtrans.TransactionDetail{}
	for _, detail := range details {
		bt.TransactionDetails.GrossAmount += detail.GrossAmount
		bt.ItemDetails = append(bt.ItemDetails, &midtrans.ItemDetail{
			ID:    detail.OrderID,
			Price: detail.GrossAmount,
		})
	}
}

func CreateBankTransferCharge(bt *midtrans.BankTransferCreateParams, opts *pg.Options) (*midtrans.ChargeResponse, error) {
	SetTransactionDetails(bt, []TransactionDetail{})
	res, err := midtrans.CreateBankTransferCharge(bt, opts)
	if err != nil {
		log.Println("failed to create bank_transfer charge with error:", err)
		return nil, errors.New("failed to create bank_transfer charge")
	}

	log.Println("response bank_transfer charge", *res)
	if len(res.VANumbers) > 0 {
		log.Println("virtual account", bt.BankTransfer.Bank, res.VANumbers[0].VANumber)
	}

	if res.PermataVANumber != "" {
		log.Println("virtual account permata", res.PermataVANumber)
	}

	if res.BillKey != "" {
		log.Println("virtual account mandiri", res.BillerCode+"-"+res.BillKey)
	}

	return res, nil
}
