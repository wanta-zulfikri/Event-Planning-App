package payment

import (
	"context"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/wanta-zulfikri/Event-Planning-App/config/common"
)

var snapClient snap.Client

func (r PaymentGateway) InitializeClientMidtrans() {
	snapClient.New(common.MIDTRANS_SERVER_KEY, midtrans.Sandbox)
}

func (r PaymentGateway) CreateTransaction(req PaymentGateway) string {
	snapUrl, err := snapClient.CreateTransactionToken(generateSnapReq(req))

	if err != nil {
		fmt.Printf("Midtrans error : %v", err.GetMessage())
	}
	return snapUrl
}

func (r PaymentGateway) CreateUrlTransactionWithGateway(req PaymentGateway) string {
	snapClient.Options.SetContext(context.Background())

	snapUrl, err := snapClient.CreateTransactionUrl(generateSnapReq(req))

	if err != nil {
		fmt.Printf("Midtrans error : %v", err.GetMessage())
	}
	return snapUrl
}

func generateSnapReq(req PaymentGateway) *snap.Request {
	reqSnap := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderId,
			GrossAmt: req.GrossAmt,
		},
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeBNIVA,
			snap.PaymentTypePermataVA,
			snap.PaymentTypeBCAVA,
			snap.PaymentTypeBRIVA,
			snap.PaymentTypeBankTransfer,
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: req.Email,
			Phone: req.Phone,
		},
		Items: &req.Items,
	}
	return reqSnap
}
