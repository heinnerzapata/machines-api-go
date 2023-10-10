package db

import (
	"context"
	"machines-api-go/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		FirstName: util.RandomString(8),
		LastName:  util.RandomString(8),
		Password:  util.RandomString(8),
		Email:     util.RandomString(8),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestCreateAccount(t *testing.T) {
	user := createRandomUser(t)

	require.NotEmpty(t, user)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.Equal(t, user1, user2)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:        user1.ID,
		FirstName: "Braham",
		LastName:  user1.LastName,
		Email:     user1.Email,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	user3, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user3)

	require.Equal(t, user2.FirstName, user3.FirstName)

}
