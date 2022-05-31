package book

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/book"
	"bookstoreapi/internal/infra/repo"
	"context"
	"encoding/json"
	"net/http"
)

func NewBook(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	dto, err := getRequestBody(w, req)

	if err != nil {
		return
	}

	repo := repo.NewBookRepo()
	bookEnt := book.NewBookEntity(repo)

	err = bookEnt.CreateNewBook(ctx, dto)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on book create", err)
	}

	writer.SendResponse(w, 204, writer.JSONResponse{
		Message: "ok",
		Data:    nil,
	})
}

func getRequestBody(w http.ResponseWriter, req *http.Request) (*book.BookDto, error) {
	var dto *book.BookDto

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dto)

	if err != nil {
		helper.HandleHttpError(w, 400, "invalid fields", nil)

		return nil, err
	}

	return dto, nil
}
