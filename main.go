package main

import (
	"TerminalAI/ai"
	"fmt"
)

func main() {

	ai.SetUpAPI()

	fmt.Println(ai.MakeRequest("How are you"))



	// logger.SetupLogger()

	// err := godotenv.Load()

	// if err != nil {
	// 	logger.Logger().Error(err.Error())
	// }

    // ctx := context.Background()
    // client, err := genai.NewClient(ctx, &genai.ClientConfig{
    //     APIKey:  os.Getenv("API_KEY"),
    //     Backend: genai.BackendGeminiAPI,
    // })

	// if err != nil {
	// 	logger.Logger().Error(err.Error())
	// }

	// result, err := client.Models.GenerateContent(
    //     ctx,
    //     "gemini-2.0-flash",
    //     genai.Text("Make a chocolate chip cookie recipe"),
    //     nil,
    // )

	// if err != nil {
	// 	logger.Logger().Error(err.Error())
	// }

	// fmt.Println(result.Text())
}