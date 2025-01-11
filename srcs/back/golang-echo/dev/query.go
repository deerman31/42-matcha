package dev

const (
	query = `
-- ユーザー登録用の関数を作成
CREATE OR REPLACE FUNCTION register_user(
    p_username VARCHAR(255),
    p_email VARCHAR(255),
    p_password_hash VARCHAR(255),
    p_is_online BOOLEAN,
    p_is_registered BOOLEAN,
    p_is_preparation BOOLEAN,
    -- user_info用のパラメータ
    p_lastname VARCHAR(50),
    p_firstname VARCHAR(50),
    p_birthdate DATE,
    p_gender gender_type,
    p_sexuality sexuality_type,
    p_area prefecture,
    p_self_intro VARCHAR(300),
    -- user_image用のパラメータ
    p_profile_image_path1 VARCHAR(255)
) RETURNS INTEGER AS $$
DECLARE
    new_user_id INTEGER;
BEGIN
    -- usersテーブルへの挿入
    INSERT INTO users (
        username,
        email,
        password_hash,
        is_online,
        is_registered,
        is_preparation
    ) VALUES (
        p_username,
        p_email,
        p_password_hash,
        p_is_online,
        p_is_registered,
        p_is_preparation
    ) RETURNING id INTO new_user_id;

    -- user_infoテーブルへの挿入
    -- トリガーで自動的に作成されるレコードを更新
    UPDATE user_info SET
        lastname = p_lastname,
        firstname = p_firstname,
        birthdate = p_birthdate,
        gender = p_gender,
        sexuality = p_sexuality,
        area = p_area,
        self_intro = p_self_intro
    WHERE user_id = new_user_id;

    -- user_imageテーブルへの挿入
    -- トリガーで自動的に作成されるレコードを更新
    UPDATE user_image SET
        profile_image_path1 = p_profile_image_path1
    WHERE user_id = new_user_id;

    RETURN new_user_id;
END;
$$ LANGUAGE plpgsql;
`
)
