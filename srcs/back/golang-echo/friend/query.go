package friend

const (
	queryGetFriendList = `
SELECT DISTINCT u.username 
FROM user_friends uf
JOIN users u ON 
   CASE 
       WHEN uf.user_id1 = $1 THEN uf.user_id2 
       WHEN uf.user_id2 = $1 THEN uf.user_id1
   END = u.id
WHERE user_id1 = $1 OR user_id2 = $1
ORDER BY uf.created_at DESC
`


	queryGetUserID = `SELECT id FROM users WHERE username = $1;`

	queryDeleteFriend = `
		DELETE FROM user_friends
		WHERE user_id1 = $1 AND user_id2 = $2
	`
)
