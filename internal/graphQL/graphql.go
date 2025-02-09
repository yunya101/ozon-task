package graphstruct

import (
	"github.com/graphql-go/graphql"
	"github.com/yunya101/ozon-task/internal/service"
)

var PostType *graphql.Object
var CommentType *graphql.Object
var UserType *graphql.Object

type GraphQlQueries struct {
	service *service.PostService
}

func (g *GraphQlQueries) InitPostType() *graphql.Object {
	PostType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Post",
			Fields: graphql.Fields{
				"id":            &graphql.Field{Type: graphql.Int},
				"author":        &graphql.Field{Type: graphql.Int},
				"title":         &graphql.Field{Type: graphql.String},
				"text":          &graphql.Field{Type: graphql.String},
				"subs":          &graphql.Field{Type: graphql.NewList(UserType)},
				"comments":      &graphql.Field{Type: graphql.NewList(CommentType)},
				"countComments": &graphql.Field{Type: graphql.Int},
				"isCommented":   &graphql.Field{Type: graphql.Boolean},
			},
		},
	)
	return PostType
}

func (g *GraphQlQueries) InitUserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":       &graphql.Field{Type: graphql.Int},
				"username": &graphql.Field{Type: graphql.String},
			},
		},
	)
}
func (g *GraphQlQueries) InitCommentType() *graphql.Object {
	CommentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"id":       &graphql.Field{Type: graphql.Int},
				"author":   &graphql.Field{Type: graphql.Int},
				"text":     &graphql.Field{Type: graphql.String},
				"post":     &graphql.Field{Type: graphql.Int},
				"patent":   &graphql.Field{Type: graphql.Int},
				"comments": &graphql.Field{Type: graphql.NewList(CommentType)},
				"createAt": &graphql.Field{Type: graphql.DateTime},
			},
		},
	)
	return CommentType
}

func (g *GraphQlQueries) GetPostById(postType *graphql.Object, page int) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"post": &graphql.Field{Type: graphql.NewList(postType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						posts, err := g.service.GetLastestPosts(page)
						if err != nil {
							return nil, err
						}
						return posts, nil
					},
				},
			},
		},
	)
}
