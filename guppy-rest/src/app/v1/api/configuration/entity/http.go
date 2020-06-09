package entity

// ConfigurationResponse ...
type ConfigurationResponse struct {
	Key  string `json:"authorization,omitempty"`
	User string `json:"user,omitempty"`
}

// AddUserConfigRequest ...
type AddUserConfigRequest struct {
	User     string `form:"user" json:"user"`
	Password string `form:"password" json:"password"`
	Roles    string `form:"role" json:"role"`
}

// ConfigurationUserRequest ...
type ConfigurationUserRequest struct {
	User     string `form:"user" json:"user"`
	Password string `form:"password" json:"password"`
}

type ConfigurationAppRequest struct {
}
