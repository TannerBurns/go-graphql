package controllers

import (
	"github.com/graphql-go/graphql"

	"../models"
)

/*
Controller - structure to make multiple controllers if needed
*/
type Controller struct {
	Name      string
	GqlSchema *graphql.Schema
	Logger    *models.Logger
}
