package gql

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

var CartList []Cart

type Cart struct {
	Sku  string `json:"sku"`
	Qty  int    `json:"qty"`
}

// define custom GraphQL ObjectType `cartType` for our Golang struct `Cart`
// Note that
// - the fields in our cartType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var cartType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Cart",
	Fields: graphql.Fields{
		"sku": &graphql.Field{
			Type: graphql.String,
		},
		"qty": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addCart": &graphql.Field{
			Type:        cartType, // the return type for this field
			Description: "Add new cart",
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"qty": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				sku, _ := params.Args["sku"].(string)
				qty,_  := params.Args["qty"].(int)


				// perform mutation operation here
				// for e.g. create a Cart and save to DB.
				newCart := Cart{
					Sku: sku,
					Qty: qty,
				}

				CartList = append(CartList, newCart)

				// return the new Cart object that we supposedly save to DB
				// Note here that
				// - we are returning a `Cart` struct instance here
				// - we previously specified the return Type to be `cartType`
				// - `Cart` struct maps to `cartType`, as defined in `cartType` ObjectConfig`
				return newCart, nil
			},
		},
	},
})

// root query
// we just define a trivial example here, since root query is required.
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		"cart": &graphql.Field{
			Type:        cartType,
			Description: "Get single cart",
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"qty": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery, isOK := params.Args["sku"].(string)
				if isOK {
					// Search for el with sku
					for _, todo := range CartList {
						if todo.Sku == idQuery {
							return todo, nil
						}
					}
				}

				return Cart{}, nil
			},
		},

		"lastCart": &graphql.Field{
			Type:        cartType,
			Description: "Last cart added",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return CartList[len(CartList)-1], nil
			},
		},

		"cartList": &graphql.Field{
			Type:        graphql.NewList(cartType),
			Description: "List of cart",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return cartType, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var CartSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func GqlExecuteQuery(c echo.Context) *graphql.Result {
	var p postData

	json.NewDecoder(c.Request().Body).Decode(&p)

	result := graphql.Do(graphql.Params{
		Context:        c.Request().Context(),
		Schema:         CartSchema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}

	return result
}