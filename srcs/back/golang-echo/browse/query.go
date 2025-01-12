package browse

const (
	query = `
WITH my_location AS (
 SELECT 
     CASE 
         WHEN is_gps = true THEN location
         ELSE location_alternative
     END as geom
 FROM user_location 
 WHERE user_id = $1
),
my_info AS (
 SELECT gender, sexuality 
 FROM user_info 
 WHERE user_id = $1
),
my_tags AS (
 SELECT array_agg(tag_id) as tag_ids
 FROM user_tags
 WHERE user_id = $1
),
user_tag_groups AS (
 SELECT 
     ut.user_id,
     array_agg(ut.tag_id) as tag_ids
 FROM user_tags ut
 GROUP BY ut.user_id
),
profile_views_count AS (
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
 SELECT 
   user_id1 as user_id, COUNT(*) * 5 as friend_points
 FROM user_friends
 GROUP BY user_id1
 UNION ALL
 SELECT 
   user_id2 as user_id, COUNT(*) * 5 as friend_points
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
),
fame_rating AS (
 SELECT 
   COALESCE(pv.viewed_id, l.liked_id, f.user_id, b.blocked_id, fr.fake_account_id) as user_id,
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
   END as rating
 FROM profile_views_count pv
 FULL OUTER JOIN like_count l ON pv.viewed_id = l.liked_id
 FULL OUTER JOIN friend_count f ON COALESCE(pv.viewed_id, l.liked_id) = f.user_id
 FULL OUTER JOIN block_count b ON COALESCE(pv.viewed_id, l.liked_id, f.user_id) = b.blocked_id
 FULL OUTER JOIN fake_report_count fr ON COALESCE(pv.viewed_id, l.liked_id, f.user_id, b.blocked_id) = fr.fake_account_id
)
SELECT 
    u.username,
    ui.birthdate,
	ui.area,
    img.profile_image_path1,
    cardinality(
        ARRAY(
            SELECT UNNEST(COALESCE(utg.tag_ids, ARRAY[]::integer[])) 
            INTERSECT 
            SELECT UNNEST((SELECT tag_ids FROM my_tags))
        )
    ) as common_tag_count,
	ROUND(CAST(ST_Distance(
        CASE 
            WHEN ul.is_gps = true THEN ul.location
            ELSE ul.location_alternative
        END,
        (SELECT geom FROM my_location)
    ) / 1000.0 AS NUMERIC), 2) as distance_km,
    COALESCE(fr.rating, 0) as fame_rating


FROM user_info ui
JOIN users u ON ui.user_id = u.id
LEFT JOIN user_image img ON ui.user_id = img.user_id
LEFT JOIN user_location ul ON ui.user_id = ul.user_id
LEFT JOIN user_tag_groups utg ON ui.user_id = utg.user_id
LEFT JOIN fame_rating fr ON ui.user_id = fr.user_id
CROSS JOIN my_info mi 
WHERE ui.user_id != $1
AND u.is_registered = true
AND u.is_preparation = true
AND EXTRACT(YEAR FROM AGE(CURRENT_DATE, ui.birthdate)) BETWEEN $4 AND $5
AND cardinality(
 ARRAY(
     SELECT UNNEST(COALESCE(utg.tag_ids, ARRAY[]::integer[])) 
     INTERSECT 
     SELECT UNNEST((SELECT tag_ids FROM my_tags))
 )
) >= $6
AND COALESCE(fr.rating, 0) >= $7
AND (
 (mi.gender::text = ui.sexuality::text OR ui.sexuality = 'male/female')
 AND
 (mi.sexuality::text = ui.gender::text OR mi.sexuality = 'male/female')
)
AND ST_Distance(
 CASE 
     WHEN ul.is_gps = true THEN ul.location
     ELSE ul.location_alternative
 END,
 (SELECT geom FROM my_location)
) BETWEEN $2 AND $3
AND NOT EXISTS (
 SELECT 1 FROM user_blocks
 WHERE (blocker_id = $1 AND blocked_id = ui.user_id)
    OR (blocker_id = ui.user_id AND blocked_id = $1)
)
AND NOT EXISTS (
 SELECT 1 FROM report_fake_accounts
 WHERE (reporter_id = $1 AND fake_account_id = ui.user_id)
    OR (reporter_id = ui.user_id AND fake_account_id = $1)
)
AND NOT EXISTS (
 SELECT 1 FROM user_friends
 WHERE (user_id1 = LEAST($1, ui.user_id) AND user_id2 = GREATEST($1, ui.user_id))
)
AND NOT EXISTS (
 SELECT 1 FROM user_likes
 WHERE liker_id = $1 AND liked_id = ui.user_id
)	

	`
)
