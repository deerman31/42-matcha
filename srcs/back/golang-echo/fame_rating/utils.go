package famerating

import (
	"database/sql"
	"fmt"
	"time"
)

func fakeAccountCount(tx *sql.Tx, userID int) (int, error) {
	const query = `
	SELECT COUNT(*)
FROM report_fake_accounts
WHERE fake_account_id = $1;`
	var count int
	if err := tx.QueryRow(query, userID).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting fake accounts: %w", err)
	}
	return count, nil
}

func blockCount(tx *sql.Tx, userID int) (int, error) {
	const query = `
	SELECT COUNT(*)
FROM user_blocks
WHERE blocked_id = $1;`
	var count int
	if err := tx.QueryRow(query, userID).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting blocks: %w", err)
	}
	return count, nil
}

func friendCount(tx *sql.Tx, userID int) (int, error) {
	const query = `
	SELECT COUNT(*)
FROM user_friends
WHERE user_id1 = $1 OR user_id2 = $1;`
	var count int
	if err := tx.QueryRow(query, userID).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting friends: %w", err)
	}
	return count, nil
}

func likedCount(tx *sql.Tx, userID int, duration time.Duration) (int, error) {
	const query = `
		SELECT COUNT(*)
		FROM user_likes
		WHERE liked_id = $1
		AND created_at >= CURRENT_TIMESTAMP - $2::interval`
	var count int
	if err := tx.QueryRow(query, userID, duration.String()).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting likes: %w", err)
	}
	return count, nil
}

func viewedCount(tx *sql.Tx, userID int, duration time.Duration) (int, error) {
	const query = `
		SELECT COUNT(*)
		FROM profile_views
		WHERE viewed_id = $1
		AND viewed_at >= CURRENT_TIMESTAMP - $2::interval`
	var count int
	if err := tx.QueryRow(query, userID, duration.String()).Scan(&count); err != nil {
		return 0, fmt.Errorf("counting profile views: %w", err)
	}
	return count, nil
}
