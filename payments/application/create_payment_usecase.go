package application

import (
	"encoding/json"
	"strconv"
	"time"

	"ModaVane/payments/domain"
	"ModaVane/payments/domain/ports"

)

type CreatePaymentUseCase struct {
	repo               ports.PaymentRepository
	broker             ports.Broker
	senderNotification ports.SenderNotification
}

func NewCreatePaymentUseCase(repo ports.PaymentRepository, broker ports.Broker, senderNotification ports.SenderNotification) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{
		repo:               repo,
		broker:             broker,
		senderNotification: senderNotification,
	}
}

func (uc *CreatePaymentUseCase) Execute(payment domain.Payment) (int, error) {
	idPago, err := uc.repo.CreatePayment(payment)
	if err != nil {
		return 0, err
	}
	idPagoStr := strconv.Itoa(idPago)

	messageJson := map[string]interface{}{
		"order_id": payment.OrderID,
	}

	messageJsonStr, err := json.Marshal(messageJson)
	if err != nil {
		return idPago, err
	}

	err = uc.broker.Publish(string(messageJsonStr))
	if err != nil {
		return idPago, err
	}

	time.Sleep(5 * time.Second)

	err = uc.senderNotification.SendNotification(map[string]interface{}{
		"event": "new-payment",
		"data":  idPagoStr,
	})

	if err != nil {
		return idPago, err
	}

	return idPago, nil
}
