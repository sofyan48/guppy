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

// MetaListParams ...
type MetaListParams struct {
	Limit string `form:"limit" json:"limit"`
	Page  string `form:"page" json:"page"`
}

// MetaUserRequest ...
type MetaUserRequest struct {
	User string `form:"user" json:"user"`
}

// UserDataConfig ...
type UserDataConfig struct {
	User     string `json:"user"`
	Roles    string `json:"role"`
	Path     string `json:"key"`
	Password string `json:"password,omitempty"`
}
