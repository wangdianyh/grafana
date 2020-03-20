package notifiers

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/alerting"
	"github.com/grafana/grafana/pkg/services/alerting/notifiers_sdk"
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
	fmt.Println("------------------------ FCM start to excute ----------------------")
	fcm.log.Info("Executing FCM notification", "ruleId", evalContext.Rule.ID, "notification", fcm.Name)
	//fmt.Printf("message: %v\v", evalContext.Rule.Message)
	var err error
	switch evalContext.Rule.State {
	case models.AlertStateAlerting:
		err = fcm.createAlert(evalContext)
	}
	return err
}

func (fcm *FCMNotifier) createAlert(evalContext *alerting.EvalContext) error {
	fmt.Println("------------------------ FCM start from here ----------------------")
	fcm.log.Info("Creating FCM notify", "ruleId", evalContext.Rule.ID, "notification", fcm.Name)

	//opt := option.WithCredentialsFile(file)
	opt := option.WithCredentialsJSON([]byte(notifiers_sdk.FB_sdk))

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
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: evalContext.Rule.Name,
			Body:  evalContext.Rule.Message,
		},
		Tokens: registrationTokens,
	}
	// [END android_message_golang]

	// Send a message to the device corresponding to the provided
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
