package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"pokeAPI/config"
)

func Pokemon() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query value if there's one in order to fetch a particular pokemon
		params := c.Request.URL.Query()
		offset := params.Get("offset")
		limit := params.Get("limit")

		// A variable to store the params that will be appended to the query
		queryParams := "/?"

		if offset != "" {
			queryParams = queryParams + "offset=" + offset + "&"
		}
		if limit != "" {
			queryParams = queryParams + "limit=" + limit
		}

		// Make the request to the actual PokeAPI with a not so reliable method...
		response, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon%s", queryParams))
		if err != nil {
			log.Fatalln(err)
			return
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
			return
		}

		var result config.PokeResponse
		json.Unmarshal(body, &result)

		// Remove the real API url from the next and previous keys to direct response to my server
		result.Next = strings.ReplaceAll(result.Next, "https://pokeapi.co/api/v2/pokemon", "")
		result.Previous = strings.ReplaceAll(result.Previous, "https://pokeapi.co/api/v2/pokemon", "")

		// Resurn of the result
		c.JSON(http.StatusOK, result)
	}
}
