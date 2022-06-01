package apirbac

// RBAC is the struct that makes all this possible.
// It holds all methods and data for your rbac rules.
type RBAC struct {
	Configs RBACConfigs
}

// RBACConfigs is the json format for configs.
// Actions are the possible actions in your program. e.g. GET, POST, Create etc
type RBACConfigs struct {
	Actions   []string   `json:"actions"`
	Resources []Resource `json:"resources"`
	Roles     []Role     `json:"roles"`
}

// Resource is an item in your program that you want to control access to.
// It could be an endpoints, table name or any other resource.
// The ID is the name you want to give the resource. For example, if it is an endpoint to
// show one user, the name could be 'user'
// The regex is the pattern for all endpoints that should be identified as the resource.
// In the single user resource, the endpoint could be /users/34 or /users/3.
// Therefore, you can use the regex `/users/[1-9]\d*` to describe all possible endpoints that match the resource.
type Resource struct {
	ID    string `json:"id"`
	Regex string `json:"regex"`
}

// Role is basically the roles in your program.
// All users can be part of a role.
// Example roles are admin, customer etc.
type Role struct {
	ID     string  `json:"id"`
	Grants []Grant `json:"grants"`
}

// Grant is an array that describes the permissions on a resource that a role has.
// It has the Resource and an array of permissions (which are actions).
// If a role has all permissions on a Resource, you can put just '*' in the Actions.
type Grant struct {
	Resource Resource `json:"resource"`
	Actions  []string `json:"actions"`
}
