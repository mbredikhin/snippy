package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
	"github.com/redis/go-redis/v9"
)

const (
	blacklisted_tokens = "blacklisted_tokens"
)

type AuthRepo struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewAuthRepo(db *sqlx.DB, rdb *redis.Client) *AuthRepo {
	return &AuthRepo{
		db:  db,
		rdb: rdb,
	}
}

func (r *AuthRepo) CreateUser(user snippets.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetUser - get user from db
func (r *AuthRepo) GetUser(username, password string) (snippets.User, error) {
	var user snippets.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthRepo) BlacklistToken(token string, expiresAt int64) error {
	ctx := context.Background()
	return r.rdb.ZAdd(ctx, blacklisted_tokens, redis.Z{
		Score:  float64(expiresAt),
		Member: token,
	}).Err()
}

func (r *AuthRepo) CheckIfTokenBlacklisted(token string) bool {
	ctx := context.Background()
	if err := r.rdb.ZRankWithScore(ctx, blacklisted_tokens, token).Err(); err == redis.Nil {
		return false
	}
	return true
}

func (r *AuthRepo) RemoveExpiredTokensFromBlacklist(timestamp int64) error {
	ctx := context.Background()
	return r.rdb.ZRemRangeByScore(ctx, blacklisted_tokens, "-inf", fmt.Sprint(timestamp)).Err()
}
