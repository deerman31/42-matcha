package like

const (
	queryGetUserID = `SELECT id FROM users WHERE username = $1;`

	queryCheckLikeExists = `
    SELECT EXISTS (
        SELECT 1 
        FROM user_likes 
        WHERE liker_id = $1 AND liked_id = $2
    )
`
	queryCheckFriendExists = `
    SELECT EXISTS (
        SELECT 1 
        FROM user_friends 
        WHERE user_id1 = $1 AND user_id2 = $2
    )
`
	queryDeleteLike = `
		DELETE FROM user_likes
		WHERE liker_id = $1 AND liked_id = $2
	`

	queryInsertLike = `
	INSERT INTO user_likes (liker_id, liked_id)
	VALUES ($1, $2)
	`

	queryInsertFriend = `
	INSERT INTO user_friends (user_id1, user_id2)
	VALUES ($1, $2)
	`

	queryGetLikedUsers = `
        SELECT DISTINCT u.username 
        FROM user_likes ul
        JOIN users u ON ul.liked_id = u.id
        WHERE ul.liker_id = $1
        ORDER BY ul.created_at DESC
    `

	queryGetLikerUsers = `
        SELECT DISTINCT u.username 
        FROM user_likes ul
        JOIN users u ON ul.liker_id = u.id
        WHERE ul.liked_id = $1
        ORDER BY ul.created_at DESC
    `
)
