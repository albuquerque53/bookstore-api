package category

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/category"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
)

func ListCategories(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	repo := repo.NewCategoryRepo()
	categoryEnt := category.NewCategoryEntity(repo)

	authors, err := categoryEnt.GetAllCategories(ctx)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on authors list", err)
		return
	}

	writer.SendResponse(w, 200, writer.JSONResponse{
		Message: "ok",
		Data:    authors,
	})
}
