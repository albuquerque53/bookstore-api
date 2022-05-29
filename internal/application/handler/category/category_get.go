package category

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/category"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
)

func GetCategory(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	id, err := helper.GetPathId(w, req.URL.Path)

	if err != nil {
		return
	}

	repo := repo.NewCategoryRepo()
	categoryEnt := category.NewCategoryEntity(repo)

	category, err := categoryEnt.GetCategoryById(ctx, id)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on category search", err)
		return
	}

	writer.SendResponse(w, 200, writer.JSONResponse{
		Message: "ok",
		Data:    category,
	})
}
