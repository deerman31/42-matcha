package myprofile

const (
	userInfoQuery = `
SELECT 
    u.username,
    u.email,
    ui.lastname,
    ui.firstname,
    ui.birthdate,
    ui.gender,
    ui.sexuality,
    ui.area,
    ui.self_intro
FROM 
    users u
    INNER JOIN user_info ui ON u.id = ui.user_id
WHERE 
    u.id = $1
`

	query2 = `
SELECT 
    ut.user_id,
    ARRAY_AGG(t.tag_name) as tags
FROM 
    user_tags ut
    INNER JOIN tags t ON ut.tag_id = t.id
WHERE 
    ut.user_id = $1  -- ここにユーザーIDを指定
GROUP BY 
    ut.user_id;
`

	// PostGISのgeography型からポイントを取得するクエリ
	getUserLocationQuery = `
SELECT 
    CASE 
        WHEN is_gps = true THEN ST_Y(location::geometry)
        ELSE ST_Y(location_alternative::geometry)
    END as latitude,
    CASE 
        WHEN is_gps = true THEN ST_X(location::geometry)
        ELSE ST_X(location_alternative::geometry)
    END as longitude,
    is_gps
FROM 
    user_location 
WHERE 
    user_id = $1
    `

	// 月間プロフィール閲覧数を取得するクエリ
	getMonthlyProfileViewsQuery = `
SELECT 
	COUNT(DISTINCT viewer_id) as view_count
FROM 
	profile_views
WHERE 
	viewed_id = $1
	AND viewed_at >= CURRENT_TIMESTAMP - INTERVAL '1 month'`

	// 月間Like数を取得するクエリ
	getMonthlyLikesQuery = `
SELECT 
    COUNT(DISTINCT liker_id) as like_count
FROM 
    user_likes
WHERE 
    liked_id = $1
    AND created_at >= CURRENT_TIMESTAMP - INTERVAL '1 month'`

	// フレンド数を取得するクエリ
	getFriendCountQuery = `
SELECT 
	(
		SELECT COUNT(*) 
		FROM user_friends 
		WHERE user_id1 = $1 
		OR user_id2 = $1
	) as friend_count`

	// ブロック数を取得するクエリ
	getBlockedCountQuery = `
    SELECT 
        COUNT(DISTINCT blocker_id) as block_count
    FROM 
        user_blocks
    WHERE 
        blocked_id = $1`

	getFakeAccountReportsQuery = `
    SELECT 
        COUNT(DISTINCT reporter_id) as report_count
    FROM 
        report_fake_accounts
    WHERE 
        fake_account_id = $1`
)
