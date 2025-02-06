package auth

const (
	// ユーザー名からユーザー情報を取得するクエリ
	selectUserByUsernameQuery = `
        SELECT id, username, password_hash, is_online, is_registered, is_preparation
        FROM users 
        WHERE username = $1
        LIMIT 1
    `

	// ユーザーのオンラインステータスを更新するクエリ
	updateUserOnlineStatusQuery = `
        UPDATE users 
        SET is_online = TRUE 
        WHERE id = $1
	`

	// ユーザーのオンラインステータスを更新するクエリ
	updateUserOfflineStatusQuery = `
        UPDATE users 
        SET is_online = FALSE 
        WHERE id = $1
    `

	// 1つのクエリで両方をチェック
	checkDuplicateCredentialsQuery = `
        SELECT 
            EXISTS(SELECT 1 FROM users WHERE username = $1) as username_exists,
            EXISTS(SELECT 1 FROM users WHERE email = $2) as email_exists
    `
	// 新規ユーザーを登録するためのクエリ
	insertNewUserQuery = `
        INSERT INTO users (
            username, 
            email, 
            password_hash
        ) VALUES ($1, $2, $3)
		 RETURNING id
    `
)
