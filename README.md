# AWS SNS - Lambda Slack integration written in Go
Using (just recently) released official support for Go in lambda.
Inspired by [terranova's article](https://medium.com/cohealo-engineering/how-set-up-a-slack-channel-to-be-an-aws-sns-subscriber-63b4d57ad3ea).

## Install
- start with `dep ensure`

Follow steps in [aws manual](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/):
1. `GOOS=linux go build -o main`
2. `zip deployment.zip main`

- upload this file in lambda function setup
- make sure you specify ENV variable `WEBHOOK_URL` with appropriate slack webhook URL
- configure lambda trigger as SNS

## My usecase
AWS Healhcheck integration with slack, where Healthcheck sends notifications to the SNS 
which triggers the lambda.

