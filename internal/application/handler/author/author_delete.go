package author

import (
	"bookstoreapi/internal/application/helper"
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
)

func DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	id, err := helper.GetPathId(w, req.URL.Path)

	if err != nil {
		return
	}

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	err = authorEnt.DeleteAuthorById(ctx, id)

	if err != nil {
		helper.HandleHttpError(w, 400, "error on author search", err)
		return
	}

	writer.SendResponse(w, 204, writer.JSONResponse{
		Message: "ok",
		Data:    nil,
	})
}
