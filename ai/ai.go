package ai

import (
	"google.golang.org/genai"
	"github.com/joho/godotenv"
	"os"
	"TerminalAI/logger"
	"sync"
)

var client genai.Client
var ctx context.Context


	init := sync.OnceFunc(
		func() error {
			defer return nil
			logger.SetupLogger()

			err := godotenv.Load()

			if err != nil {
				logger.Logger().Error(err.Error())
				return err
			}

			ctx := context.Background()
			client, err := genai.NewClient(ctx, &genai.ClientConfig{
				APIKey:  os.Getenv("API_KEY"),
				Backend: genai.BackendGeminiAPI,
			})

			if err != nil {
				logger.Logger().Error(err.Error())
				return err
			}
		}
	)

func SetUpAPI() error {
	init()
}


func MakeRequest(prompt string) string{	
	result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.0-flash",
        genai.Text(prompt),
        nil,
    )

	if err != nil {
		logger.Logger().Error(err.Error())
	}

	return result.Text()
}