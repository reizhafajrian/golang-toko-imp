package main

import (
	"Test/models"
	"Test/provider"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var apiV1 = "/api/v1"
var saveData = "/save-data"

type RequestBody struct {
	Data []models.JsonToko `json:"data"`
}

func main() {
	collection := provider.ConnectingMongo().Database("tokopedia").Collection("toko")
	r := gin.Default()
	api := r.Group(apiV1)
	api.POST(saveData, func(c *gin.Context) {
		var reqBody RequestBody
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, "Error reading request body")
			return
		}
		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  "error",
					"message": "Error parsing request body",
				})
			return
		}
		for _, dataToko := range reqBody.Data {
			_, err := collection.InsertOne(context.Background(), dataToko)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusCreated, gin.H{
			"status":      "success",
			"status_code": http.StatusCreated,
		})
	})
	r.Run()
}
