package apirbac

type Action string

type RBAC struct {
	Configs RBAConfigs
}

type RBAConfigs struct {
	// these describe json fields
	Actions   []string `json:"actions"`
	Resources []string `json:"resources"` // v0.1.3 - use regexes in place of resources
	Roles     []Role   `json:"roles"`
}

type Role struct {
	ID     string  `json:"id"`
	Grants []Grant `json:"grants"`
}

type Grant struct {
	Resource    string   `json:"resource"`
	Permissions []string `json:"permissions"`
}
