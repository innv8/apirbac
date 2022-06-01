/*

Package apirbac is a role based access control library.
It is meant to be used specifically in APIs

apirbac depends on these terms:

		- Role defines a group of people e.g admin, customer etc
		- Resource defines a shared item that can be accessed by or/more roles
		- Action defines the REST verbs e.g POST, GET etc


There are two ways of using this package:

1. Defining rules in code

2. Loading rules from a JSON config.

---


1. Defining rules in code:

		r := rbac.Init()

		// register all actions that will be available in your app
		r.AddActions("GET", "POST", "PUT")

		// register all the available resources
		r.AddResource("users", "/users")
		r.AddResource("user", `/users/[1-9]\d*`)
		r.AddResource("login", "/login")

		// add roles and their permitted actions on the resources
		err := r.AddPermission("admin", "users", "*") // admin has all permissions on users
		if err != nil {
			fmt.Printf("unable to add admin role because %v", err)
			return
		}

		// here we are defining a customer role that only has GET and POST
		// permissions on the users resource.
		err = r.AddPermission("customer", "users", "GET", "POST")
		if err != nil {
			fmt.Printf("unable to add customer role because %v", err)
			return
		}

		// to check whether a role is allowed an action to a resource
		allowed := r.IsAllowed("admin", "/users/3543", "GET") // returns true because the users/3543 matches the user regex and admin is allowed to do any action on the resource
		allowed = r.IsAllowed("customer", "/login", "GET") // returns false

		// to export the configurations
		err = r.SaveConfigs("./rbac-configs.json")

2. Loading rules from a JSON config.

		r := rbac.Init()
		err := r.LoadConfigs("./rbac-configs.json")

		// you can then go on to check if roles are allowed to access a resource
		allowed := r.IsAllowed("admin", "/users/3543", "GET") // returns true because the users/3543 matches the user regex and admin is allowed to do any action on the resource
		allowed = r.IsAllowed("customer", "/login", "GET") // returns false

		// to export the configurations
		err = r.SaveConfigs("./rbac-configs.json")

*/
package apirbac
