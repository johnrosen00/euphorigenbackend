package sessions

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

//RedisStore represents a session.Store backed by redis.
type RedisStore struct {
	//Redis client used to talk to redis server.
	Client *redis.Client
	//Used for key expiry time on redis.
	SessionDuration time.Duration
	Context         context.Context
}

//NewRedisStore constructs a new RedisStore
func NewRedisStore(ctx context.Context, client *redis.Client, sessionDuration time.Duration) *RedisStore {
	//initialize and return a new RedisStore struct
	return &RedisStore{client, sessionDuration, ctx}
}

//Store implementation

//Save saves the provided `sessionState` and associated SessionID to the store.
//The `sessionState` parameter is typically a pointer to a struct containing
//all the data you want to associated with the given SessionID.
func (rs *RedisStore) Save(sid SessionID, sessionState interface{}) error {
	// marshals the `sessionState` to JSON and save it in the redis database,
	//using `sid.getRedisKey()` for the key.
	//return any errors that occur along the way.
	j, err := json.Marshal(sessionState)

	if err != nil {
		return err
	}

	sidKey := sid.getRedisKey()
	rs.Client.Set(rs.Context, sidKey, j, rs.SessionDuration)
	return nil
}

//Get populates `sessionState` with the data previously saved
//for the given SessionID
func (rs *RedisStore) Get(sid SessionID, sessionState interface{}) error {
	//gets the previously-saved session state data from redis,
	//unmarshal it back into the `sessionState` parameter
	//and reset the expiry time, so that it doesn't get deleted until
	//the SessionDuration has elapsed.

	sidKey := sid.getRedisKey()

	v := rs.Client.Get(rs.Context, sidKey)

	if v.Err() != nil {
		return ErrStateNotFound
	}

	vB, err := v.Bytes()
	if err != nil {
		return errors.New("no session state was found in the session store")
	}

	err1 := json.Unmarshal(vB, sessionState)
	rs.Client.Set(rs.Context, sidKey, rs.Client.Get(rs.Context, sidKey), rs.SessionDuration)

	return err1
}

//Delete deletes all state data associated with the SessionID from the store.
func (rs *RedisStore) Delete(sid SessionID) error {
	sidKey := sid.getRedisKey()

	if ret := rs.Client.Del(rs.Context, sidKey).Err(); ret != nil {
		return ErrStateNotFound
	}

	return nil
}

//getRedisKey() returns the redis key to use for the SessionID
func (sid SessionID) getRedisKey() string {
	//convert the SessionID to a string and add the prefix "sid:" to keep
	//SessionID keys separate from other keys that might end up in this
	//redis instance
	return "sid:" + sid.String()
}
