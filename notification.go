package fcm

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrInvalidMessage occurs if push notitication message is nil.
	ErrInvalidMessage = errors.New("message is invalid")

	// ErrInvalidTarget occurs if message topic is empty.
	ErrInvalidTarget = errors.New("topic is invalid or registration ids are not set")
)

// Message represents list of targets, options, and payload for HTTP JSON
// messages.
// See https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#resource:-message
type Message struct {
	Name         string            `json:"name,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
	Notification *Notification     `json:"notification,omitempty"`
	Android      *AndroidConfig    `json:"android,omitempty"`

	// one of
	Token     string `json:"token,omitempty"`
	Topic     string `json:"topic,omitempty"`
	Condition string `json:"condition,omitempty"`
}

// See https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#Notification
type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Image string `json:"image,omitempty"`
}

// See https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#AndroidConfig
type AndroidConfig struct {
	CollapseKey           string                 `json:"collapse_key,omitempty"`
	Priority              AndroidMessagePriority `json:"priority,omitempty"`
	Ttl                   string                 `json:"ttl,omitempty"`
	RestrictedPackageName string                 `json:"restricted_package_name,omitempty"`
	Data                  map[string]string      `json:"data,omitempty"`
	Notification          *AndroidNotification   `json:"notification,omitempty"`
}

// Notification specifies the predefined, user-visible key-value pairs of the
// notification payload.
// See https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#AndroidNotification
type AndroidNotification struct {
	Title                 string               `json:"title,omitempty"`
	Body                  string               `json:"body,omitempty"`
	Icon                  string               `json:"icon,omitempty"`
	Color                 string               `json:"color,omitempty"`
	Sound                 string               `json:"sound,omitempty"`
	Tag                   string               `json:"tag,omitempty"`
	ClickAction           string               `json:"click_action,omitempty"`
	BodyLocKey            string               `json:"body_loc_key,omitempty"`
	BodyLocArgs           []string             `json:"body_loc_args,omitempty"`
	TitleLocKey           string               `json:"title_loc_key,omitempty"`
	TitleLocArgs          []string             `json:"title_loc_args,omitempty"`
	ChannelId             string               `json:"channel_id,omitempty"`
	Ticker                string               `json:"ticker,omitempty"`
	Sticky                bool                 `json:"sticky,omitempty"`
	EventName             string               `json:"event_name,omitempty"`
	LocalOnly             bool                 `json:"local_only,omitempty"`
	NotificationPriority  NotificationPriority `json:"notification_priority,omitempty"`
	DefaultSound          bool                 `json:"default_sound,omitempty"`
	DefaultVibrateTimings bool                 `json:"default_vibrate_timings,omitempty"`
	DefaultLightSettings  bool                 `json:"default_light_settings,omitempty"`
	VibrateTimings        []string             `json:"vibrate_timings,omitempty"`
	Visibility            Visibility           `json:"visibility,omitempty"`
	NotificationCount     int                  `json:"notification_count,omitempty"`
	LightSettings         *LightSettings       `json:"light_settings,omitempty"`
	Image                 string               `json:"image,omitempty"`
}

type Color struct {
	Red   float32 `json:"red,omitempty"`
	Green float32 `json:"green,omitempty"`
	Blue  float32 `json:"blue,omitempty"`
	Alpha float32 `json:"alpha,omitempty"`
}

type LightSettings struct {
	Color            Color
	LightOnDuration  string `json:"light_on_duration,omitempty"`
	LightOffDuration string `json:"light_off_duration,omitempty"`
}

type AndroidMessagePriority string

type Visibility string

type NotificationPriority string

const (
	NotificationPriorityUnspecified NotificationPriority = "PRIORITY_UNSPECIFIED"
	NotificationPriorityMin         NotificationPriority = "PRIORITY_MIN"
	NotificationPriorityLow         NotificationPriority = "PRIORITY_LOW"
	NotificationPriorityDefault     NotificationPriority = "PRIORITY_DEFAULT"
	NotificationPriorityHigh        NotificationPriority = "PRIORITY_HIGH"
	NotificationPriorityMax         NotificationPriority = "PRIORITY_MAX"

	VisibilityPrivate Visibility = "PRIVATE"
	VisibilityPublic  Visibility = "PUBLIC"
	VisibilitySecret  Visibility = "SECRET"

	AndroidMessagePriorityNormal AndroidMessagePriority = "NORMAL"
	AndroidMessagePriorityHigh   AndroidMessagePriority = "HIGH"
)

// Validate returns an error if the message is not well-formed.
func (msg *Message) Validate() error {
	if msg == nil {
		return ErrInvalidMessage
	}

	// validate target identifier: `to` or `condition`, or `registration_ids`
	opCnt := strings.Count(msg.Condition, "&&") + strings.Count(msg.Condition, "||")
	if msg.Token == "" && (msg.Condition == "" || opCnt > 2) && len(msg.Topic) == 0 {
		return ErrInvalidTarget
	}
	return nil
}
