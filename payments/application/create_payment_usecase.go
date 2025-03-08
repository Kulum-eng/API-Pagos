package aplication

import (
	"ModaVane/payments/domain"
	"ModaVane/payments/domain/ports"
)

type CreatePaymentUseCase struct {
	repo               ports.PaymentRepository
	senderNotification ports.SenderNotification
}

func NewCreatePaymentUseCase(repo ports.PaymentRepository, senderNotification ports.SenderNotification) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{repo: repo, senderNotification: senderNotification}
}

func (uc *CreatePaymentUseCase) Execute(payment domain.Payment) (int, error) {
	idPago, err := uc.repo.CreatePayment(payment)
	if err != nil {
		return 0, err
	}

	err = uc.senderNotification.SendNotification(map[string]interface{}{
		"event": "new-order",
		"data":  idPago,
	})
	if err != nil {
		return idPago, err
	}

	return idPago, nil
}
