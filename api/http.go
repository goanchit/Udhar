package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"udhar/mock"
	"udhar/models"
	"udhar/repository"

	"github.com/gin-gonic/gin"
)

type Server struct {
	repository repository.Repository
}

func NewServer(repo repository.Repository) *Server {
	return &Server{
		repository: repo,
	}
}

func (s Server) CreateLoanApplication(c *gin.Context) {

	generateProduct := mock.MockCreateProduct()

	// Creates 2 Loan Application
	mockLoanApplication := mock.MockGenerateMultipleLoanApplicationData(generateProduct)
	fmt.Println(mockLoanApplication.LoanApplications[0].LoanApplicationStatus)
	fmt.Println(mockLoanApplication.LoanApplications[1].LoanApplicationStatus)
	loanApplicationStringified, err := json.Marshal(mockLoanApplication)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to marshall data %v", err)
		log.Fatalln(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
	}

	s.repository.PublishToTopic(context.Background(), loanApplicationStringified, "arn:aws:sns:us-east-1:000000000000:CREATE-LOAN-APPLICATION")
	c.JSON(http.StatusOK, gin.H{
		"message": "Loan Application Request Sent",
		"data": map[string]interface{}{
			"traceId":   mockLoanApplication.Metadata.TraceID,
			"timeStamp": time.Now().Format(time.RFC3339),
			"error":     nil,
		},
	})
}

func (s Server) CreateConcentOnAA(c *gin.Context) {
	body := models.RequestBodyWithTraceId{}

	if err := c.ShouldBind(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Body Parsed",
		})
		return
	}

	r, err := json.Marshal(body)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to marshall data %v", err)
		log.Fatalln(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
	}

	s.repository.PublishToTopic(context.Background(), r, "arn:aws:sns:us-east-1:000000000000:CREATE-CONCENT-ON-AA")
	c.JSON(http.StatusOK, gin.H{
		"message": "Lenders have been notified to generate data",
		"data": map[string]interface{}{
			"traceId":   body.TraceId,
			"timeStamp": time.Now().Format(time.RFC3339),
			"error":     nil,
		},
	})
}

func (s Server) ConcentStatusFromBA(c *gin.Context) {
	body := models.RequestBodyWithTraceId{}

	if err := c.ShouldBind(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Body Parsed",
		})
		return
	}
	consentPayload := mock.MockGenerateConcentRequestStatusData(body.TraceId)

	r, err := json.Marshal(consentPayload)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to marshall data %v", err)
		log.Fatalln(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
	}
	s.repository.PublishToTopic(context.Background(), r, "arn:aws:sns:us-east-1:000000000000:NOTIFY-BA")

	c.JSON(http.StatusOK, gin.H{
		"message": "Notification has been sent to BA",
		"data": map[string]interface{}{
			"traceId":   body.TraceId,
			"timeStamp": time.Now().Format(time.RFC3339),
			"error":     nil,
		},
	})
}

func (s Server) GenerateOfferRequest(c *gin.Context) {
	body := models.RequestBodyWithTraceId{}

	if err := c.ShouldBind(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Body Parsed",
		})
		return
	}

	mockOfferRequest := mock.MockGenerateOffersRequest(body.TraceId)

	r, err := json.Marshal(mockOfferRequest)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to marshall data %v", err)
		log.Fatalln(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
	}
	s.repository.PublishToTopic(context.Background(), r, "arn:aws:sns:us-east-1:000000000000:GENERATE-OFFERS")
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification sent to Lenders to generate offers",
		"data": map[string]interface{}{
			"traceId":   body.TraceId,
			"timeStamp": time.Now().Format(time.RFC3339),
			"error":     nil,
		},
	})

}

func (s Server) RequestAdditionalDocumentsFromLenders(c *gin.Context) {
	body := models.RequestBodyWithTraceId{}

	if err := c.ShouldBind(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Body Parsed",
		})
		return
	}

	mockAdditionalData := mock.MockAdditionDocumentsRequest(body.TraceId)

	r, err := json.Marshal(mockAdditionalData)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to marshall data %v", err)
		log.Fatalln(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
	}

	s.repository.PublishToTopic(context.Background(), r, "arn:aws:sns:us-east-1:000000000000:SUBMIT-ADDITIONAL-DOCUMENTS")
	c.JSON(http.StatusOK, gin.H{
		"message": "Additional Documents Submit Request Sent To Lenders",
		"data": map[string]interface{}{
			"traceId":   body.TraceId,
			"timeStamp": time.Now().Format(time.RFC3339),
			"error":     nil,
		},
	})
}

func (s Server) SetOfferRequestByBA(c *gin.Context) {
	body := models.RequestBodyWithTraceId{}
	if err := c.ShouldBind(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Body Parsed",
		})
		return
	}

	mockSetOfferRequest := mock.MockGenerateSetOfferRequest(body.TraceId)
	r, err := json.Marshal(mockSetOfferRequest)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to marshall data %v", err)
		log.Fatalln(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
	}

	s.repository.PublishToTopic(context.Background(), r, "arn:aws:sns:us-east-1:000000000000:SET-OFFER")
	c.JSON(http.StatusOK, gin.H{
		"message": "Sent Notification to Lender With The Selected Offer",
		"data": map[string]interface{}{
			"traceId":   body.TraceId,
			"timeStamp": time.Now().Format(time.RFC3339),
			"error":     nil,
		},
	})

}
