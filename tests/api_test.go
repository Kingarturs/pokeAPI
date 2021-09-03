package tests

import (
	"pokeAPI/config"
	"pokeAPI/routes"
	"testing"

	test_utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/", routes.Pokemon())
	router.GET("/hello", routes.HelloWorld())

	test_utils.SetRouter(router)
}

func TestPokemonWithParams(t *testing.T) {
	t.Log("/ with offset and limit should return a valid result")

	var result config.PokeResponse
	params := map[string]int{"offset": 20, "limit": 20}

	err := test_utils.TestHandlerUnMarshalResp("GET", "/", "form", params, &result)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}

	if result.Count == 0 {
		t.Errorf("Result count empty")
	}
	if result.Next == "" {
		t.Errorf("Next shouldn't be empty")
	}
	if result.Previous == "" {
		t.Errorf("Previous shouldn't be empty")
	}
	if len(result.Results) == 0 {
		t.Errorf("Results shouldn't be empty")
	}
}

func TestPokemonWithoutParams(t *testing.T) {
	t.Log("/ without params should return a valid result")

	var result config.PokeResponse

	err := test_utils.TestHandlerUnMarshalResp("GET", "/", "form", "", &result)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}

	if result.Count == 0 {
		t.Errorf("Result count empty")
	}
	if result.Next == "" {
		t.Errorf("Next shouldn't be empty")
	}
	if result.Previous != "" {
		t.Errorf("Previous should be empty")
	}
	if len(result.Results) == 0 {
		t.Errorf("Results shouldn't be empty")
	}
}

func TestHelloWorldHandler(t *testing.T) {
	t.Log("/hello should return Hello, World")

	var result string

	body, err := test_utils.TestOrdinaryHandler("GET", "/hello", "form", "")
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}

	result = string(body)

	if result != "Hello, World" {
		t.Errorf("Result should be Hello, World")
	}
}
