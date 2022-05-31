package book

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/book"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
)

func ListBooks(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	repo := repo.NewBookRepo()
	bookEnt := book.NewBookEntity(repo)

	authors, err := bookEnt.GetAllBooks(ctx)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on books list", err)
		return
	}

	writer.SendResponse(w, 200, writer.JSONResponse{
		Message: "ok",
		Data:    authors,
	})
}
