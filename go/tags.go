package main

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var idTagMap = map[int64]*Tag{}
var nameTagMap = map[string]*Tag{}

func LoadTags(ctx context.Context) error {
	var tagModels []*TagModel
	if err := sqlx.SelectContext(ctx, dbConn, &tagModels, "SELECT * FROM tags"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get tags: "+err.Error())
	}

	for i := range tagModels {
		tag := &Tag{
			ID:   tagModels[i].ID,
			Name: tagModels[i].Name,
		}
		idTagMap[tag.ID] = tag
		nameTagMap[tag.Name] = tag
	}

	return nil
}

func FindTagByID(id int64) *Tag {
	return idTagMap[id]
}

func FindTagByName(name string) *Tag {
	return nameTagMap[name]
}
