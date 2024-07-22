package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
)

// GraphQLRequest represents a typical GraphQL request payload
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
}

// ReadGraphQLRequest reads the GraphQL request from the specified file
func ReadGraphQLRequest(filePath string) (GraphQLRequest, error) {
	var graphqlRequest GraphQLRequest

	// Read the file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return graphqlRequest, err
	}

	// Unmarshal JSON content into GraphQLRequest struct
	err = json.Unmarshal(data, &graphqlRequest)
	return graphqlRequest, err
}

// EncodeToFormURLEncoded encodes the GraphQL request to x-www-form-urlencoded format
func EncodeToFormURLEncoded(graphqlRequest GraphQLRequest) (string, error) {
	formData := url.Values{}

	// Encode each part of the GraphQL request
	formData.Set("query", graphqlRequest.Query)
	if graphqlRequest.OperationName != "" {
		formData.Set("operationName", graphqlRequest.OperationName)
	}
	if graphqlRequest.Variables != nil {
		variablesJSON, err := json.Marshal(graphqlRequest.Variables)
		if err != nil {
			return "", err
		}
		formData.Set("variables", string(variablesJSON))
	}

	return formData.Encode(), nil
}

func main() {
	// Define a command-line flag for the input file
	filePath := flag.String("file", "", "Path to the file containing the GraphQL request")
	flag.Parse()

	// Check if the file path is provided
	if *filePath == "" {
		fmt.Println("Usage: go run main.go -file <path_to_graphql_request_file>")
		os.Exit(1)
	}

	// Read the GraphQL request from the file
	graphqlRequest, err := ReadGraphQLRequest(*filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Encode the GraphQL request into x-www-form-urlencoded format
	encodedQuery, err := EncodeToFormURLEncoded(graphqlRequest)
	if err != nil {
		fmt.Println("Error encoding the query:", err)
		os.Exit(1)
	}

	// Print the encoded query
	fmt.Println(encodedQuery)
}
