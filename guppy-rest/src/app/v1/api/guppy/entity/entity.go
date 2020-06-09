package entity

// InsertDataModels ...
type InsertDataModels struct {
	Path      string `json:"path"`
	Value     string `json:"value"`
	IsEncrypt bool   `json:"is_encrypt"`
}

// RequestPayload ...
type RequestPayload struct {
	Stage      string             `json:"stage"`
	Name       string             `json:"name"`
	Parameters []InsertDataModels `json:"parameters"`
}

// ParametersRequest ...
type ParametersRequest struct {
	Path      string `form:"path" json:"path"`
	IsEncrypt bool   `form:"decryption" json:"decryption"`
}

// GetResponse ...
type GetResponse struct {
	Path     string `json:"path"`
	Value    string `json:"value"`
	Revision int64  `json:"revision,omitempty"`
	Version  int64  `json:"version"`
}
