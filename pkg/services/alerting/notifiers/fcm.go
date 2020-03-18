package notifiers

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/alerting"
	"github.com/grafana/grafana/pkg/util"
	"google.golang.org/api/option"
)

func init() {
	alerting.RegisterNotifier(&alerting.NotifierPlugin{
		Type:        "FCM",
		Name:        "FCM",
		Description: "Send notifications to FCM notify",
		Factory:     NewFCMNotifier,
		OptionsTemplate: `
		<h3 class="page-heading">FCM settings</h3>
		<div class="gf-form">
			<label class="gf-form-label width-8">
				Tokens
			</label>
			<textarea rows="7" class="gf-form-input width-27" required ng-model="ctrl.model.settings.tokens"></textarea>
		</div>
		<div class="gf-form offset-width-8">
			<span>You can enter multiple tokens using a ";" separator</span>
		</div>
`,
	})
}

// FCMNotifier is responsible for sending
// alert notifications to FCM.
type FCMNotifier struct {
	NotifierBase
	Token []string
	log   log.Logger
}

// NewFCMNotifier is the constructor for the FCM notifier
func NewFCMNotifier(model *models.AlertNotification) (alerting.Notifier, error) {
	token := model.Settings.Get("tokens").MustString()
	if token == "" {
		return nil, alerting.ValidationError{Reason: "Could not find token in settings"}
	}
	// split toke
	tokens := util.SplitEmails(token)
	return &FCMNotifier{
		NotifierBase: NewNotifierBase(model),
		Token:        tokens,
		log:          log.New("alerting.notifier.FCM"),
	}, nil
}

// Notify send an alert notification to FCM
func (fcm *FCMNotifier) Notify(evalContext *alerting.EvalContext) error {
	fcm.log.Info("Executing FCM notification", "ruleId", evalContext.Rule.ID, "notification", fcm.Name)

	var err error

	if evalContext.Rule.State == models.AlertStateAlerting {
		err = fcm.createAlert(evalContext)
	}
	return err
}

func (fcm *FCMNotifier) createAlert(evalContext *alerting.EvalContext) error {
	fcm.log.Info("Creating FCM notify", "ruleId", evalContext.Rule.ID, "notification", fcm.Name)

	fmt.Println("token: ", fcm.Token)

	//opt := option.WithCredentialsFile(file)
	opt := option.WithCredentialsJSON([]byte(`{
		"type": "service_account",
		"project_id": "grafana-notification",
		"private_key_id": "92c0d03925ac945e5f80e5b96e77f823cc6c4fee",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCw0UgfU2UKAZva\n65o2DFiLwjgq/sRxdNbpKxx4WpwtvfmN7p8imqmJ3kF4kJhklUVxt4ChQWdaRup9\nnR0iSI+9awez6RbUx48/+n+L86kCGGSKNDy5LW9CD75N8xt6dWZvduhIcP0KQFVH\nH8lENRbAP1kkZNHjNvNFj4JgZGnULe5ZrgzlR22hzzATITuCsVV6IR7MlN8IBqSz\n8HD0hUH1lUha9co0n7tVw5JiWMbWG4qDKgfltpnAlXwcU+y4FlY+G6N5Oa8A7Mr2\n9AzBNcsWhohqCmNRNqGw+r5KjHza+TsZckXqZOqe/cpQAsUPI9IDIFh5idi6yKhU\nIFZHrp3dAgMBAAECggEAAfrStOiviLu/emVvvGUwMWucsPzzR1Xu+UnIFgZ/TMY3\nK7V24K6DCw3xa7kqvtpd0bBW5124TrsQJK63eDFmmQQmxocfAB9P5YAXh4I6uGqv\ngLPYoWldBb9WLfQOSoIgAhixvYXn8pwJZTRGggkzvQLrxTJZZN+yCThaVfWQRC9O\ntN7JKAuMI0+tsSUlnu68kXmpMsbn3AwJLy8jPRLWp9g2VZTdb31/+A1VoqeNNZ1p\nJ/V77cdUnp2n6Y3xOPEtg3FLgvD8WI84buuR8PqKVuaF+Yxur5v7BZ91op6sECkC\n4ijRMxPPKhEmNsVtaagjcEFdD7LrhSWSiw1IsNtfQQKBgQDZ1EZYRnR9ZOmR/yZN\nSC4Jz88P+I3GOxf0TLGrRJq+Y+k5ZKBhLGpM6aVlVzku6m6bwA0mTulPDyqDgg6g\nQrIVRWU3zTiK6S5s7JtSW6Slh31fLwkZd7EX9Oy5B8TFdZoJzEN026sbisS622Bk\ns0Gq6lXcpLRAUksDcHVqrhjScQKBgQDPzT1pz+iS/+ivuPtKlNCjtzfv1eAUoW4V\nxxJjXzsOvR4Tf+2C3S5LNA2tx6geUQy+ADh7I6eLdjktfM6fHG4SkJtvJd2GoehN\nzcgiqBpH34hm9IfwmDCupMqrJ72AH7YvbYXWvvipGHCAli5ZJJHhL4lYFd66pNAV\n2DcdpfugLQKBgBXiNY74xQsz8CMytu5cqgNiVTMNjXC0zxtD+TVzlvg5oVyat2IL\nzEId1vfvY1dLRgFvseJ/WwEOTP8ZOc7v5GQurJSGkX+jHX7j5lbHziqzCe1eFFPy\nql/1wzJzjVkpD2iclMpQp0gFEO6Uy4JSX+6DzEx2X4V2vwKBccpd4zCBAoGAHsjs\nGTno3aY15ZqE9+aWBjsFeW149fV4ZpeIXNpl2GgiBYeFO0bjLdb3U9BpUpx1Q8yq\nkWuVza5lCB0eSyoeEHgF3vCAIgrobGZZCPFYe19dSMtfPEB/rc/SCosnosyP4/TY\nyBigpARv3kzhbulhBzhQo5ER3xq9jQ7sE2NcpL0CgYAcuNPH3LTsVm32IUvtGT9s\npZ7zF80/75owNTQPa9rKPVdOlJO/ffecBta0s0T027LgmBEEqOOs6e3JqVg6lk/S\nExK7F/RrF3o8/z82LYPNP0KeQEBsqYOIaTkIcCkC2Jw7lV8O+rbmHHj/SAYLV7Yt\nEBWSL36Dte+p7RTwkqpoIw==\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-mqs3o@grafana-notification.iam.gserviceaccount.com",
		"client_id": "113789093807351232244",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-mqs3o%40grafana-notification.iam.gserviceaccount.com"
	  }`))

	var err error
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v\n", err)
		return err
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		fmt.Printf("error getting Messaging client: %v\n", err)
		return err
	}

	// This registration token comes from the client FCM SDKs.
	registrationTokens := fcm.Token

	// See documentation on defining a message payload.
	// [START android_message_golang]
	//oneHour := time.Duration(1) * time.Hour
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: evalContext.Rule.Name,
			Body:  evalContext.Rule.Message,
		},
		Tokens: registrationTokens,
	}
	// [END android_message_golang]

	// Send a message to the device corresponding to the provided
	// registration token.
	br, err := client.SendMulticast(context.Background(), message)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	if br.FailureCount > 0 {
		var failedTokens []string
		for idx, resp := range br.Responses {
			if !resp.Success {
				// The order of responses corresponds to the order of the registration tokens.
				failedTokens = append(failedTokens, registrationTokens[idx])
			}
		}

		fmt.Printf("List of tokens that caused failures: %v\n", failedTokens)
	}

	// Response is a message ID string.
	fmt.Println("Successfully sent message:", br.SuccessCount)

	return nil
}
