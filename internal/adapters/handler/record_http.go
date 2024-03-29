package handler

import (
	"net/http"

	"example.com/internal/core/domain"
	"example.com/internal/core/services"
	"example.com/utils"
	"github.com/IBM/fp-go/function"
	E "github.com/IBM/fp-go/ioeither"
	"github.com/IBM/fp-go/option"
)

type HTTPRecordHandler struct {
	svc services.RecordService
}

func NewHTTPRecordHandler(message services.RecordService) *HTTPRecordHandler {
	return &HTTPRecordHandler{
		svc: message,
	}
}

func (h *HTTPRecordHandler) AddRecord(w http.ResponseWriter, r *http.Request) {
	function.Pipe3(r.Body,
		utils.ParseJSON[domain.Record],
		E.Chain(h.svc.AddRecord),
		E.Fold[error, domain.Record, int](WriteLeft(w), WriteRight[domain.Record](http.StatusCreated, w, option.None[string]())))()
}
