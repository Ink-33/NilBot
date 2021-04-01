package nilbot

// Message is used to unmarshal pushed event.
type Message []struct {
	Type string `json:"type"`
	Data struct {
		Content string `json:"content"`
		Cover   string `json:"cover"`
		Date    string `json:"data"`
		File    string `json:"file"`
		ID      string `json:"id"`
		Image   string `json:"image"`
		Magic   string `json:"magic"`
		QQ      string `json:"qq"`
		ResID   string `json:"resid"`
		Text    string `json:"text"`
		Title   string `json:"title"`
		Type    string `json:"type"`
		URL     string `json:"url"`
	} `json:"data"`
}
