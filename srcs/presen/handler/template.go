package handler

import (
	"github.com/shwatanap/go-backend-template/srcs/usecase"
)

type TemplateHandler interface{}

type templateHandler struct {
	templateUsecase usecase.TemplateUsecase
}

func NewTemplateHandler(ut usecase.TemplateUsecase) TemplateHandler {
	return &templateHandler{templateUsecase: ut}
}
