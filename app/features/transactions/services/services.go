package services

// import (
// 	"errors"
// 	"time"

// 	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
// 	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
// 	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
// )

// type TransactionService struct {
// 	m transactions.Repository
// 	u users.Repository
// 	e events.Repository
// }

// func New(r transactions.Repository) transactions.Service {
// 	return &TransactionService{m: r}
// }

// // cek event yang user pilih, jika eventid tersedia maka appen to slice event
// // calculate total transaction
// // initiate the transaction, then create transaction
// // initiate the transaction detail with loop over the events above and append to slice eventAttends
// // update availability of each event to "0", which is unavailable
// // then create the transaction detail with status
// // setup response from my app

// func (ts TransactionService) CreateTransaction(newTransaction transactions.Core, userId int) (map[string]interface{}, error) {
// 	// Get the events for the transaction
// 	var events []events.Core
// 	for _, ticket := range newTransaction.Tickets {
// 		event, err := ts.e.GetEventByEventID(int(newTransaction.Event.ID))
// 		if err != nil {
// 			return nil, err
// 		}
// 		if Ticket.Availability < ticket.TicketQuantity {
// 			return nil, errors.New("Kuota event sudah habis")
// 		}
// 		event.Availability -= ticket.TicketQuantity
// 		events = append(events, event)
// 	}

// 	// Calculate total payments
// 	var subtotal uint
// 	for _, ticket := range newTransaction.Tickets {
// 		subtotal += ticket.TicketPrice * ticket.TicketQuantity
// 	}
// 	grandTotal := subtotal // no additional cost for now

// 	// Initiate the payment
// 	payment, err := ts.m.InitiatePayment(newTransaction.PaymentType, grandTotal)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create the transaction
// 	newTransaction.Status = "pending"
// 	newTransaction.PurchaseStartDate = time.Now()
// 	newTransaction.PurchaseDueDate = time.Now().Add(time.Hour * 24 * 1) // 1 day
// 	newTransaction.StatusDate = time.Now()
// 	newTransaction.GrandTotal = grandTotal
// 	newTransaction.PaymentStatus = "pending"
// 	newTransaction.PaymentLink = payment.RedirectUrl // use payment redirect url
// 	if err := ts.m.CreateTransaction(&newTransaction); err != nil {
// 		return nil, err
// 	}

// 	// Create the transaction details
// 	var transactionDetails []transactions.TransactionDetail
// 	for i, ticket := range newTransaction.Tickets {
// 		transactionDetail := transactions.TransactionDetail{
// 			TicketCategory:  ticket.TicketCategory,
// 			TicketQuantity:  ticket.TicketQuantity,
// 			TicketPrice:     ticket.TicketPrice,
// 			Event:           events[i],
// 			Transaction:     &newTransaction,
// 			TransactionDate: time.Now(),
// 			Status:          "pending",
// 		}
// 		transactionDetails = append(transactionDetails, transactionDetail)

// 		if err := ts.m.UpdateEventAvailability(&events[i]); err != nil {
// 			return nil, err
// 		}
// 	}
// 	if err := ts.m.CreateTransactionDetails(transactionDetails); err != nil {
// 		return nil, err
// 	}

// 	// Create user history
// 	history := users.History{
// 		UserId: userId,
// 		Items: []users.HistoryItem{{
// 			ID:          newTransaction.Invoice,
// 			Name:        "Event Transaction",
// 			Description: "Purchase of tickets for events",
// 			Amount:      int64(grandTotal),
// 			Quantity:    1,
// 			Date:        time.Now(),
// 		}},
// 	}
// 	if err := ts.m.CreateUserHistory(&history); err != nil {
// 		return nil, err
// 	}

// 	data := map[string]interface{}{
// 		"order_id":       newTransaction.Invoice,
// 		"total_payments": grandTotal,
// 		"payments": map[string]interface{}{
// 			"id":             payment.ID,
// 			"payment_status": payment.PaymentStatus,
// 			"payment_type":   payment.PaymentType,
// 			"created_at":     payment.CreatedAt,
// 			"updated_at":     payment.UpdatedAt,
// 		},
// 		"payment_link": payment.RedirectUrl,
// 	}

// 	return data, nil
// }
