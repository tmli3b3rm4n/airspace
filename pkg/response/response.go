package response

type Message struct {
	Endpoint string `json:"endpoint"`
	Value    bool   `json:"value"`
}

type Response struct {
	Status  string  `json:"status"`
	Message Message `json:"message"`
	Error   string  `json:"error,omitempty"`
}
