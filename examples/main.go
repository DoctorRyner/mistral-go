package main

import (
	"log"

	"github.com/gage-technologies/mistral-go"
)

func main() {
	// If api key is empty it will load from MISTRAL_API_KEY env var
	client := mistral.NewMistralClientDefault("your-api-key")

	// Example: Using Chat Completions
	chatRes, err := client.Chat("mistral-tiny", []mistral.ChatMessage{{Content: "Hello, world!", Role: mistral.RoleUser}}, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion: %v", err)
	}
	log.Printf("Chat completion: %+v\n", chatRes)

	// Example: Using Chat Completions Stream
	chatResChan, err := client.ChatStream("mistral-tiny", []mistral.ChatMessage{{Content: "Hello, world!", Role: mistral.RoleUser}}, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion stream: %v", err)
	}

	for chatResChunk := range chatResChan {
		if chatResChunk.Error != nil {
			log.Fatalf("Error while streaming response: %v", chatResChunk.Error)
		}
		log.Printf("Chat completion stream part: %+v\n", chatResChunk)
	}

	// Example: Using Embeddings
	embsRes, err := client.Embeddings("mistral-embed", []string{"Embed this sentence.", "As well as this one."})
	if err != nil {
		log.Fatalf("Error getting embeddings: %v", err)
	}

	log.Printf("Embeddings response: %+v\n", embsRes)

	// Example: Using Agent Chat Completions
	agentReqParam := mistral.DefaultChatRequestParams
	agentReqParam.AgentId = "your-agent-id"
	agentRes, err := client.Chat("mistral-tiny", []mistral.ChatMessage{{Content: "Hello, world!", Role: mistral.RoleUser}}, &agentReqParam)
	if err != nil {
		log.Fatalf("Error getting chat completion: %v", err)
	}
	log.Printf("Agent chat completion: %+v\n", agentRes)

	// Example: Using Agent Chat Completions Stream
	agentStreamReqParam := mistral.DefaultChatRequestParams
	agentStreamReqParam.AgentId = "your-agent-id"
	agentResChan, err := client.ChatStream("mistral-tiny", []mistral.ChatMessage{{Content: "Hello, world!", Role: mistral.RoleUser}}, &agentStreamReqParam)
	if err != nil {
		log.Fatalf("Error getting chat completion stream: %v", err)
	}
	for agentResChunk := range agentResChan {
		if agentResChunk.Error != nil {
			log.Fatalf("Error while streaming response: %v", agentResChunk.Error)
		}
		log.Printf("Agent chat completion stream part: %+v\n", agentResChunk)
	}
}
