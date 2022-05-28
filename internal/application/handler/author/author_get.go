package author

import (
	"bookstoreapi/internal/application/writer"
	"bookstoreapi/internal/domain/author"
	"bookstoreapi/internal/infra/repo"
	"context"
	"net/http"
	"strconv"
	"strings"
)

func GetAuthor(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	id, err := getAuthorId(w, req.URL.Path)

	if err != nil {
		return
	}

	repo := repo.NewAuthorRepo()
	authorEnt := author.NewAuthorEntity(repo)

	author := authorEnt.GetAuthorById(ctx, id)

	writer.SendResponse(w, 200, writer.JSONResponse{
		Message: "ok",
		Data:    author,
	})
}

func getAuthorId(w http.ResponseWriter, urlPath string) (int, error) {
	strId := strings.TrimPrefix(urlPath, "/authors/get/")

	id, err := strconv.Atoi(strId)

	if err != nil {
		writer.SendResponse(w, 400, writer.JSONResponse{
			Message: "error",
			Data:    "invalid author ID",
		})

		return 0, err
	}

	return id, nil
}
