package tokenstore

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
)

type TokenStoreSuite struct {
	suite.Suite
	store *Store
}

func TestTokenStoreSuite(t *testing.T) {
	suite.Run(t, &TokenStoreSuite{})
}

func (t *TokenStoreSuite) SetupSuite() {
	r := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	s, err := New(r)
	if err != nil {
		panic(err)
	}

	t.store = s
}

func (t *TokenStoreSuite) TearDownTest() {
	t.store.List = []string{}
	t.store.Redis.Del(t.store.RedisKey).Err()
}

func (t *TokenStoreSuite) TestSave() {
	err := t.store.Save("new_token_1")
	t.NoError(err)

	err = t.store.Save("new_token_2")
	t.NoError(err)

	err = t.store.Save("new_token_3")
	t.NoError(err)

	tokens, err := t.store.All()
	t.NoError(err)
	t.Equal([]string{"new_token_1", "new_token_2", "new_token_3"}, tokens)
}

func (t *TokenStoreSuite) TestRemove() {
	err := t.store.Save("new_token_1")
	t.NoError(err)

	err = t.store.Save("new_token_2")
	t.NoError(err)

	err = t.store.Save("new_token_3")
	t.NoError(err)

	err = t.store.Remove("new_token_2")
	t.NoError(err)

	tokens, err := t.store.All()
	t.NoError(err)
	t.Equal([]string{"new_token_1", "new_token_3"}, tokens)
}
