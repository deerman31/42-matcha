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

)
