package models

import (
	"github.com/mongodb/mongo-go-driver/bson"
	// "github.com/mongodb/mongo-go-driver/x/bsonx"
	"fmt"
)

// Recipe struct to model a recipe in mongo.
type Recipe struct {
	Ingredients  []string `json:"ingredients" bson:"ingredients"`
	Instructions []string `json:"instructions" bson:"instructions"`
	Name         string   `json:"name" bson:"name"`
	Picture      string   `json:"picture" bson:"picture"`
	Rating       int      `json:"rating" bson:"rating"`
	Tags         []string `json:"tags" bson:"tags"`
	Thoughts     string   `json:"thoughts" bson:"thoughts"`
}

func RecipeQuery(name string, ingredients, tags []string, rating int) []interface{} {
	var matches []interface{}
	var agg []interface{}

	if name != "" {
		matches = append(matches, bson.M{"name": name})
	}

	if len(ingredients) > 0 {
		agg = append(agg, bson.M{"$unwind": "$ingredients"})
		agg = append(agg, bson.M{"$group": bson.M{
			"_id":          "$_id",
			"ingredients":  bson.M{"$push": "$ingredients"},
			"instructions": bson.M{"$first": "$instructions"},
			"name":         bson.M{"$first": "$name"},
			"picture":      bson.M{"$first": "$picture"},
			"rating":       bson.M{"$first": "$rating"},
			"tags":         bson.M{"$first": "$tags"},
			"thoughts":     bson.M{"$first": "$thoughts"},
		}})
		for _, ingredient := range ingredients {
			matches = append(matches, bson.M{"ingredients": ingredient})
		}
	}

	if len(tags) > 0 {
		fmt.Println("tags", tags)
		agg = append(agg, bson.M{"$unwind": "$tags"})
		agg = append(agg, bson.M{"$group": bson.M{
			"_id":          "$_id",
			"ingredients":  bson.M{"$first": "$ingredients"},
			"instructions": bson.M{"$first": "$instructions"},
			"name":         bson.M{"$first": "$name"},
			"picture":      bson.M{"$first": "$picture"},
			"rating":       bson.M{"$first": "$rating"},
			"tags":         bson.M{"$push": "$tags"},
			"thoughts":     bson.M{"$first": "$thoughts"},
		}})
		for _, tag := range tags {
			matches = append(matches, bson.M{"tags": tag})
		}
	}

	if rating > 0 && rating < 6 {
		matches = append(matches, bson.M{"rating": rating})
	}

	agg = append(agg,
		bson.M{"$match": bson.D{{"$and", matches}}},
	)

	fmt.Println("agg", agg)
	return agg
}
