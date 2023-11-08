package main

import (
	"fmt"
	"log"
	"udhar/api"
	"udhar/config"
	"udhar/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	cfg := config.ConnectToAWS() // Establish AWS Connection

	// Setup SNS and SQS
	sqsClient, _, queueUrls := config.SetupInfra(cfg)

	api.RouteHandler(r, cfg)

	channel := make(chan map[string]interface{})

	for k, v := range queueUrls {
		if v != "" {
			go handler.PollMessages(sqsClient, k, channel)
		}
	}
	go func() {
		for message := range channel {
			queueUrl, _ := message["queueUrl"].(string)
			msg, msgExists := message["message"].(string)
			if msgExists {
				switch queueUrls[queueUrl] {
				case "CREATE-LOAN-APPLICATION":
					handler.LoanApplicationResponseListener(msg)
				case "CREATE-CONCENT-ON-AA":
					handler.ConcentOnAAResponse(msg)
				case "NOTIFY-BA":
					handler.NotifyBA(msg)
				case "GENERATE-OFFERS":
					handler.GenerateOfferResponse(msg)
				case "SUBMIT-ADDITIONAL-DOCUMENTS":
					handler.SubmitAdditionalDocumentsResponse(msg)
				case "SET-OFFER":
					handler.SetOfferHandler(msg)
				default:
					fmt.Print("No Queue Matched")
				}

			}
		}
	}()

	r.Run(":8000")
}
