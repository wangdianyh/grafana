package notifiers

import (
	"context"
	"fmt"
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
		Type:        "FCMS",
		Name:        "FCMS",
		Description: "Send notifications to FCMS notify",
		Factory:     NewFCMSNotifier,
		OptionsTemplate: `
		<div class="gf-form-group">
      <h3 class="page-heading">FCMS notify settings</h3>
      <div class="gf-form">
        <span class="gf-form-label width-14">Token</span>
        <input type="text" required class="gf-form-input max-width-22" ng-model="ctrl.model.settings.token" placeholder="FCM notify token key"></input>
      </div>
    </div>
`,
	})
}

// FCMNotifier is responsible for sending
// alert notifications to FCM.
type FCMSNotifier struct {
	NotifierBase
	Token string
	log   log.Logger
}

// NewFCMNotifier is the constructor for the FCM notifier
func NewFCMSNotifier(model *models.AlertNotification) (alerting.Notifier, error) {
	token := model.Settings.Get("token").MustString()
	if token == "" {
		return nil, alerting.ValidationError{Reason: "Could not find token in settings"}
	}
	// split toke
	//var line = strings.TrimSuffix(token, "\n")
	//fmt.Printf("'%s'\n", line)
	//tokens := strings.Split(line, ";")
	//fmt.Printf("%q\n", tokens)
	return &FCMSNotifier{
		NotifierBase: NewNotifierBase(model),
		Token:        token,
		log:          log.New("alerting.notifier.FCMS"),
	}, nil
}

// Notify send an alert notification to FCM
func (fcms *FCMSNotifier) Notify(evalContext *alerting.EvalContext) error {
	fcms.log.Info("Executing FCM notification", "ruleId", evalContext.Rule.ID, "notification", fcms.Name)

	var err error
	/*
		switch evalContext.Rule.State {
		case models.AlertStateAlerting:
			err = fcm.createAlert(evalContext)
		}
	*/
	if evalContext.Rule.State == models.AlertStateAlerting {
		err = fcms.createAlert(evalContext)
	}
	return err
}

func (fcms *FCMSNotifier) createAlert(evalContext *alerting.EvalContext) error {
	fcms.log.Info("Creating FCM notify", "ruleId", evalContext.Rule.ID, "notification", fcms.Name)
	ruleURL, err := evalContext.GetRuleURL()
	if err != nil {
		fcms.log.Error("Failed get rule link", "error", err)
		return err
	}
	fmt.Println("token: ", fcms.Token)
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
	registrationTokens := fcms.Token

	// See documentation on defining a message payload.
	// [START android_message_golang]
	//oneHour := time.Duration(1) * time.Hour
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "$GOOG up 1.43% on the day",
			Body:  "$GOOG gained 11.80 points to close at 835.67, up 1.43% on the day.",
		},
		Token: registrationTokens,
	}
	// [END android_message_golang]
	br, err := client.Send(ctx, message)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Response is a message ID string.
	fmt.Println("Successfully sent message:", br)

	return nil
}
