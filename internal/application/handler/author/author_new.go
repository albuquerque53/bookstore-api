package author

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"encoding/json"
	"net/http"
)

func NewAuthor(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	dto, err := getRequestBody(w, req)

	if err != nil {
		return
	}

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	authorEnt.CreateNewAuthor(ctx, dto)

	writer.SendResponse(w, 204, writer.JSONResponse{
		Message: "ok",
		Data:    nil,
	})
}

func getRequestBody(w http.ResponseWriter, req *http.Request) (*author.AuthorDto, error) {
	var dto *author.AuthorDto

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dto)
	if err != nil {
		helper.HandleHttpError(w, 400, "invalid fields", nil)

		return nil, err
	}

	return dto, nil
}
