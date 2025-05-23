package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const POSTGRES_URL = "postgres://lukefisher:postgres@localhost:5432/grove"

func GetNote(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": recover()})
		}
	}()

	query_id := c.Query("id")
	if query_id == "" {
		log.Panicln("No id query param provided")
	}

	con, err := pgx.Connect(context.Background(), POSTGRES_URL)
	defer con.Close(context.Background())

	if err != nil {
		log.Panicln("Error connecting to the notes database.")
	}

	var created, last_updated pgtype.Timestamp
	var content, title pgtype.Text
	err = con.QueryRow(context.Background(), "SELECT id, created, last_updated, content, title FROM notes where id=$1", query_id).Scan(&query_id, &created, &last_updated, &content, &title)
	if !created.Valid || !last_updated.Valid || !content.Valid || !title.Valid {
		log.Panicln("Internal server error, most likely not your fault.")
	}
	id, _ := strconv.Atoi(query_id)
	
	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"created":      created.Time.UTC(),
		"last_updated": last_updated.Time.UTC(),
		"content":      content.String,
		"title":        title.String,
	})
}
