package notifiers

import (
	"testing"

	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFCMNotifier(t *testing.T) {
	Convey("FCM notifier tests", t, func() {
		Convey("empty settings should return error", func() {
			json := `{ }`

			settingsJSON, _ := simplejson.NewJson([]byte(json))
			model := &models.AlertNotification{
				Name:     "FCM_testing",
				Type:     "FCM",
				Settings: settingsJSON,
			}

			_, err := NewFCMNotifier(model)
			So(err, ShouldNotBeNil)

		})
		Convey("settings should trigger incident", func() {
			json := `
			{
  "token": "['abcdefgh0123456789']"
			}`
			settingsJSON, _ := simplejson.NewJson([]byte(json))
			model := &models.AlertNotification{
				Name:     "FCM_testing",
				Type:     "FCM",
				Settings: settingsJSON,
			}

			not, err := NewFCMNotifier(model)
			FCMNotifier := not.(*FCMNotifier)

			So(err, ShouldBeNil)
			So(FCMNotifier.Name, ShouldEqual, "line_testing")
			So(FCMNotifier.Type, ShouldEqual, "line")
			So(FCMNotifier.Token, ShouldEqual, "['abcdefgh0123456789']")
		})
	})
}
