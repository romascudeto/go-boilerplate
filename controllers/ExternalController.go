package controllers

import (
	"encoding/json"
	"go-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

//GetExternal1 ... Get data from external
func GetExternal1(c *gin.Context) {
	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get("https://private-2489f-roommeapi.apiary-mock.com/v2/config/checkoutLabel")

	var restyResponse helpers.Response
	json.Unmarshal(resp.Body(), &restyResponse)
	response := helpers.NewResponse(restyResponse.Meta, restyResponse.Data)
	c.JSON(http.StatusOK, response)
}

//GetExternal2 ... Get data from external
func GetExternal2(c *gin.Context) {
	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get("https://jsonplaceholder.typicode.com/todos")

	byt := []byte(resp.Body())
	var dat interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	meta := map[string]interface{}{}
	response := helpers.NewResponse(meta, dat)
	c.JSON(http.StatusOK, response)
}
