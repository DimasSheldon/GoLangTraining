package controller

import (
	"assignment3/model"
	"assignment3/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetStatus(ctx *gin.Context) {
	util.UpdateJSONData()

	jsonData, err := os.Open("data.json")

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	bytes, err := ioutil.ReadAll(jsonData)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	var data model.Status

	json.Unmarshal(bytes, &data)

	waterStatus, windStatus := util.GetWaterStatus(data.Status.Water), util.GetWindStatus(data.Status.Wind)
	waterClass, winedClass := util.GetWaterClass(data.Status.Water), util.GetWindClass(data.Status.Wind)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"water":       data.Status.Water,
		"waterStatus": waterStatus,
		"waterClass":  waterClass,
		"wind":        data.Status.Wind,
		"windStatus":  windStatus,
		"windClass":   winedClass,
	})
}
