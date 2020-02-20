package fcm

// Response represents the FCM server's response to the application
// server's sent message.
type Response struct {
	MulticastID  int64    `json:"multicast_id"`
	Success      int      `json:"success"`
	Failure      int      `json:"failure"`
	CanonicalIDs int      `json:"canonical_ids"`
	StatusCode   int      `json:"error_code"`
	Results      []Result `json:"results"`
}

// Result represents the status of a processed message.
type Result struct {
	MessageID      string `json:"message_id"`
	RegistrationID string `json:"registration_id"`
	Error          string `json:"error"`
}

// Unregistered checks if the device token is unregistered,
// according to response from FCM server. Useful to determine
// if app is uninstalled.
func (r Result) Unregistered() bool {
	switch r.Error {
	case "NotRegistered", "MismatchSenderID", "MissingRegistration", "InvalidRegistration":
		return true

	default:
		return false
	}
}
