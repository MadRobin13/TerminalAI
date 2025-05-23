package ai

import (
	"TerminalAI/logger"
	"context"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

var client *genai.Client = nil
var ctx context.Context = context.Background()
var once sync.Once

func SetUpAPI() {
	once.Do(func() {
		logger.SetupLogger()

		err := godotenv.Load("./.env")

		if err != nil {
			logger.Logger().Error(err.Error())
			panic(err)
		}

		client, err = genai.NewClient(ctx, &genai.ClientConfig{
			APIKey:  os.Getenv("API_KEY"),
			Backend: genai.BackendGeminiAPI,
		})

		if err != nil {
			logger.Logger().Error(err.Error())
			panic(err)
		}
	})
}


func MakeRequest(prompt string) string{	
	result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.0-flash",
        genai.Text("You are going to generate a respose based on this current chat history between a user and yourself. Do what the user says and speak directly to the user at all times: \n" + prompt),
        nil,
    )

	if err != nil {
		logger.Logger().Error(err.Error())
	}

	return result.Text()
}