package apirbac

import (
	"fmt"
	"log"
	"regexp"
)

func Init() *RBAC {
	return &RBAC{}
}

func (r *RBAC) AddActions(actions ...string) {
	r.Configs.Actions = append(r.Configs.Actions, actions...)
}

func (r *RBAC) AddResource(resourceID, regex string) {
	resource := Resource{
		ID:    resourceID,
		Regex: regex,
		Rgx:   regexp.MustCompile(regex),
	}
	r.Configs.Resources = append(r.Configs.Resources, resource)
}

func (r *RBAC) GetResource(resourceID string) (Resource, error) {
	for _, resource := range r.Configs.Resources {
		if resource.ID == resourceID {
			return resource, nil
		}
	}
	return Resource{}, fmt.Errorf("resource not found")
}
func (r *RBAC) AddPermission(roleID, resourceID string, permissions ...string) error {
	resource, err := r.GetResource(resourceID)
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
					Resource:    resource,
					Permissions: permissions,
				},
			},
		})
		return nil
	}
	// here role exists, check if the resource already exists
	for _, grant := range role.Grants {
		if grant.Resource.ID == resourceID {
			return fmt.Errorf("role %s already has resource %s", roleID, resourceID)
		}
	}

	// add grants
	r.Configs.Roles[roleIndex].Grants = append(role.Grants, Grant{
		Resource:    resource,
		Permissions: permissions,
	})
	return nil
}

func (r *RBAC) GetRole(roleID string) (Role, int, error) {
	for i, _r := range r.Configs.Roles {
		if _r.ID == roleID {
			return _r, i, nil
		}
	}
	return Role{}, 0, fmt.Errorf("role %s not found", roleID)
}

func (r *RBAC) RoleExists(roleID string) bool {
	for _, role := range r.Configs.Roles {
		if role.ID == roleID {
			return true
		}
	}
	return false
}

func (r *RBAC) IsAllowed(roleID, resourceValue, action string) bool {
	role, _, err := r.GetRole(roleID)
	if err != nil {
		return false
	}

	grant, err := getResourceFromValue(resourceValue, role)
	if err != nil {
		return false
	}

	for _, permission := range grant.Permissions {
		if permission == "*" {
			return true
		}

		if permission == action {
			return true
		}
	}
	return false
}

func getResourceFromValue(val string, role Role) (Grant, error) {
	for _, g := range role.Grants {
		log.Println("---------------------", g.Resource.Regex)
		log.Println("---------------------", g.Resource.Rgx == nil)

		match := g.Resource.Rgx.MatchString(val)
		if match {
			return g, nil
		}
	}
	return Grant{}, fmt.Errorf("resource not found")
}
