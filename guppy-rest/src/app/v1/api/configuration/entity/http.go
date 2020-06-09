package entity

// ConfigurationResponse ...
type ConfigurationResponse struct {
	Key  string `json:"key,omitempty"`
	User string `json:"user,omitempty"`
}

// ConfigurationUserRequest ...
type ConfigurationUserRequest struct {
	User     string `form:"user" json:"user"`
	Password string `form:"password" json:"password"`
}

type ConfigurationAppRequest struct {
}
