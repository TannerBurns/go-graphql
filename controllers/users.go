package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
)

type reqBody struct {
	Query string `json:"query"`
}

func (c *Controller) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := r.URL.Query().Get("query")
	// Check to ensure query was provided in the request body
	if query == "" {
		c.Logger.Logging(r, 400)
		http.Error(w, "Must provide graphql query in params", 400)
		return
	}

	// Execute graphql query
	result := models.ExecuteQuery(query, *c.GqlSchema)

	// marshalling to json,
	// the Content-Type as application/json.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(r, 200)
	json.NewEncoder(w).Encode(result)
	return
}
