package aplication

import (
	"ModaVane/payments/domain"
	"ModaVane/payments/domain/ports"

)

type UpdatePaymentUseCase struct {
	repo ports.PaymentRepository
}

func NewUpdatePaymentUseCase(repo ports.PaymentRepository) *UpdatePaymentUseCase {
	return &UpdatePaymentUseCase{repo: repo}
}

func (uc *UpdatePaymentUseCase) Execute(payment domain.Payment) error {
	return uc.repo.UpdatePayment(payment)
}
