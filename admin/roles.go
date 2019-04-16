package admin

type Role struct {
	Name        string
	Permissions []string
}

var AdminRole = &Role{
	Name:        "Admin",
	Permissions: []string{"read", "create", "delete", "update"},
}

var GuestRole = &Role{
	Name:        "Guest",
	Permissions: []string{"read"},
}
