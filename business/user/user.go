package user

import "time"

//User product User that available to rent or sell
type User struct {
	ID         int
	Name       string
	Username   string
	Password   string
	Address   string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
	Version    int
}

//NewUser create new User
func NewUser(
	name string,
	username string,
	password string,
	creator string,
	createdAt time.Time) User {

	return User{
		Name:       name,
		Username:   username,
		Password:   password,
		CreatedAt:  createdAt,
		CreatedBy:  creator,
		ModifiedAt: createdAt,
		ModifiedBy: creator,
		Version:    1,
	}
}

//ModifyUser update existing User data
func (oldData *User) ModifyUser(newName string, modifiedAt time.Time, updater string) User {
	return User{
		ID:         oldData.ID,
		Name:       newName,
		Username:   oldData.Username,
		Password:   oldData.Password,
		CreatedAt:  oldData.CreatedAt,
		CreatedBy:  oldData.CreatedBy,
		ModifiedAt: modifiedAt,
		ModifiedBy: updater,
		Version:    oldData.Version + 1,
	}
}

//ModifyUser Address update existing User data
func (oldData *User) ModifyUserAdress(address string, modifiedAt time.Time, updater string) User {
	return User{
		ID:         oldData.ID,
		Name:       oldData.Name,
		Username:   oldData.Username,
		Password:   oldData.Password,
		Address: 	address,
		CreatedAt:  oldData.CreatedAt,
		CreatedBy:  oldData.CreatedBy,
		ModifiedAt: modifiedAt,
		ModifiedBy: updater,
		Version:    oldData.Version + 1,
	}
}
