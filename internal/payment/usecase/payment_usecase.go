package payment

import (
	"os"

	paymentDto "edtech.id/internal/payment/dto"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type PaymentUsecase interface {
	Create(paymentDto paymentDto.PaymentRequestBody) (*xendit.Invoice, error)
}

type PaymentUsecaseImpl struct {
}

// Create implements PaymentUsecase
func (pu *PaymentUsecaseImpl) Create(paymentDto paymentDto.PaymentRequestBody) (*xendit.Invoice, error) {
	data := invoice.CreateParams{
		ExternalID:  paymentDto.ExternalID,
		Description: paymentDto.Description,
		Amount:      float64(paymentDto.Amount),
		Customer: xendit.InvoiceCustomer{
			Email: paymentDto.PayerEmail,
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"email"},
			InvoicePaid:     []string{"email"},
			InvoiceReminder: []string{"email"},
			InvoiceExpired:  []string{"email"},
		},
		InvoiceDuration:    86400,
		SuccessRedirectURL: os.Getenv("XENDIT_SUCCESS_URL"),
	}

	response, err := invoice.Create(&data)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewPaymentUsecase() PaymentUsecase {

	xendit.Opt.SecretKey = os.Getenv("XENDIT_API_KEY")

	return &PaymentUsecaseImpl{}
}
