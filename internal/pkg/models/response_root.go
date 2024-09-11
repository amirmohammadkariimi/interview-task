package models

// model for root handler json response
type RootHandlerResponse struct {
	Date       int64  `json:"time"`
	Kubernetes bool   `json:"kubernetes"`
	Version    string `json:"version"`
}
