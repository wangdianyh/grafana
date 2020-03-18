package notifiers

import (
	"context"
	"fmt"
	"strings"
	//"net/url"

	//"github.com/grafana/grafana/pkg/bus"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/alerting"
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
	var line = strings.TrimSuffix(token, "\n")
	fmt.Printf("'%s'\n", line)
	tokens := strings.Split(line, ";")
	fmt.Printf("%q\n", tokens)
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
	/*
		switch evalContext.Rule.State {
		case models.AlertStateAlerting:
			err = fcm.createAlert(evalContext)
		}
	*/
	if evalContext.Rule.State == models.AlertStateAlerting {
		err = fcm.createAlert(evalContext)
	}
	return err
}

func (fcm *FCMNotifier) createAlert(evalContext *alerting.EvalContext) error {
	fcm.log.Info("Creating FCM notify", "ruleId", evalContext.Rule.ID, "notification", fcm.Name)
	ruleURL, err := evalContext.GetRuleURL()
	if err != nil {
		fcm.log.Error("Failed get rule link", "error", err)
		return err
	}
	fmt.Println("token: ", fcm.Token)
	//body := fmt.Sprintf("%s - %s\n%s", evalContext.Rule.Name, ruleURL, evalContext.Rule.Message)
	fmt.Println("name: ", evalContext.Rule.Name)
	fmt.Println("url: ", ruleURL)
	fmt.Println("message: ", evalContext.Rule.Message)

	opt := option.WithCredentialsFile("grafana-notification-firebase-adminsdk-mqs3o-92c0d03925.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v\n", err)
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		fmt.Printf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationTokens := []string{
		"e2bPi8R8CC4:APA91bG_7VfqGSCRalDoa8sR9Ggsng5m2FKFsW0CkocYkKpEbHQHA2Q8kSKx5ddP0ZDSN7NPRf_Cyqga1ObQveqpdxmDsYhLDe1n4vWEtpYQwgnuJAn3r2BGzPTMpUGokF-kDxRqO6vM",
	}

	// See documentation on defining a message payload.
	// [START android_message_golang]
	//oneHour := time.Duration(1) * time.Hour
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "$GOOG up 1.43% on the day",
			Body:  "$GOOG gained 11.80 points to close at 835.67, up 1.43% on the day.",
		},
		Tokens: registrationTokens,
	}
	// [END android_message_golang]

	// Send a message to the device corresponding to the provided
	// registration token.
	br, err := client.SendMulticast(context.Background(), message)
	if err != nil {
		fmt.Println("err: ", err)
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
