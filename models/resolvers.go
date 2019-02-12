package models

import "github.com/graphql-go/graphql"

// Resolver struct holds a connection to our database
type Resolver struct {
	db *Database
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) UserResolver(p graphql.ResolveParams) (users interface{}, err error) {
	// Strip the name from arguments and assert that it's a string
	name, ok := p.Args["name"].(string)
	if ok {
		users := r.db.GetUsersByName(name)
		return users, nil
	}

	return
}
