package finduser

import "fmt"

type User struct {
	ID  int32
	Pos string
}

const posInCache = "Returning user from cache"
const posInDB = "Returning user from database"

// The UserFinder is the subject interface.
type UserFinder interface {
	FindUser(id string) (User, error)
}

// UserList is a type of slice of users.
type UserList []User

// FindUser iterates over the list to try to find a user with the same ID that the param or returns an error if it can't find it.
func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}

	return User{}, fmt.Errorf("User %d could not be found\n", id)
}

// AddUser adds a new user to the end of the Users slice
func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

// UserListProxy is the proxy object, is composed of a UserList slice, which is our database representation.
// The StackCache is also a UserList type for simplicity.
// The StackSize to give our stack the size we want, it will cache a maximum of StackSize users and rotate the cache if it reaches this limit.
// The LastSearchUsedCache will hold if the last performed search has used the cache, or has accessed the database
// The MockedDatabase is a pointer, because it will be used to store 1 million users. So it uses a pointer to reference is better than value.
type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

// addUserToStack takes the user argument and adds it to the stack in place.
// If the stack is full it removes the first element on it before adding, follow FIFO rules.
func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

// FindUser will search for the specified ID in the cache list.
// If it finds it, it will return the ID. If not, it will search in the mock database.
// Finally, if it's not in the mock database., it will return an error (generated from the mock database)
func (u *UserListProxy) FindUser(id int32) (User, error) {
	// Search for the object in the cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		user.Pos = posInCache
		u.LastSearchUsedCache = true
		return user, nil
	}

	// Object is not in the cache list. Search in the heavy list
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	// Adds the new user to the stack, removing the last if necessary
	u.addUserToStack(user)

	user.Pos = posInDB
	u.LastSearchUsedCache = false
	return user, nil
}
