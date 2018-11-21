package graphql

import (
	"context"
	"github.com/graphql-go/graphql"
	"recipebook/models"
	"recipebook/mongo"
	"recipebook/utils"
)

var (
	recipeType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Recipe",
		Description: "A recipe.",
		Fields: graphql.Fields{
			"ingredients": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"instructions": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"picture": &graphql.Field{
				Type: graphql.String,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
			"tags": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"thoughts": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	recipeQueryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"searchRecipe": &graphql.Field{
				Type: &graphql.List{
					OfType: recipeType,
				},
				Description: "Search for recipe by name, ingredients, tags, or rating.",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "The name of a recipe dish.",
					},
					"ingredients": &graphql.ArgumentConfig{
						Type: &graphql.List{
							OfType: graphql.String,
						},
						Description: "The ingredients of a recipe dish.",
					},
					"tags": &graphql.ArgumentConfig{
						Type: &graphql.List{
							OfType: graphql.String,
						},
						Description: "The tags of a recipe dish.",
					},
					"rating": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "The rating of a recipe dish.",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rn := utils.GetString(p.Args["name"])
					ings := utils.GetStringSlice(p.Args["ingredients"])
					tags := utils.GetStringSlice(p.Args["tags"])
					rat := utils.GetInt(p.Args["rating"])
					q := models.RecipeQuery(rn, ings, tags, rat)
					rb := mongo.GetCollection("recipebook", "recipes")

					cur, err := rb.Aggregate(context.Background(), q)
					if err != nil {
						return nil, err
					}
					defer cur.Close(context.Background())
					var recipes []models.Recipe

					for cur.Next(context.Background()) {
						elem := models.Recipe{}
						err := cur.Decode(&elem)
						if err != nil {

							return nil, err
						}
						recipes = append(recipes, elem)
					}

					return recipes, nil
				},
			},
		},
	})
)
