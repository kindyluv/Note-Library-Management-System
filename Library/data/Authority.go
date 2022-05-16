package data

import (
	"errors"
	"gorm.io/gorm"
)

type Authority struct {
	DB *gorm.DB
}

type Options struct {
	TablePrefix string
	DB          *gorm.DB
}

var (
	ErrPermissionInUse     = errors.New("cannot delete assigned permission")
	ErrPermissionNotFound  = errors.New("permission not found")
	ErrRoleAlreadyAssigned = errors.New("this role is already assigned to the user")
	ErrRoleInUse           = errors.New("cannot delete assigned role")
	ErrRoleNotFound        = errors.New("role not found")
)

var tablePrefix string

var auth *Authority

func New(opts Options) *Authority {
	tablePrefix = opts.TablePrefix
	auth = &Authority{
		DB: opts.DB,
	}
	migrateTables(opts.DB)
	return auth
}

func Resolve() *Authority {
	return auth
}

func (a *Authority) CreateRole(roleName string) error {
	var dbRole Role
	res := a.DB.Where("name = ?", roleName).First(&dbRole)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			a.DB.Create(&Role{Name: roleName})
			return nil
		}
	}

	return res.Error
}

func (a *Authority) CreatePermission(permName string) error {
	var dbPerm UserPermission
	res := a.DB.Where("name = ?", permName).First(&dbPerm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// create
			a.DB.Create(&UserPermission{Name: permName})
			return nil
		}
	}

	return res.Error
}

func (a *Authority) AssignPermissions(roleName string, permNames []string) error {
	// get the role id
	var role Role
	rRes := a.DB.Where("name = ?", roleName).First(&role)
	if rRes.Error != nil {
		if errors.Is(rRes.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	var perms []UserPermission
	// get the permissions ids
	for _, permName := range permNames {
		var perm UserPermission
		pRes := a.DB.Where("name = ?", permName).First(&perm)
		if pRes.Error != nil {
			if errors.Is(pRes.Error, gorm.ErrRecordNotFound) {
				return ErrPermissionNotFound
			}

		}

		perms = append(perms, perm)
	}

	// insert data into RolePermissions table
	for _, perm := range perms {
		// ignore any assigned permission
		var rolePerm RolePermission
		res := a.DB.Where("role_id = ?", role.ID).Where("permission_id =?", perm.ID).First(&rolePerm)
		if res.Error != nil {
			// assign the record
			cRes := a.DB.Create(&RolePermission{RoleID: role.ID, PermissionID: perm.ID})
			if cRes.Error != nil {
				return cRes.Error
			}
		}
	}

	return nil
}

func (a *Authority) AssignRole(userID uint, roleName string) error {
	// make sure the role exist
	var role Role
	res := a.DB.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}
	}

	// check if the role is already assigned
	var userRole UserRole
	res = a.DB.Where("user_id = ?", userID).Where("role_id = ?", role.ID).First(&userRole)
	if res.Error == nil {
		//found a record, this role is already assigned to the same user
		return ErrRoleAlreadyAssigned
	}

	// assign the role
	a.DB.Create(&UserRole{UserID: userID, RoleID: role.ID})

	return nil
}

func (a *Authority) CheckRole(userID uint, roleName string) (bool, error) {
	// find the role
	var role Role
	res := a.DB.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, ErrRoleNotFound
		}

	}

	// check if the role is a assigned
	var userRole UserRole
	res = a.DB.Where("user_id = ?", userID).Where("role_id = ?", role.ID).First(&userRole)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

	}

	return true, nil
}

func (a *Authority) CheckPermission(userID uint, permName string) (bool, error) {
	// the user role
	var userRoles []UserRole
	res := a.DB.Where("user_id = ?", userID).Find(&userRoles)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
	}

	//prepare an array of role ids
	var roleIDs []uint
	for _, r := range userRoles {
		roleIDs = append(roleIDs, r.RoleID)
	}

	// find the permission
	var perm UserPermission
	res = a.DB.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, ErrPermissionNotFound
		}

	}

	// find the role permission
	var rolePermission RolePermission
	res = a.DB.Where("role_id IN (?)", roleIDs).Where("permission_id = ?", perm.ID).First(&rolePermission)
	if res.Error != nil {
		return false, nil
	}

	return true, nil
}

func (a *Authority) CheckRolePermission(roleName string, permName string) (bool, error) {
	// find the role
	var role Role
	res := a.DB.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, ErrRoleNotFound
		}

	}

	// find the permission
	var perm UserPermission
	res = a.DB.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, ErrPermissionNotFound
		}

	}

	// find the rolePermission
	var rolePermission RolePermission
	res = a.DB.Where("role_id = ?", role.ID).Where("permission_id = ?", perm.ID).First(&rolePermission)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

	}

	return true, nil
}

// it returns a error in case of any
func (a *Authority) RevokeRole(userID uint, roleName string) error {
	// find the role
	var role Role
	res := a.DB.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	// revoke the role
	a.DB.Where("user_id = ?", userID).Where("role_id = ?", role.ID).Delete(UserRole{})

	return nil
}

func (a *Authority) RevokePermission(userID uint, permName string) error {
	// revoke the permission from all roles of the user
	// find the user roles
	var userRoles []UserRole
	res := a.DB.Where("user_id = ?", userID).Find(&userRoles)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil
		}

	}

	// find the permission
	var perm UserPermission
	res = a.DB.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrPermissionNotFound
		}

	}

	for _, r := range userRoles {
		// revoke the permission
		a.DB.Where("role_id = ?", r.RoleID).Where("permission_id = ?", perm.ID).Delete(RolePermission{})
	}

	return nil
}

func (a *Authority) RevokeRolePermission(roleName string, permName string) error {
	// find the role
	var role Role
	res := a.DB.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	// find the permission
	var perm UserPermission
	res = a.DB.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrPermissionNotFound
		}

	}

	// revoke the permission
	a.DB.Where("role_id = ?", role.ID).Where("permission_id = ?", perm.ID).Delete(RolePermission{})

	return nil
}

func (a *Authority) GetRoles() ([]string, error) {
	var result []string
	var roles []Role
	a.DB.Find(&roles)

	for _, role := range roles {
		result = append(result, role.Name)
	}

	return result, nil
}

// GetUserRoles returns all user assigned roles
func (a *Authority) GetUserRoles(userID uint) ([]string, error) {
	var result []string
	var userRoles []UserRole
	a.DB.Where("user_id = ?", userID).Find(&userRoles)

	for _, r := range userRoles {
		var role Role
		// for every user role get the role name
		res := a.DB.Where("id = ?", r.RoleID).Find(&role)
		if res.Error == nil {
			result = append(result, role.Name)
		}
	}

	return result, nil
}

func (a *Authority) GetPermissions() ([]string, error) {
	var result []string
	var perms []UserPermission
	a.DB.Find(&perms)

	for _, perm := range perms {
		result = append(result, perm.Name)
	}

	return result, nil
}

func (a *Authority) DeleteRole(roleName string) error {
	// find the role
	var role Role
	res := a.DB.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	// check if the role is assigned to a user
	var userRole UserRole
	res = a.DB.Where("role_id = ?", role.ID).First(&userRole)
	if res.Error == nil {
		// role is assigned
		return ErrRoleInUse
	}

	// revoke the assignment of permissions before deleting the role
	a.DB.Where("role_id = ?", role.ID).Delete(RolePermission{})

	// delete the role
	a.DB.Where("name = ?", roleName).Delete(Role{})

	return nil
}

func (a *Authority) DeletePermission(permName string) error {
	// find the permission
	var perm UserPermission
	res := a.DB.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrPermissionNotFound
		}

	}

	// check if the permission is assigned to a role
	var rolePermission RolePermission
	res = a.DB.Where("permission_id = ?", perm.ID).First(&rolePermission)
	if res.Error == nil {
		// role is assigned
		return ErrPermissionInUse
	}

	// delete the permission
	a.DB.Where("name = ?", permName).Delete(UserPermission{})

	return nil
}

func migrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&Role{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&UserPermission{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&RolePermission{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&UserRole{})
	if err != nil {
		return
	}

}
