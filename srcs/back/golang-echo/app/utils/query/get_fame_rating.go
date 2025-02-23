package query

import "database/sql"

const queryGetFameRatingByUserID = `
WITH profile_views_count AS (
    SELECT viewed_id, COUNT(*) as view_count
    FROM profile_views
    WHERE viewed_at >= CURRENT_TIMESTAMP - INTERVAL '1 month'
    GROUP BY viewed_id
),
like_count AS (
    SELECT liked_id, COUNT(*) * 3 as like_points
    FROM user_likes
    WHERE created_at >= CURRENT_TIMESTAMP - INTERVAL '1 month'
    GROUP BY liked_id
),
friend_count AS (
    SELECT user_id1 as user_id, COUNT(*) * 5 as friend_points
    FROM user_friends
    GROUP BY user_id1
    UNION ALL
    SELECT user_id2 as user_id, COUNT(*) * 5 as friend_points
    FROM user_friends
    GROUP BY user_id2
),
block_count AS (
    SELECT blocked_id, COUNT(*) * -5 as block_points
    FROM user_blocks
    GROUP BY blocked_id
),
fake_report_count AS (
    SELECT fake_account_id, COUNT(*) * -5 as fake_points
    FROM report_fake_accounts
    GROUP BY fake_account_id
)
SELECT 
    CASE 
        WHEN (COALESCE(pv.view_count, 0) + 
              COALESCE(l.like_points, 0) + 
              COALESCE(f.friend_points, 0) + 
              COALESCE(b.block_points, 0) + 
              COALESCE(fr.fake_points, 0)) >= 100 THEN 5
        WHEN (COALESCE(pv.view_count, 0) + 
              COALESCE(l.like_points, 0) + 
              COALESCE(f.friend_points, 0) + 
              COALESCE(b.block_points, 0) + 
              COALESCE(fr.fake_points, 0)) >= 80 THEN 4
        WHEN (COALESCE(pv.view_count, 0) + 
              COALESCE(l.like_points, 0) + 
              COALESCE(f.friend_points, 0) + 
              COALESCE(b.block_points, 0) + 
              COALESCE(fr.fake_points, 0)) >= 60 THEN 3
        WHEN (COALESCE(pv.view_count, 0) + 
              COALESCE(l.like_points, 0) + 
              COALESCE(f.friend_points, 0) + 
              COALESCE(b.block_points, 0) + 
              COALESCE(fr.fake_points, 0)) >= 40 THEN 2
        WHEN (COALESCE(pv.view_count, 0) + 
              COALESCE(l.like_points, 0) + 
              COALESCE(f.friend_points, 0) + 
              COALESCE(b.block_points, 0) + 
              COALESCE(fr.fake_points, 0)) >= 20 THEN 1
        ELSE 0
    END as fame_rating
FROM profile_views_count pv
FULL OUTER JOIN like_count l ON pv.viewed_id = l.liked_id
FULL OUTER JOIN friend_count f ON COALESCE(pv.viewed_id, l.liked_id) = f.user_id
FULL OUTER JOIN block_count b ON COALESCE(pv.viewed_id, l.liked_id, f.user_id) = b.blocked_id
FULL OUTER JOIN fake_report_count fr ON COALESCE(pv.viewed_id, l.liked_id, f.user_id, b.blocked_id) = fr.fake_account_id
WHERE COALESCE(pv.viewed_id, l.liked_id, f.user_id, b.blocked_id, fr.fake_account_id) = $1;
`

func GetFameRating(tx *sql.Tx, userID int) (int, error) {
	fameRating := 0
	err := tx.QueryRow(queryGetFameRatingByUserID, userID).Scan(&fameRating)
	if err == sql.ErrNoRows {
		// 結果がない場合は0を返す
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return fameRating, nil
}
