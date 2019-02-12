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

	// Check to ensure query was provided in the request body
	if r.Body == nil {
		http.Error(w, "Must provide graphql query in request body", 400)
		return
	}

	var rBody reqBody
	// Decode the request body into rBody
	err := json.NewDecoder(r.Body).Decode(&rBody)
	if err != nil {
		http.Error(w, "Error parsing JSON request body", 400)
	}

	// Execute graphql query
	result := models.ExecuteQuery(rBody.Query, *c.GqlSchema)

	// marshalling to json,
	// the Content-Type as application/json.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(r, 200)
	json.NewEncoder(w).Encode(result)
	return
}
