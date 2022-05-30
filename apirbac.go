package apirbac

import (
	"fmt"
	"regexp"
)

func Init() *RBAC {
	return &RBAC{}
}

func (r *RBAC) AddActions(actions ...string) {
	r.Configs.Actions = append(r.Configs.Actions, actions...)
}

func (r *RBAC) AddResources(resources ...string) {
	r.Configs.Resources = append(r.Configs.Resources, resources...)
}

func (r *RBAC) AddRole(roleID, resource string, permissions ...string) error {
	role, err := r.GetRole(roleID)
	if err != nil {
		// here the role does not exist, create it
		r.Configs.Roles = append(r.Configs.Roles, Role{
			ID: roleID,
			Grants: []Grant{
				{
					Resource:    resource,
					Permissions: permissions,
				},
			},
		})
		return nil
	}
	// here role exists, add grants
	role.Grants = append(role.Grants, Grant{
		Resource:    resource,
		Permissions: permissions,
	})
	return fmt.Errorf("role %s already exists", roleID)
}

func (r *RBAC) GetRole(roleID string) (role *Role, err error) {
	for _, _r := range r.Configs.Roles {
		if _r.ID == roleID {
			return &_r, nil
		}
	}
	return nil, fmt.Errorf("role %s not found", roleID)
}

func (r *RBAC) RoleExists(roleID string) bool {
	for _, role := range r.Configs.Roles {
		if role.ID == roleID {
			return true
		}
	}
	return false
}

func (r *RBAC) IsAllowed(roleID, resource, action string) bool {
	role, err := r.GetRole(roleID)
	if err != nil {
		return false
	}

	for _, grant := range role.Grants {

		if grant.Resource == resource {
			// search for the action
			for _, p := range grant.Permissions {
				// if a permission is *, return true
				if p == "*" {
					return true
				}

				if matched, _ := regexp.MatchString(p, action); matched {
					return true
				}
			}
		}
	}
	return false
}
