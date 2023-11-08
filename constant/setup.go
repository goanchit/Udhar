package constant

type Configure struct {
	PublisherType     string
	RecieverType      string
	Name              string
	SNSDelayInSeconds string
}

var SETUP = [...]Configure{
	{
		PublisherType:     "SNS",
		RecieverType:      "SQS",
		Name:              "CREATE-LOAN-APPLICATION",
		SNSDelayInSeconds: "5",
	},
	{
		PublisherType:     "SNS",
		RecieverType:      "SQS",
		Name:              "CREATE-CONCENT-ON-AA",
		SNSDelayInSeconds: "10",
	},
	{
		PublisherType:     "SNS",
		RecieverType:      "SQS",
		Name:              "NOTIFY-BA",
		SNSDelayInSeconds: "5",
	},
	{
		PublisherType:     "SNS",
		RecieverType:      "SQS",
		Name:              "GENERATE-OFFERS",
		SNSDelayInSeconds: "10",
	},
	{
		PublisherType:     "SNS",
		RecieverType:      "SQS",
		Name:              "SUBMIT-ADDITIONAL-DOCUMENTS",
		SNSDelayInSeconds: "10",
	},
	{
		PublisherType:     "SNS",
		RecieverType:      "SQS",
		Name:              "SET-OFFER",
		SNSDelayInSeconds: "5",
	},
}
