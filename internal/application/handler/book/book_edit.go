package book

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/book"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
)

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	id, err := helper.GetPathId(w, req.URL.Path)

	if err != nil {
		return
	}

	dto, err := getRequestBody(w, req)

	if err != nil {
		return
	}

	repo := repo.NewBookRepo()
	bookEnt := book.NewBookEntity(repo)

	err = bookEnt.UpdateBookById(ctx, id, dto)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on book update", err)
		return
	}

	writer.SendResponse(w, 204, writer.JSONResponse{
		Message: "ok",
		Data:    nil,
	})
}
