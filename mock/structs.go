package mock

import "github.com/google/uuid"

type Product struct {
	Data struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"data"`
	Network struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	} `json:"network"`
}

type Metadata struct {
	Version                 string    `json:"version"`
	OriginatorOrgID         string    `json:"originatorOrgId"`
	OriginatorParticipantID string    `json:"originatorParticipantId"`
	Timestamp               string    `json:"timestamp"`
	TraceID                 uuid.UUID `json:"traceId"`
	RequestID               uuid.UUID `json:"requestId"`
}

type ProductData struct {
	ProductID        uuid.UUID `json:"productId"`
	ProductNetworkID uuid.UUID `json:"productNetworkId"`
}

type LoanApplication struct {
	CreatedDate           string    `json:"createdDate"`
	LoanApplicationID     uuid.UUID `json:"loanApplicationId"`
	LoanApplicationStatus string    `json:"loanApplicationStatus" fake:"{randomstring:[CREATED]}"`
	PledgedDocuments      []struct {
		Source           string `json:"source"`
		SourceIdentifier string `json:"sourceIdentifier"`
		Format           string `json:"format"`
		Reference        string `json:"reference"`
		Type             string `json:"type"`
		IsDataInline     bool   `json:"isDataInline"`
		Data             string `json:"data"`
	} `json:"pledgedDocuments"`
	Borrower struct {
		PrimaryID             uuid.UUID `json:"primaryId"`
		PrimaryIDType         string    `json:"primaryIdType"`
		AdditionalIdentifiers []struct {
			Key            string `json:"key"`
			Value          string `json:"value"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"additionalIdentifiers"`
		Name           string `json:"name"`
		Category       string `json:"category"`
		ContactDetails []struct {
			Type        string `json:"type"`
			Desctiption string `json:"desctiption"`
			Phone       string `json:"phone"`
			Email       string `json:"email"`
			Address     struct {
				Hba            string `json:"hba"`
				Srl            string `json:"srl"`
				Landmark       string `json:"landmark"`
				Als            string `json:"als"`
				Vtc            string `json:"vtc"`
				PinCode        string `json:"pinCode"`
				Po             string `json:"po"`
				District       string `json:"district"`
				State          string `json:"state"`
				Country        string `json:"country"`
				URL            string `json:"url"`
				Latitude       string `json:"latitude"`
				Longitude      string `json:"longitude"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"address"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"contactDetails"`
		Documents []struct {
			Source           string `json:"source"`
			SourceIdentifier string `json:"sourceIdentifier"`
			Format           string `json:"format"`
			Reference        string `json:"reference"`
			Type             string `json:"type"`
			IsDataInline     bool   `json:"isDataInline"`
			Data             string `json:"data"`
		} `json:"documents"`
		URL            string `json:"url"`
		ExtensibleData struct {
		} `json:"extensibleData"`
		AaIdentifier string `json:"aaIdentifier"`
	} `json:"borrower"`
	Guarantors []struct {
		PrimaryID             uuid.UUID `json:"primaryId"`
		PrimaryIDType         string    `json:"primaryIdType"`
		Description           string    `json:"description"`
		AdditionalIdentifiers []struct {
			Key            string `json:"key"`
			Value          string `json:"value"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"additionalIdentifiers"`
		Name                     string `json:"name"`
		RelationshipWithBorrower string `json:"relationshipWithBorrower"`
		Category                 string `json:"category"`
		ContactDetails           []struct {
			Type        string `json:"type"`
			Desctiption string `json:"desctiption"`
			Phone       string `json:"phone"`
			Email       string `json:"email"`
			Address     struct {
				Hba            string `json:"hba"`
				Srl            string `json:"srl"`
				Landmark       string `json:"landmark"`
				Als            string `json:"als"`
				Vtc            string `json:"vtc"`
				PinCode        string `json:"pinCode"`
				Po             string `json:"po"`
				District       string `json:"district"`
				State          string `json:"state"`
				Country        string `json:"country"`
				URL            string `json:"url"`
				Latitude       string `json:"latitude"`
				Longitude      string `json:"longitude"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"address"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"contactDetails"`
		Documents []struct {
			Source           string `json:"source"`
			SourceIdentifier string `json:"sourceIdentifier"`
			Format           string `json:"format"`
			Reference        string `json:"reference"`
			Type             string `json:"type"`
			IsDataInline     bool   `json:"isDataInline"`
			Data             string `json:"data"`
		} `json:"documents"`
		URL            string `json:"url"`
		ExtensibleData struct {
		} `json:"extensibleData"`
	} `json:"guarantors"`
	Applicants []struct {
		PrimaryID             string `json:"primaryId"`
		PrimaryIDType         string `json:"primaryIdType"`
		AdditionalIdentifiers []struct {
			Key            string `json:"key"`
			Value          string `json:"value"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"additionalIdentifiers"`
		Name                     string `json:"name"`
		Category                 string `json:"category"`
		RelationshipWithBorrower string `json:"relationshipWithBorrower"`
		ContactDetails           []struct {
			Type        string `json:"type"`
			Desctiption string `json:"desctiption"`
			Phone       string `json:"phone"`
			Email       string `json:"email"`
			Address     struct {
				Hba            string `json:"hba"`
				Srl            string `json:"srl"`
				Landmark       string `json:"landmark"`
				Als            string `json:"als"`
				Vtc            string `json:"vtc"`
				PinCode        string `json:"pinCode"`
				Po             string `json:"po"`
				District       string `json:"district"`
				State          string `json:"state"`
				Country        string `json:"country"`
				URL            string `json:"url"`
				Latitude       string `json:"latitude"`
				Longitude      string `json:"longitude"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"address"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"contactDetails"`
		Documents []struct {
			Source           string `json:"source"`
			SourceIdentifier string `json:"sourceIdentifier"`
			Format           string `json:"format"`
			Reference        string `json:"reference"`
			Type             string `json:"type"`
			IsDataInline     bool   `json:"isDataInline"`
			Data             string `json:"data"`
		} `json:"documents"`
		URL            string `json:"url"`
		ExtensibleData struct {
		} `json:"extensibleData"`
	} `json:"applicants"`
	Terms struct {
		RequestedAmount      int    `json:"requestedAmount"`
		Currency             string `json:"currency"`
		SanctionedAmount     int    `json:"sanctionedAmount"`
		NetDisbursedAmount   int    `json:"netDisbursedAmount"`
		InterestType         string `json:"interestType"`
		InterestRate         int    `json:"interestRate"`
		AnnualPercentageRate int    `json:"annualPercentageRate"`
		CoolingOffPeriod     struct {
			Duration int    `json:"duration"`
			Unit     string `json:"unit"`
		} `json:"coolingOffPeriod"`
		TotalRepayableAmount int    `json:"totalRepayableAmount"`
		InterestAmount       int    `json:"interestAmount"`
		Description          string `json:"description"`
		Tenure               struct {
			Duration int    `json:"duration"`
			Unit     string `json:"unit"`
		} `json:"tenure"`
		LegalAgreement struct {
			Type string `json:"type"`
			Data string `json:"data"`
		} `json:"legalAgreement"`
		Documents []struct {
			Source           string `json:"source"`
			SourceIdentifier string `json:"sourceIdentifier"`
			Format           string `json:"format"`
			Reference        string `json:"reference"`
			Type             string `json:"type"`
			IsDataInline     bool   `json:"isDataInline"`
			Data             string `json:"data"`
		} `json:"documents"`
		Charges struct {
			Prepayment struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"prepayment"`
			Bounce struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"bounce"`
			LatePayment struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"latePayment"`
			Processing struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"processing"`
			Other struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"other"`
		} `json:"charges"`
		URL            string `json:"url"`
		ExtensibleData struct {
		} `json:"extensibleData"`
	} `json:"terms"`
	URL              string `json:"url"`
	Description      string `json:"description"`
	CreditGuarantees struct {
		CgtmseEligible bool `json:"cgtmseEligible"`
	} `json:"creditGuarantees"`
}

type LoanApplicationRequest struct {
	Metadata         Metadata          `json:"metadata"`
	ProductData      ProductData       `json:"productData"`
	LoanApplications []LoanApplication `json:"loanApplications" fakesize:"2"`
}

type Response struct {
	Status         string `json:"status"`
	ResponseDetail string `json:"responseDetail"`
}

type LoanApplicationResponse struct {
	LoanApplicationRequest
	Response Response `json:"response"`
}

type ConcentRequestedByBA struct {
	Metadata           Metadata    `json:"metadata"`
	ProductData        ProductData `json:"productData"`
	LoanApplicationIds []string    `json:"loanApplicationIds"`
	Consent            struct {
		ConsentFetchType     string `json:"consentFetchType"`
		Vua                  string `json:"vua"`
		Description          string `json:"description"`
		IsAggregationEnabled bool   `json:"isAggregationEnabled"`
		ConsentAggregationID string `json:"consentAggregationId"`
		ConsentHandle        string `json:"consentHandle"`
		ConsentStatus        string `json:"consentStatus"`
		URL                  string `json:"url"`
		ExtensibleData       struct {
		} `json:"extensibleData"`
	} `json:"consent"`
}

type ConcentProvidedByLender struct {
	ConcentRequestedByBA
	Response Response `json:"response"`
}

type ConcentStatusRequest struct {
	Metadata      Metadata    `json:"metadata"`
	ProductData   ProductData `json:"productData"`
	ConsentHandle string      `json:"consentHandle"`
}

type ConcentStatusResponse struct {
	ConcentStatusRequest
	Response Response `json:"response"`
}

type GenerateOfferRequest struct {
	Metadata           Metadata    `json:"metadata"`
	ProductData        ProductData `json:"productData"`
	LoanApplicationIds []string    `json:"loanApplicationIds"`
}

type Offer struct {
	ID        string `json:"id"`
	ValidTill string `json:"validTill"`
	Terms     struct {
		RequestedAmount      int    `json:"requestedAmount"`
		Currency             string `json:"currency"`
		SanctionedAmount     int    `json:"sanctionedAmount"`
		NetDisbursedAmount   int    `json:"netDisbursedAmount"`
		InterestType         string `json:"interestType"`
		InterestRate         int    `json:"interestRate"`
		AnnualPercentageRate int    `json:"annualPercentageRate"`
		CoolingOffPeriod     struct {
			Duration int    `json:"duration"`
			Unit     string `json:"unit"`
		} `json:"coolingOffPeriod"`
		TotalRepayableAmount int    `json:"totalRepayableAmount"`
		InterestAmount       int    `json:"interestAmount"`
		Description          string `json:"description"`
		Tenure               struct {
			Duration int    `json:"duration"`
			Unit     string `json:"unit"`
		} `json:"tenure"`
		LegalAgreement struct {
			Type string `json:"type"`
			Data string `json:"data"`
		} `json:"legalAgreement"`
		Documents []struct {
			Source           string `json:"source"`
			SourceIdentifier string `json:"sourceIdentifier"`
			Format           string `json:"format"`
			Reference        string `json:"reference"`
			Type             string `json:"type"`
			IsDataInline     bool   `json:"isDataInline"`
			Data             string `json:"data"`
		} `json:"documents"`
		Charges struct {
			Prepayment struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"prepayment"`
			Bounce struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"bounce"`
			LatePayment struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"latePayment"`
			Processing struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"processing"`
			Other struct {
				ChargeType string `json:"chargeType"`
				Data       struct {
					Rate                int    `json:"rate"`
					Amount              int    `json:"amount"`
					ApplicableParameter string `json:"applicableParameter"`
					Description         string `json:"description"`
					URL                 string `json:"url"`
				} `json:"data"`
				URL            string `json:"url"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"other"`
		} `json:"charges"`
		URL            string `json:"url"`
		ExtensibleData struct {
		} `json:"extensibleData"`
	} `json:"terms"`
	Disbursement struct {
		Plans []struct {
			ID                  string `json:"id"`
			Title               string `json:"title"`
			ShortDescription    string `json:"shortDescription"`
			Description         string `json:"description"`
			PaymentURL          string `json:"paymentUrl"`
			PayNowAllowed       bool   `json:"payNowAllowed"`
			EditPlanAllowed     bool   `json:"editPlanAllowed"`
			ChangeMethodAllowed bool   `json:"changeMethodAllowed"`
			Automatic           bool   `json:"automatic"`
			ScheduleType        string `json:"scheduleType"`
			NoOfInstallment     int    `json:"noOfInstallment"`
			Frequency           string `json:"frequency"`
			Tenure              struct {
				Duration int    `json:"duration"`
				Unit     string `json:"unit"`
			} `json:"tenure"`
			TotalAmount    int    `json:"totalAmount"`
			Principal      int    `json:"principal"`
			InterestAmount int    `json:"interestAmount"`
			Penalty        int    `json:"penalty"`
			StartDate      string `json:"startDate"`
			Status         string `json:"status"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"plans"`
		AccountDetails []struct {
			ID              string `json:"id"`
			Description     string `json:"description"`
			Status          string `json:"status"`
			AccountDataType string `json:"accountDataType"`
			Data            struct {
				AccountType         string `json:"accountType"`
				AccountIFSC         string `json:"accountIFSC"`
				AccountNumber       string `json:"accountNumber"`
				Vpa                 string `json:"vpa"`
				MaskedAccountNumber string `json:"maskedAccountNumber"`
				AccountHolderName   string `json:"accountHolderName"`
				URL                 string `json:"url"`
			} `json:"data"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"accountDetails"`
	} `json:"disbursement"`
	Repayment struct {
		Plans []struct {
			ID                  string `json:"id"`
			Title               string `json:"title"`
			ShortDescription    string `json:"shortDescription"`
			Description         string `json:"description"`
			PaymentURL          string `json:"paymentUrl"`
			PayNowAllowed       bool   `json:"payNowAllowed"`
			EditPlanAllowed     bool   `json:"editPlanAllowed"`
			ChangeMethodAllowed bool   `json:"changeMethodAllowed"`
			Automatic           bool   `json:"automatic"`
			ScheduleType        string `json:"scheduleType"`
			NoOfInstallment     int    `json:"noOfInstallment"`
			Frequency           string `json:"frequency"`
			Tenure              struct {
				Duration int    `json:"duration"`
				Unit     string `json:"unit"`
			} `json:"tenure"`
			TotalAmount    int    `json:"totalAmount"`
			Principal      int    `json:"principal"`
			InterestAmount int    `json:"interestAmount"`
			Penalty        int    `json:"penalty"`
			StartDate      string `json:"startDate"`
			Status         string `json:"status"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"plans"`
	} `json:"repayment"`
	Description string `json:"description"`
	Documents   []struct {
		Source           string `json:"source"`
		SourceIdentifier string `json:"sourceIdentifier"`
		Format           string `json:"format"`
		Reference        string `json:"reference"`
		Type             string `json:"type"`
		IsDataInline     bool   `json:"isDataInline"`
		Data             string `json:"data"`
	} `json:"documents"`
	URL            string `json:"url"`
	ExtensibleData struct {
	} `json:"extensibleData"`
}

type GenerateOfferMetadata struct {
	LoanApplications []struct {
		CreatedDate           string `json:"createdDate"`
		LoanApplicationID     string `json:"loanApplicationId"`
		LoanApplicationStatus string `json:"loanApplicationStatus"`
		Kyc                   struct {
			KycRefNo    string `json:"kycRefNo"`
			Description string `json:"description"`
			Business    struct {
				Type     string `json:"type"`
				Scale    string `json:"scale"`
				Category string `json:"category"`
				Name     string `json:"name"`
				Address  struct {
					Hba            string `json:"hba"`
					Srl            string `json:"srl"`
					Landmark       string `json:"landmark"`
					Als            string `json:"als"`
					Vtc            string `json:"vtc"`
					PinCode        string `json:"pinCode"`
					Po             string `json:"po"`
					District       string `json:"district"`
					State          string `json:"state"`
					Country        string `json:"country"`
					URL            string `json:"url"`
					Latitude       string `json:"latitude"`
					Longitude      string `json:"longitude"`
					ExtensibleData struct {
					} `json:"extensibleData"`
				} `json:"address"`
				Email             string `json:"email"`
				PhoneNumber       string `json:"phoneNumber"`
				IncorporationDate string `json:"incorporationDate"`
				CommencementDate  string `json:"commencementDate"`
				Udyam             struct {
					RegistrationNumber string `json:"registrationNumber"`
					RegistrationDate   string `json:"registrationDate"`
					ExtensibleData     struct {
					} `json:"extensibleData"`
				} `json:"udyam"`
				Status         string `json:"status"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"business"`
			Individual struct {
				Type    string `json:"type"`
				Name    string `json:"name"`
				Dob     string `json:"dob"`
				Address struct {
					Hba            string `json:"hba"`
					Srl            string `json:"srl"`
					Landmark       string `json:"landmark"`
					Als            string `json:"als"`
					Vtc            string `json:"vtc"`
					PinCode        string `json:"pinCode"`
					Po             string `json:"po"`
					District       string `json:"district"`
					State          string `json:"state"`
					Country        string `json:"country"`
					URL            string `json:"url"`
					Latitude       string `json:"latitude"`
					Longitude      string `json:"longitude"`
					ExtensibleData struct {
					} `json:"extensibleData"`
				} `json:"address"`
				Status         string `json:"status"`
				Email          string `json:"email"`
				PhoneNumber    string `json:"phoneNumber"`
				ExtensibleData struct {
				} `json:"extensibleData"`
			} `json:"individual"`
			Next struct {
				Type string `json:"type"`
				Data struct {
					URL string `json:"url"`
				} `json:"data"`
			} `json:"next"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"kyc"`
		Offers         []Offer `json:"offers"`
		Description    string  `json:"description"`
		URL            string  `json:"url"`
		ExtensibleData struct {
		} `json:"extensibleData"`
		RejectionDetails []struct {
			Reason         string `json:"reason"`
			Description    string `json:"description"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"rejectionDetails"`
		ActionRequired []struct {
			ActionType  string `json:"actionType"`
			Description string `json:"description"`
			Reference   struct {
				Object string `json:"object"`
				Value  string `json:"value"`
			} `json:"reference"`
			URL            string `json:"url"`
			ExtensibleData struct {
			} `json:"extensibleData"`
		} `json:"actionRequired"`
	} `json:"loanApplications"`
}

type GenerateOfferResponse struct {
	GenerateOfferRequest
	GenerateOfferMetadata
}

type AdditionalDocumentsRequest struct {
	Metadata         Metadata    `json:"metadata"`
	ProductData      ProductData `json:"productData"`
	LoanApplications []struct {
		LoanApplicationID   string `json:"loanApplicationId"`
		AdditionalDocuments []struct {
			Source           string `json:"source"`
			SourceIdentifier string `json:"sourceIdentifier"`
			Format           string `json:"format"`
			Reference        string `json:"reference"`
			Type             string `json:"type"`
			IsDataInline     bool   `json:"isDataInline"`
			Data             string `json:"data"`
		} `json:"additionalDocuments"`
	} `json:"loanApplications"`
}

type AdditionalDocumentResponse struct {
	Metadata    Metadata    `json:"metadata"`
	Response    Response    `json:"response"`
	ProductData ProductData `json:"productData"`
	GenerateOfferMetadata
}

type SetOfferRequest struct {
	Metadata          Metadata    `json:"metadata"`
	ProductData       ProductData `json:"productData"`
	LoanApplicationID string      `json:"loanApplicationId"`
	Offer             Offer       `json:"offer"`
}

type SetOfferResponse struct {
	Metadata              Metadata    `json:"metadata"`
	Response              Response    `json:"response"`
	ProductData           ProductData `json:"productData"`
	LoanApplicationID     string      `json:"loanApplicationId"`
	LoanApplicationStatus string      `json:"loanApplicationStatus"`
	ExtensibleData        struct {
	} `json:"extensibleData"`
}
