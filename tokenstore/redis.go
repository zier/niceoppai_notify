package tokenstore

import "github.com/go-redis/redis"

// Store ...
type Store struct {
	RedisKey string
	Redis    *redis.Client
	List     []string
}

// New ...
func New(redis *redis.Client) (*Store, error) {
	s := &Store{
		RedisKey: "tokens",
		Redis:    redis,
		List:     []string{},
	}

	if s.Redis.Exists(s.RedisKey).Val() == 0 {
		return s, nil
	}

	list, err := s.Redis.LRange(s.RedisKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	s.List = list
	return s, nil
}

func (s *Store) isTokenExist(token string) bool {
	for _, t := range s.List {
		if token == t {
			return true
		}
	}

	return false
}

func (s *Store) redisSync() error {
	if err := s.Redis.Del(s.RedisKey).Err(); err != nil {
		return err
	}

	for _, val := range s.List {
		if err := s.Redis.LPush(s.RedisKey, val).Err(); err != nil {
			return err
		}
	}

	return nil
}

// Save ...
func (s *Store) Save(token string) error {
	if s.isTokenExist(token) {
		return nil
	}

	s.List = append(s.List, token)
	return s.redisSync()
}

// Remove ...
func (s *Store) Remove(token string) error {
	for i, v := range s.List {
		if v == token {
			s.List = append(s.List[:i], s.List[i+1:]...)
			break
		}
	}

	return s.redisSync()
}

// All ...
func (s *Store) All() ([]string, error) {
	return s.List, nil
}
