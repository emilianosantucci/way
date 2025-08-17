package rest

type NewApplication struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
