# Simple PokeAPI made with Go
This API works just as a sort of intermediary between the PokeAPI found in https://pokeapi.co/ and it's made for Wizeline Academy's Go Bootcamp

### General information
- Go Language version used: `1.15.6`
- Framework used: `Gin Gonic`

### How to run it
1. Clone this repo.
2. Copy the `.env.example` file into a `.env` file
3. Replace the values of the variables in the `.env` file with the values required
4. Install the dependencies (This project is made with Go, so you will need to have it installed on your machine) using the command `go mod download`
5. Navigate to the main folder and run the project with `go run main.go`

### How to run the tests
1. Navigate to the tests folder
2. Run the command `go test` on your terminal

---

### Endpoints
Required endpoints for the challenge

#### Hello world
```GET /hello```
This endpoint returns a simple string with a "Hello, World"

#### Pokemon list
```GET /```

- query parameters
  - `offset` - Same as the original PokeAPI is the offset of the queried pokemon
  - `limit` - The size limit of the array of results

This endpoint returns exactly the same response structure as the original PokeAPI, it returns the following info:
- `count` - The number of results in the query
- `next` - The url for the next page of results
- `previuos` - The url for the previous page of results
- `results` - An array containing the actual list of pokemon
