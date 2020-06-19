package fcm

// easyjson:json
type sendRequest struct {
	Message *Message `json:"message"`
}
