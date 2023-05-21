package webhook

import (
	"errors"
	"os"
	"strings"

	classRoomDto "edtech.id/internal/class_room/dto"
	classRoomUsecase "edtech.id/internal/class_room/usecase"
	orderDto "edtech.id/internal/order/dto"
	orderUsecase "edtech.id/internal/order/usecase"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type WebhookUsecase interface {
	UpdatePayment(id string) error
}

type WebhookUsecaseImpl struct {
	orderUsecase     orderUsecase.OrderUsecase
	classRoomUsecase classRoomUsecase.ClassRoomUsecase
}

// UpdatePayment implements WebhookUsecase
func (wu *WebhookUsecaseImpl) UpdatePayment(id string) error {
	invoiceParams := invoice.GetParams{
		ID: id,
	}

	dataXendit, err := invoice.Get(&invoiceParams)

	if err != nil {
		return err
	}

	dataOrder, orderError := wu.orderUsecase.FindByExternalId(dataXendit.ExternalID)

	if orderError != nil {
		return orderError
	}

	if dataOrder == nil {
		return errors.New("order not found")
	}

	if dataXendit.Status == "settled" {
		return errors.New("order already paid")
	}

	if dataXendit.Status != "paid" {
		if dataXendit.Status == "PAID" || dataXendit.Status == "SETTLED" {
			// add classrooms
			for _, orderDetail := range dataOrder.OrderDetails {
				dataClassRoom := classRoomDto.ClassRoomRequestBody{
					UserID:    dataOrder.UserID,
					ProductID: orderDetail.ProductID,
				}

				_, err := wu.classRoomUsecase.Create(dataClassRoom)

				if err != nil {
					return err
				}
			}

			// TODO: send notif
		}
	}

	// Update data order
	orderDto := orderDto.OrderRequestBody{
		Status: strings.ToLower(dataXendit.Status),
	}

	_, updateOrderErr := wu.orderUsecase.Update(int(dataOrder.ID), orderDto)

	if updateOrderErr != nil {
		return updateOrderErr
	}

	return nil

}

func NewWebhookUsecase(
	orderUsecase orderUsecase.OrderUsecase,
	classRoomUsecase classRoomUsecase.ClassRoomUsecase,
) WebhookUsecase {
	xendit.Opt.SecretKey = os.Getenv("XENDIT_API_KEY")
	return &WebhookUsecaseImpl{orderUsecase, classRoomUsecase}
}
