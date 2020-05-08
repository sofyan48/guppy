package entity

// InsertDataModels ...
type InsertDataModels struct {
	Path      string `json:"path"`
	Value     string `json:"value"`
	IsEncrypt bool   `json:"isEncrypt"`
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
	IsEncrypt bool   `form:"is_encrypt" json:"is_encrypt"`
}

type GetResponse struct {
	Path           string `json:"path"`
	Value          string `json:"value"`
	CreateRevision int64  `json:"create_revision"`
	UpdateRevision int64  `json:"update_revision"`
}
