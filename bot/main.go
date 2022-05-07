package main

import (
	"encoding/json"
	"log"
	"os"

	oumu "yukinissie.com/slack-reaction-bot/oumu"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var api = slack.New(os.Getenv("SLACK_BOT_TOKEN"))

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// JSONパース
	body := request.Body
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		log.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}

	// チャレンジレスポンス
	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		if err := json.Unmarshal([]byte(body), &r); err != nil {
			log.Println(err.Error())
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
		}
		return events.APIGatewayProxyResponse{Body: string([]byte(r.Challenge)), StatusCode: 200}, nil
	}

	// 任意の処理
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		switch ev := eventsAPIEvent.InnerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			oumu.Gaeshi(api, ev.Channel, ev.Text, ev.BotID)
			return events.APIGatewayProxyResponse{StatusCode: 200}, nil
		}
	}

	// 例外処理
	return events.APIGatewayProxyResponse{StatusCode: 400}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
