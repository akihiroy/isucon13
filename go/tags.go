package main

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TagModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type TagsResponse struct {
	Tags []*Tag `json:"tags"`
}

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

func getTagHandler(c echo.Context) error {
	tags := make([]*Tag, 0, len(idTagMap))
	for _, tag := range idTagMap {
		tags = append(tags, tag)
	}
	return c.JSON(http.StatusOK, &TagsResponse{
		Tags: tags,
	})
}
