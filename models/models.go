package models

import "github.com/google/uuid"

type RequestBodyWithTraceId struct {
	TraceId uuid.UUID `json:"traceId"`
}
