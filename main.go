package main

// TODO: Get only the temperature key and value, parse it to html if possible
import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func main() {
	// & Creation of the listening server
	router_object := gin.Default()

	router_object.GET("/get_weather", func(c *gin.Context) {
		c.String(200, API_Request())
	})
	router_object.Run(":3000") //? Localhost 3000
}

func API_Request() string {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true")
	if err != nil {
		log.Fatalln(err)
	}

	//& Read response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// & Convert the body to type string, filtring it
	json := string(body) // * All the JSON content is in here
	value := gjson.Get(json, "current_weather.temperature")
	return value.String()
}
