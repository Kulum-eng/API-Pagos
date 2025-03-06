package aplication

import (
	"ModaVane/payments/domain/ports"

)

type DeletePaymentUseCase struct {
	repo ports.PaymentRepository
}

func NewDeletePaymentUseCase(repo ports.PaymentRepository) *DeletePaymentUseCase {
	return &DeletePaymentUseCase{repo: repo}
}

func (uc *DeletePaymentUseCase) Execute(id int) error {
	return uc.repo.DeletePayment(id)
}
