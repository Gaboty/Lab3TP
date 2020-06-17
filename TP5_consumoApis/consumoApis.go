package TP5_consumoApis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Item struct  {
	Name               string `json:"name"`
	CountryId          string `json:"country_id"`
	Currencies[] struct{
		Id     string
		Symbol string
	} `json:"currencies"`
}

func ConsumoApi(){
	r := gin.Default()
	resp, err := http.Get("https://api.mercadolibre.com/sites/MLA")
	if err != nil {
		log.Fatal(err)
	}
	var resultado Item
	err = json.NewDecoder(resp.Body).Decode(&resultado)
	if err != nil {
		log.Fatal(err)
	}
	r.GET("api.mercadolibre.com/sites/MLA", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Item": resultado})
	})
	r.Run(":3000")
}
