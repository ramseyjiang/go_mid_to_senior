package finduser

import (
	"math/rand"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUserListProxy(t *testing.T) {
	mockedDatabase := UserList{}

	rand.Seed(8888888) // The preceding test creates a user list of million users with random IDs.
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, User{ID: n})
	}

	// a proxy object composed of a mock database with the million users, and a cache implemented as a FIFO stack with a size of 2.
	proxy := UserListProxy{
		MockedDatabase: &mockedDatabase,
		StackCache:     UserList{},
		StackSize:      2,
	}

	// took the fourth, fifth, and sixth IDs from the slice
	knownIDs := [3]int32{mockedDatabase[3].ID, mockedDatabase[4].ID, mockedDatabase[5].ID}

	// t.Run() is a test closure.
	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		// check whether the returned user has the same ID as that of the expected user at index 0 of the knownIDs slice
		assert.Equal(t, user.ID, knownIDs[0])
		assert.Equal(t, 1, len(proxy.StackCache))
		assert.Equal(t, false, proxy.LastSearchUsedCache)
		assert.Equal(t, posInDB, user.Pos)
	})

	// The second embedded test for the Proxy pattern is to ask for the same user as before, which must now be returned from the cache.
	t.Run("FindUser - One user asking for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, user.ID, knownIDs[0])
		assert.Equal(t, 1, len(proxy.StackCache))
		assert.Equal(t, true, proxy.LastSearchUsedCache)
		assert.Equal(t, posInCache, user.Pos)
	})

	t.Run("FindUser - overflowing stack", func(t *testing.T) {
		user1, _ := proxy.FindUser(knownIDs[0])
		assert.Equal(t, posInCache, user1.Pos)

		user2, _ := proxy.FindUser(knownIDs[1])
		assert.Equal(t, false, proxy.LastSearchUsedCache)
		assert.Equal(t, posInDB, user2.Pos)

		user3, _ := proxy.FindUser(knownIDs[2])
		assert.Equal(t, false, proxy.LastSearchUsedCache)
		assert.Equal(t, posInDB, user3.Pos)

		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}

		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users the cache should not grow" + " more than to two")
		}

		for _, v := range proxy.StackCache {
			if v.ID != user2.ID && v.ID != user3.ID {
				t.Error("A non expected user was found on the cache")
			}
		}
	})
}
