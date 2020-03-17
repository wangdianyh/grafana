package notifiers

import (
	"fmt"
	"strings"
	//"net/url"

	//"github.com/grafana/grafana/pkg/bus"
	firebase "firebase.google.com/go"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/alerting"
	//"firebase.google.com/go/messaging"
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

	return nil
}
