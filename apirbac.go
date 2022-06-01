package apirbac

import (
	"fmt"
	"regexp"
)

// Init function returns a pointer to RBAC.
// By returning a pointer, it ensures that you only have one instance of rbac configs
// in your application.
func Init() *RBAC {
	return &RBAC{}
}

// AddActions adds a new action.
// in a REST API, an action is like GET, POST, PUT etc or one of CRUD.
func (r *RBAC) AddActions(actions ...string) {
	r.Configs.Actions = append(r.Configs.Actions, actions...)
}

//AddResource adds a new resource.
// In a REST API, a resource can be an endpoint or a table etc.
// resourceID is the unique name of the resource
// e.g. an endpoint /users/4 can have a resourceID user
// and the regex can be /users/[0-9]+
func (r *RBAC) AddResource(resourceID, regex string) {
	resource := Resource{
		ID:    resourceID,
		Regex: regex,
	}
	r.Configs.Resources = append(r.Configs.Resources, resource)
}

// GetResource returns a resource or an error when the resource is not registered.
// it can be used in situations where you want to confirm if a resource is registered.
func (r *RBAC) GetResource(resourceID string) (Resource, error) {
	for _, resource := range r.Configs.Resources {
		if resource.ID == resourceID {
			return resource, nil
		}
	}
	return Resource{}, fmt.Errorf("resource not found")
}

// AddPermission adds a role and their permitted actions to a resource.
// if the role should have all permissions on a resource, use "*" as the action.
// Add all actions for a resource to a role in one line. If a resource is added twice, it will be rejected.
// If you want to add permissions for the same role but for different resources, do it in two calls.
func (r *RBAC) AddPermission(roleID, resourceID string, actions ...string) error {
	_, err := r.GetResource(resourceID)
	if err != nil {
		return err
	}

	role, roleIndex, err := r.GetRole(roleID)
	if err != nil {
		// here the role does not exist, create it
		r.Configs.Roles = append(r.Configs.Roles, Role{
			ID: roleID,
			Grants: []Grant{
				{
					ResourceID: resourceID,
					Actions:    actions,
				},
			},
		})
		return nil
	}
	// here role exists, check if the resource already exists
	for _, grant := range role.Grants {
		if grant.ResourceID == resourceID {
			return fmt.Errorf("role %s already has resource %s", roleID, resourceID)
		}
	}

	// add grants
	r.Configs.Roles[roleIndex].Grants = append(role.Grants, Grant{
		ResourceID: resourceID,
		Actions:    actions,
	})
	return nil
}

// GetRole returns a role or an error if it is not registered.
func (r *RBAC) GetRole(roleID string) (Role, int, error) {
	for i, _r := range r.Configs.Roles {
		if _r.ID == roleID {
			return _r, i, nil
		}
	}
	return Role{}, 0, fmt.Errorf("role %s not found", roleID)
}

// IsAllowed returns true if a role is allowed to perform an action on a resource
func (r *RBAC) IsAllowed(roleID, resourceValue string, action string) bool {
	role, _, err := r.GetRole(roleID)
	if err != nil {
		return false
	}

	grant, err := r.getResourceFromValue(resourceValue, role)
	if err != nil {
		return false
	}

	for _, permission := range grant.Actions {
		if permission == "*" {
			return true
		}

		if permission == action {
			return true
		}
	}
	return false
}

// getResourceFromValue returns the resource value when given a resourceID
// this is where the resource value is compared against stored resource regexes.
func (r *RBAC) getResourceFromValue(val string, role Role) (Grant, error) {
	for _, g := range role.Grants {
		resource, _ := r.GetResource(g.ResourceID)
		match, _ := regexp.MatchString(resource.Regex, val)
		if match {
			return g, nil
		}
	}
	return Grant{}, fmt.Errorf("resource not found")
}

// roleExists returns true if a roleID is registered and false otherwise
func (r *RBAC) roleExists(roleID string) bool {
	for _, role := range r.Configs.Roles {
		if role.ID == roleID {
			return true
		}
	}
	return false
}
