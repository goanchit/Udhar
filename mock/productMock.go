package mock

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func MockCreateProduct() *Product {
	product := Product{}
	gofakeit.Struct(&product)
	return &product
}

func MockGenerateMultipleLoanApplicationData(product *Product) *LoanApplicationRequest {
	loanApplicationrequest := LoanApplicationRequest{}
	gofakeit.Struct(&loanApplicationrequest)
	loanApplicationrequest.ProductData.ProductID = product.Data.Id
	loanApplicationrequest.ProductData.ProductNetworkID = product.Network.Id
	return &loanApplicationrequest
}

func MockGenerateConcentRequestByLender(traceId uuid.UUID) *ConcentRequestedByBA {
	data := ConcentRequestedByBA{}
	gofakeit.Struct(&data)
	data.Metadata.TraceID = traceId
	return &data
}

func MockGenerateConcentRequestStatusData(traceId uuid.UUID) *ConcentStatusRequest {
	data := ConcentStatusRequest{}
	gofakeit.Struct(&data)
	data.Metadata.TraceID = traceId
	return &data
}

func MockGenerateOffersRequest(traceId uuid.UUID) *GenerateOfferRequest {
	data := GenerateOfferRequest{}
	gofakeit.Struct(&data)
	data.Metadata.RequestID = traceId
	return &data
}

func MockGenerateOffersMetadata() *GenerateOfferMetadata {
	data := GenerateOfferMetadata{}
	gofakeit.Struct(&data)
	return &data
}

func MockAdditionDocumentsRequest(traceId uuid.UUID) *AdditionalDocumentsRequest {
	data := AdditionalDocumentsRequest{}
	gofakeit.Struct(&data)
	data.Metadata.TraceID = traceId
	return &data
}

func MockAdditionDocumentsResponse(traceId uuid.UUID) *AdditionalDocumentResponse {
	data := AdditionalDocumentResponse{}
	gofakeit.Struct(&data)
	data.Metadata.TraceID = traceId
	return &data
}

func MockGenerateSetOfferRequest(traceId uuid.UUID) *SetOfferRequest {
	data := SetOfferRequest{}
	gofakeit.Struct(&data)
	data.Metadata.TraceID = traceId
	return &data
}

func MockGenerateSetOfferResponse(requestData GenerateOfferRequest) *SetOfferResponse {
	data := SetOfferResponse{}
	gofakeit.Struct(&data)
	data.Metadata = requestData.Metadata
	data.ProductData = requestData.ProductData
	data.LoanApplicationStatus = "CREATED"
	return &data
}
