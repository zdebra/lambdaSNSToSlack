package main

import (
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidRecords = errors.New("invalid records")
	ErrInvalidMessage = errors.New("ivnalid records")
	ErrInvalidSubject = errors.New("invalid subject")
)

var sampleMessage = `
{
  "Records": [
    {
      "EventSource": "aws:sns",
      "EventVersion": "1.0",
      "EventSubscriptionArn": "arn:aws:sns:us-east-1...ElasticBeanstalkNotifications-Environment-foo-app...",
      "Sns": {
        "Type": "Notification",
        "MessageId": "111",
        "TopicArn": "arn:aws:sns:us-east-1...ElasticBeanstalkNotifications-Environment-foo-app",
        "Subject": "AWS Elastic Beanstalk Notification - New application version was deployed to running EC2 instances",
        "Message": "Timestamp: Thu May 07 23:38:22 UTC 2015\nMessage: Lambda Function Test: New application version was deployed to running EC2 instances.\n\nEnvironment: foo-app\nApplication: FooApp\n\nEnvironment URL: http://foo.com\nRequestId: 222\nNotificationProcessId: 333",
        "Timestamp": "2015-05-07T23:39:18.628Z",
        "SignatureVersion": "1",
        "Signature": "hello-sig",
        "SigningCertUrl": "https://sns.us-east-1.amazonaws.com/...",
        "UnsubscribeUrl": "https://sns.us-east-1.amazonaws.com/...",
        "MessageAttributes": {}
      }
    }
  ]
}`

// SNSMessage for our case is containing just a single record
type SNSMessage struct {
	Records []Record `json:"Records"`
}

type Record struct {
	EventSource          string `json:"EventSource"`
	EventVersion         string `json:"EventVersion"`
	EventSubscriptionArn string `json:"EventSubscriptionArn"`
	SNS                  SNS    `json:"Sns"`
}

type SNS struct {
	Type              string                 `json:"Type"`
	MessageID         string                 `json:"MessageId"`
	TopicARN          string                 `json:"TopicArn"`
	Subject           string                 `json:"Subject"`
	Message           string                 `json:"Message"`
	Timestamp         time.Time              `json:"Timestamp"`
	SignatureVersion  string                 `json:"SignatureVersion"`
	Signature         string                 `json:"Signature"`
	SigningCertURL    string                 `json:"SigningCertUrl"`
	UnsubscribeURL    string                 `json:"UnsubscribeUrl"`
	MessageAttributes map[string]interface{} `json:"MessageAttributes"`
}

// Validate returns error if something is wrong with SNSMessage struct
func (m SNSMessage) Validate() error {
	switch {
	case len(m.Records) != 1:
		return ErrInvalidRecords
	case m.Records[0].SNS.Message == "":
		return ErrInvalidMessage
	case m.Records[0].SNS.Subject == "":
		return ErrInvalidSubject
	default:
		return nil
	}
}
