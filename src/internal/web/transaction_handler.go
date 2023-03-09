package web

import (
	"encoding/json"
	"net/http"

	createtransaction "github.com/deividroger/ms-wallet/src/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	TransactionUseCase createtransaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(transactionUseCase createtransaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{
		TransactionUseCase: transactionUseCase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var dto createtransaction.CreateTransactionInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.TransactionUseCase.Execute(dto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
