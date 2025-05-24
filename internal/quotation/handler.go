package quotation

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"rest_go_quotation_book/pkg/res"
	"rest_go_quotation_book/pkg/utils"
	"sync"
)

type QuotationHandler struct {
	mu         sync.Mutex
	quotations []Quotation
	nextID     int
}

func NewQuotationHandler(router *http.ServeMux) {
	handler := &QuotationHandler{
		quotations: make([]Quotation, 0),
		nextID:     1,
	}

	router.HandleFunc("POST /quotes", handler.Create())
	router.HandleFunc("GET /quotes", handler.GetAll())
	router.HandleFunc("GET /quotes/random", handler.GetRandom())
	router.HandleFunc("GET /quotes/author/{author}", handler.GetAuthor())
	router.HandleFunc("DELETE /quotes/{id}", handler.Delete())
}

// Create godoc
// @Summary Создать цитату
// @Description Добавляет новую цитату в список
// @Tags quotes
// @Accept json
// @Produce json
// @Param quote body CreateQuotationRequest true "Цитата для добавления"
// @Success 201 {object} CreateQuotationResponse
// @Failure 400 {object} map[string]string
// @Router /quotes [post]
func (handler *QuotationHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Чтение body
		var payload CreateQuotationRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			res.Json(w, map[string]string{"error": "author and quote are required"}, http.StatusBadRequest)
			return
		}

		// Валидация
		if payload.Author == "" || payload.Quote == "" {
			res.Json(w, map[string]string{"error": "author and quote are required"}, http.StatusBadRequest)
			return
		}

		// Создание и сохранение цитаты
		handler.mu.Lock()
		quote := Quotation{
			ID:     handler.nextID,
			Author: payload.Author,
			Quote:  payload.Quote,
		}
		handler.quotations = append(handler.quotations, quote)
		handler.nextID++
		handler.mu.Unlock()

		// Формируем ответ
		response := CreateQuotationResponse{
			ID:     quote.ID,
			Author: quote.Author,
			Quote:  quote.Quote,
		}

		res.Json(w, response, 201)
	}
}

// GetAll godoc
// @Summary Получить все цитаты
// @Description Возвращает список всех цитат
// @Tags quotes
// @Produce json
// @Success 200 {array} Quotation
// @Router /quotes [get]
func (handler *QuotationHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.mu.Lock()
		defer handler.mu.Unlock()

		res.Json(w, handler.quotations, http.StatusOK)
	}
}

// GetRandom godoc
// @Summary Получить случайную цитату
// @Description Возвращает случайную цитату из списка
// @Tags quotes
// @Produce json
// @Success 200 {object} Quotation
// @Failure 404 {object} map[string]string
// @Router /quotes/random [get]
func (handler *QuotationHandler) GetRandom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.mu.Lock()
		defer handler.mu.Unlock()

		if len(handler.quotations) == 0 {
			res.Json(w, handler.quotations, http.StatusNotFound)
			return
		}

		randomIndex := rand.Intn(len(handler.quotations))
		randomQuotation := handler.quotations[randomIndex]

		res.Json(w, randomQuotation, http.StatusOK)
	}
}

// GetAuthor godoc
// @Summary Получить цитаты по автору
// @Description Возвращает все цитаты указанного автора
// @Tags quotes
// @Produce json
// @Param author path string true "Автор цитаты"
// @Success 200 {array} Quotation
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /quotes/author/{author} [get]
func (handler *QuotationHandler) GetAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author, err := utils.ParseAuthor(r)
		if err != nil {
			http.Error(w, "author not specified", http.StatusBadRequest)
			return
		}

		handler.mu.Lock()
		defer handler.mu.Unlock()

		var result []Quotation
		for _, q := range handler.quotations {
			if q.Author == author {
				result = append(result, q)
			}
		}

		if len(result) == 0 {
			http.Error(w, "no quotes found for this author", http.StatusNotFound)
			return
		}

		res.Json(w, result, http.StatusOK)
	}
}

// Delete godoc
// @Summary Удалить цитату
// @Description Удаляет цитату по ID
// @Tags quotes
// @Produce json
// @Param id path int true "ID цитаты"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /quotes/{id} [delete]
func (handler *QuotationHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseID(r)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		handler.mu.Lock()
		defer handler.mu.Unlock()

		index := -1
		for i, q := range handler.quotations {
			if q.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			http.Error(w, "quote not found", http.StatusNotFound)
			return
		}

		// Удаляем элемент с индексом index
		handler.quotations = append(handler.quotations[:index], handler.quotations[index+1:]...)

		res.Json(w, map[string]string{"message": "quote deleted"}, http.StatusOK)

	}
}
