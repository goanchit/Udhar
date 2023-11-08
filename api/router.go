package api

import (
	"udhar/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
)

func RouteHandler(g *gin.Engine, awscfg *aws.Config) {
	repository := repository.NewAWSClient(awscfg)
	server := NewServer(repository)

	requestRouter := g.Group("/v1/")
	{
		requestRouter.POST("/createLoanRequest", server.CreateLoanApplication)
		requestRouter.POST("/createConcentOnAA", server.CreateConcentOnAA)
		requestRouter.POST("/getConcentStatus", server.ConcentStatusFromBA)
		requestRouter.POST("/generateOffers", server.GenerateOfferRequest)
		requestRouter.POST("/generateAdditionalDocuments", server.RequestAdditionalDocumentsFromLenders)
		requestRouter.POST("/setOffers", server.SetOfferRequestByBA)
	}
}
