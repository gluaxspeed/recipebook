package api

import (
	"github.com/graphql-go/handler"
	"net/http"
	"recipebook/api/graphql"
)

func RecipeGraphqlAPI() http.Handler {
	h := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: false,
	})

	return h
}
