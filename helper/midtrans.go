package helper

import (
	"log"

	"github.com/google/uuid"
	pg "github.com/pandudpn/go-payment-gateway"
	"github.com/pandudpn/go-payment-gateway/gateway/midtrans"
)

const SandBoxServerKey string = "SB-Mid-server-OvkShVQ_Lno2P2-jMFR9m1Hq"
const SandBoxClientKey string = "SB-Mid-client-dWlmfZziZUWwJy87"

func BankTransferCharge(opts *pg.Options) {
	id := uuid.NewString()
	bt := &midtrans.BankTransferCreateParams{
		PaymentType: midtrans.PaymentTypeMandiri,
		TransactionDetails: &midtrans.TransactionDetail{
			OrderID:     id,
			GrossAmount: 10000,
		},
		ItemDetails: []*midtrans.ItemDetail{
			{
				ID:       uuid.NewString(),
				Price:    10000,
				Name:     "abc",
				Quantity: 1,
			},
		},
		EChannel: &midtrans.EChannel{
			BillKey:   "123456789",
			BillInfo1: "Pembayaran:",
			BillInfo2: id,
		},
		BankTransfer: &midtrans.BankTransfer{
			Bank:     midtrans.BankPermata,
			VANumber: "1234567890",
		},
	}

	res, err := midtrans.CreateBankTransferCharge(bt, opts)
	if err != nil {
		log.Fatalln("failed to create bank_transfer charge with error:", err)
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
}
