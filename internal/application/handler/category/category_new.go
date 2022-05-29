package category

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/category"
	"bookstoreapi/internal/infra/repo"
	"context"
	"encoding/json"
	"net/http"
)

func NewCategory(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	dto, err := getRequestBody(w, req)

	if err != nil {
		return
	}

	repo := repo.NewCategoryRepo()
	categoryEnt := category.NewCategoryEntity(repo)

	err = categoryEnt.CreateNewCategory(ctx, dto)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on category create", err)
	}

	writer.SendResponse(w, 204, writer.JSONResponse{
		Message: "ok",
		Data:    nil,
	})
}

func getRequestBody(w http.ResponseWriter, req *http.Request) (*category.CategoryDto, error) {
	var dto *category.CategoryDto

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dto)

	if err != nil {
		helper.HandleHttpError(w, 400, "invalid fields", nil)

		return nil, err
	}

	return dto, nil
}
