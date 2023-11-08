package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"udhar/mock"
	"udhar/models"

	"github.com/brianvoe/gofakeit/v6"
)

func LoanApplicationResponseListener(message string) {
	var data mock.LoanApplicationRequest
	msgBytes := []byte(message)
	if err := json.Unmarshal(msgBytes, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
		return
	}
	responseData := mock.LoanApplicationResponse{
		LoanApplicationRequest: data,
	}
	for i := range responseData.LoanApplications {
		responseData.LoanApplications[i].LoanApplicationStatus = gofakeit.RandomString([]string{"PROCESSING", "DECLINED"})
	}
	responseData.Response.Status = "SUCCESS"
	responseData.Response.ResponseDetail = "Request is Valid"

	for i := range responseData.LoanApplications {
		status := fmt.Sprintf("Loan Response Status %v", responseData.LoanApplications[i].LoanApplicationStatus)
		fmt.Println(status)
	}
	log.Println(map[string]string{
		"traceId":   data.Metadata.TraceID.String(),
		"timeStamp": time.Now().Format(time.RFC3339),
		"message":   "Loan Application Response",
	})
}

func ConcentOnAAResponse(message string) {
	var data models.RequestBodyWithTraceId
	msgBytes := []byte(message)
	if err := json.Unmarshal(msgBytes, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
		return
	}
	mockedResponse := mock.MockGenerateConcentRequestByLender(data.TraceId)
	responseData := mock.ConcentProvidedByLender{
		ConcentRequestedByBA: *mockedResponse,
	}
	responseData.Response.Status = "Success"
	responseData.Response.ResponseDetail = "Request to AA Submitted"
	log.Println(map[string]string{
		"traceId":   data.TraceId.String(),
		"timeStamp": time.Now().Format(time.RFC3339),
		"message":   "Concent Response By Lenders",
	})
}

func NotifyBA(message string) {
	var data mock.ConcentStatusRequest
	msgBytes := []byte(message)
	if err := json.Unmarshal(msgBytes, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
		return
	}
	responseData := mock.ConcentStatusResponse{
		ConcentStatusRequest: data,
	}
	responseData.Response.Status = "SUCCESS"
	responseData.Response.ResponseDetail = "Successfully fetched response"
	log.Println(responseData, "Concent Status Response By BA")
}

func GenerateOfferResponse(message string) {
	var data mock.GenerateOfferRequest
	msgBytes := []byte(message)
	if err := json.Unmarshal(msgBytes, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
		return
	}

	d := mock.MockGenerateOffersMetadata()

	responseData := mock.GenerateOfferResponse{
		GenerateOfferRequest:  data,
		GenerateOfferMetadata: *d,
	}

	log.Println(responseData, "Generated offer response sent by Lenders to BA")
}

func SubmitAdditionalDocumentsResponse(message string) {
	var data mock.AdditionalDocumentsRequest
	msgBytes := []byte(message)
	if err := json.Unmarshal(msgBytes, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
		return
	}

	d := mock.MockAdditionDocumentsResponse(data.Metadata.RequestID)

	log.Println(d, "Additional Documents Details Submitted By Lender To BA")

}

func SetOfferHandler(message string) {
	var data mock.GenerateOfferRequest

	msgBytes := []byte(message)
	if err := json.Unmarshal(msgBytes, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v\n", err)
		return
	}

	d := mock.MockGenerateSetOfferResponse(data)
	log.Println(d, "Set Offer Response Sent by BA to Lender")
}
