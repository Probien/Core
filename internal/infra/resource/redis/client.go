package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JairDavid/Probien-Backend/internal/domain/dto"
	"github.com/JairDavid/Probien-Backend/internal/infra/component"
	"github.com/go-redis/redis/v8"
	"github.com/satori/go.uuid"
	"time"
)

type Client struct {
	conn *redis.Client
}

func New(opts ...string) *Client {

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     opts[0],
		Password: opts[1],
		DB:       0,
	})

	errHealthCheck := redisConnection.Ping(redisConnection.Context())
	if errHealthCheck.Err() != nil {
		panic(errors.New("Redis connection failed: " + errHealthCheck.Err().Error()))
	}

	return &Client{
		conn: redisConnection,
	}
}

func (c *Client) GenerateSessionID(employee *dto.Employee, session chan<- component.SessionCredential) {
	sessionID := uuid.NewV4()
	sessionClaims := component.SessionCredential{
		ID:        sessionID.String(),
		Username:  employee.Email,
		ExpiresAt: time.Now().Add(time.Minute * 30),
	}
	sessionBytes, err := json.Marshal(sessionClaims)
	if err != nil {
		fmt.Errorf("error marshalling session: %v", err)
	}
	cmd := c.conn.Set(context.Background(), sessionID.String(), string(sessionBytes[:]), time.Minute*30)
	if err := cmd.Err(); err != nil {
		fmt.Errorf("error writing session to Redis: %v", err)
	}
	session <- sessionClaims
	close(session)
	return
}

func (c *Client) ClearSessionID(cookie string) error {
	c.conn.Del(context.Background(), cookie)
	return nil
}

func (c *Client) ExistCookie(cookie string, checker chan<- bool) {
	var sessionID component.SessionCredential
	val := c.conn.Get(context.Background(), cookie).Val()
	err := json.Unmarshal([]byte(val), &sessionID)
	if err != nil {
		fmt.Errorf("error getting session from Redis: %v", err)
	}
	checker <- val != "" && cookie == sessionID.ID
	close(checker)
	return
}
