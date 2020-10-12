package helpers

type (
	ResponseError struct {
		Message    string `json:"message"`
		Resolution string `json:"resolution,omitempty"`
		Error      error  `json:"error,omitempty"`
	}

	Json map[string]interface{}
)
