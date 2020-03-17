package notifiers

import (
	"fmt"
	//"net/url"

	//"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/alerting"
	"github.com/grafana/grafana/pkg/util"
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
	token := model.Settings.Get("token").MustString()
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

	//body := fmt.Sprintf("%s - %s\n%s", evalContext.Rule.Name, ruleURL, evalContext.Rule.Message)
	fmt.Printf("name: ", evalContext.Rule.Name)
	fmt.Printf("url: ", ruleURL)
	fmt.Printf("message: ", evalContext.Rule.Message)

	return nil
}
