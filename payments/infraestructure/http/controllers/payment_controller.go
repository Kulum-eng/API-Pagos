package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ModaVane/payments/application"
	"ModaVane/payments/domain"
	"ModaVane/payments/infraestructure/http/responses"

)

type PaymentController struct {
	createPaymentUseCase *aplication.CreatePaymentUseCase
	getPaymentUseCase    *aplication.GetPaymentUseCase
	updatePaymentUseCase *aplication.UpdatePaymentUseCase
	deletePaymentUseCase *aplication.DeletePaymentUseCase
}

func NewPaymentController(createUC *aplication.CreatePaymentUseCase, getUC *aplication.GetPaymentUseCase, updateUC *aplication.UpdatePaymentUseCase, deleteUC *aplication.DeletePaymentUseCase) *PaymentController {
	return &PaymentController{
		createPaymentUseCase: createUC,
		getPaymentUseCase:    getUC,
		updatePaymentUseCase: updateUC,
		deletePaymentUseCase: deleteUC,
	}
}

func (ctrl *PaymentController) Create(ctx *gin.Context) {
	var payment domain.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("los datos son inválidos", err.Error()))
		return
	}

	idPayment, err := ctrl.createPaymentUseCase.Execute(payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear pago", err.Error()))
		return
	}

	payment.ID = idPayment
	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Pago creado exitosamente", payment))
}

func (ctrl *PaymentController) GetAll(ctx *gin.Context) {
	payments, err := ctrl.getPaymentUseCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener pagos", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Pagos obtenidos exitosamente", payments))
}

func (ctrl *PaymentController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	payment, err := ctrl.getPaymentUseCase.ExecuteByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener pago", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Pago obtenido exitosamente", payment))
}

func (ctrl *PaymentController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	var payment domain.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	payment.ID = id
	if err := ctrl.updatePaymentUseCase.Execute(payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al actualizar pago", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Pago actualizado exitosamente", payment))
}

func (ctrl *PaymentController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	if err := ctrl.deletePaymentUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al eliminar pago", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Pago eliminado exitosamente", nil))
}
