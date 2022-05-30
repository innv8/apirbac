package apirbac

import "regexp"

type Action string

type RBAC struct {
	Configs RBAConfigs
}

type RBAConfigs struct {
	// these describe json fields
	Actions   []string   `json:"actions"`
	Resources []Resource `json:"resources"` // v0.1.3 - use regexes in place of resources
	Roles     []Role     `json:"roles"`
}

type Resource struct {
	ID    string         `json:"id"`
	Regex string         `json:"regex"`
	Rgx   *regexp.Regexp `json:"rgx,omitempty"`
}

type Role struct {
	ID     string  `json:"id"`
	Grants []Grant `json:"grants"`
}

type Grant struct {
	Resource    Resource `json:"resource"`
	Permissions []string `json:"permissions"`
}
