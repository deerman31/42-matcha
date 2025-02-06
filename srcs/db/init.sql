-- 1. PostGIS拡張を最初に有効化
CREATE EXTENSION IF NOT EXISTS postgis;

-- userはuserの基本的な情報が入っているテーブル
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_online BOOLEAN DEFAULT FALSE,
    is_registered BOOLEAN DEFAULT TRUE, -- is_registeredはsignup後にメールで認証したかどうかを表すものだが、開発の最初ではスピードを重視し、でふぉるとでtrue
    is_preparation BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster lookups
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);

-- usersテーブルの更新時刻を自動更新する関数
CREATE OR REPLACE FUNCTION update_users_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- usersテーブルの更新時刻自動更新トリガー
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_users_updated_at_column();


-- 都道府県のENUM型を作成
CREATE TYPE prefecture AS ENUM (
    'Hokkaido', 'Aomori', 'Iwate', 'Miyagi', 'Akita', 'Yamagata', 'Fukushima',
    'Ibaraki', 'Tochigi', 'Gunma', 'Saitama', 'Chiba', 'Tokyo', 'Kanagawa',
    'Niigata', 'Toyama', 'Ishikawa', 'Fukui', 'Yamanashi', 'Nagano',
    'Gifu', 'Shizuoka', 'Aichi', 'Mie',
    'Shiga', 'Kyoto', 'Osaka', 'Hyogo', 'Nara', 'Wakayama',
    'Tottori', 'Shimane', 'Okayama', 'Hiroshima', 'Yamaguchi',
    'Tokushima', 'Kagawa', 'Ehime', 'Kochi',
    'Fukuoka', 'Saga', 'Nagasaki', 'Kumamoto', 'Oita', 'Miyazaki', 'Kagoshima', 'Okinawa'
);

-- 性別のENUM型を作成
CREATE TYPE gender_type AS ENUM ('male', 'female');

-- 性的指向のENUM型を作成
CREATE TYPE sexuality_type AS ENUM ('male', 'female', 'male/female');

CREATE TABLE IF NOT EXISTS user_info (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    lastname VARCHAR(50) NOT NULL DEFAULT '', -- 姓
    firstname VARCHAR(50) NOT NULL DEFAULT '', -- 名前
    birthdate DATE NOT NULL DEFAULT '2000-04-02', -- 生年月日
    --is_gps BOOLEAN DEFAULT FALSE, -- 位置情報を利用するか
    gender gender_type NOT NULL DEFAULT 'male', -- 性別
    sexuality sexuality_type NOT NULL DEFAULT 'male', -- 性的対象
    area prefecture NOT NULL DEFAULT 'Tokyo', -- 都道府県
    self_intro VARCHAR(300) NOT NULL DEFAULT '', -- 自己紹介
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- インデックスの作成
CREATE INDEX idx_user_info_user_id ON user_info(user_id);

-- user_infosテーブルの更新時刻を自動更新する関数
CREATE OR REPLACE FUNCTION update_user_info_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- user_infoテーブルの更新時刻自動更新トリガー
CREATE TRIGGER update_user_info_updated_at
    BEFORE UPDATE ON user_info
    FOR EACH ROW
    EXECUTE FUNCTION update_user_info_updated_at_column();


-- usersテーブルにレコードが挿入された時、自動的にuser_infoにレコードを作成するトリガー
CREATE OR REPLACE FUNCTION create_user_info_on_user_creation()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO user_info (user_id)
    VALUES (NEW.id);
    RETURN NEW;
END;
$$ language 'plpgsql';

-- トリガーの作成
CREATE TRIGGER create_user_info_after_user_creation
    AFTER INSERT ON users
    FOR EACH ROW
    EXECUTE FUNCTION create_user_info_on_user_creation();


CREATE TABLE IF NOT EXISTS user_image (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    profile_image_path1 VARCHAR(255) DEFAULT NULL, /* プロフィール画像のパス */
    profile_image_path2 VARCHAR(255) DEFAULT NULL, /* プロフィール画像のパス */
    profile_image_path3 VARCHAR(255) DEFAULT NULL, /* プロフィール画像のパス */
    profile_image_path4 VARCHAR(255) DEFAULT NULL, /* プロフィール画像のパス */
    profile_image_path5 VARCHAR(255) DEFAULT NULL, /* プロフィール画像のパス */

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- インデックスの作成
CREATE INDEX idx_user_image_user_id ON user_image(user_id);

-- user_imageテーブルの更新時刻を自動更新する関数
CREATE OR REPLACE FUNCTION update_user_image_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- user_imageテーブルの更新時刻自動更新トリガー
CREATE TRIGGER update_user_image_updated_at
    BEFORE UPDATE ON user_image
    FOR EACH ROW
    EXECUTE FUNCTION update_user_image_updated_at_column();


-- usersテーブルにレコードが挿入された時、自動的にuser_imageにレコードを作成するトリガー
CREATE OR REPLACE FUNCTION create_user_image_on_user_creation()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO user_image (user_id)
    VALUES (NEW.id);
    RETURN NEW;
END;
$$ language 'plpgsql';

-- トリガーの作成
CREATE TRIGGER create_user_image_after_user_creation
    AFTER INSERT ON users
    FOR EACH ROW
    EXECUTE FUNCTION create_user_image_on_user_creation();


-- tag master table
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    tag_name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_tag_name ON tags(tag_name);

CREATE TABLE IF NOT EXISTS user_tags (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    UNIQUE (user_id, tag_id), -- UNIQUE KEYではなくUNIQUEを使用
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
-- 複合インデックス
CREATE INDEX idx_user_tag ON user_tags(user_id, tag_id);


-- PostGISの有効化（必須）

CREATE TABLE IF NOT EXISTS user_location (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    location GEOGRAPHY(POINT, 4326) DEFAULT ST_SetSRID(ST_MakePoint(139.7454, 35.6586), 4326)::geography, -- WGS84形式の位置情報
    --location_alternative GEOGRAPHY(POINT, 4326), --is_gpsがfalseの場合に使う位置情報、ユーザーが自分で設定する
    location_alternative GEOGRAPHY(POINT, 4326) DEFAULT ST_SetSRID(ST_MakePoint(139.7454, 35.6586), 4326)::geography, -- Defaultは東京タワー
    --Alternative
    location_updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- 位置情報の最終更新時刻
    is_gps BOOLEAN DEFAULT FALSE, -- 位置情報を利用するか
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_location UNIQUE (user_id)
);

-- 空間インデックス
CREATE INDEX idx_user_location_gist ON user_location USING GIST (location);
CREATE INDEX idx_user_location_alternative_gist ON user_location USING GIST (location_alternative);

-- 更新時刻の自動更新トリガー
CREATE OR REPLACE FUNCTION update_user_location_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    NEW.location_updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_location_updated_at
    BEFORE UPDATE ON user_location
    FOR EACH ROW
    EXECUTE FUNCTION update_user_location_updated_at_column();

-- ユーザー作成時の自動レコード作成トリガー
CREATE OR REPLACE FUNCTION create_user_location_on_user_creation()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO user_location (user_id)
    VALUES (NEW.id);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER create_user_location_after_user_creation
    AFTER INSERT ON users
    FOR EACH ROW
    EXECUTE FUNCTION create_user_location_on_user_creation();



-- プロフィール閲覧履歴テーブル
CREATE TABLE IF NOT EXISTS profile_views (
    id SERIAL PRIMARY KEY,
    viewer_id INTEGER NOT NULL, -- 閲覧したユーザーのID
    viewed_id INTEGER NOT NULL, -- 閲覧されたユーザーのID
    viewed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (viewer_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (viewed_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT unique_profile_view UNIQUE (viewer_id, viewed_id)
);
-- インデックス
CREATE INDEX idx_profile_views_viewer ON profile_views(viewer_id);
CREATE INDEX idx_profile_views_viewed ON profile_views(viewed_id);
CREATE INDEX idx_profile_views_timestamp ON profile_views(viewed_at);

-- ユーザー間のlike関係を管理するテーブル
CREATE TABLE IF NOT EXISTS user_likes (
    id SERIAL PRIMARY KEY,
    liker_id INTEGER NOT NULL, -- likeを送信したユーザーのID
    liked_id INTEGER NOT NULL, -- likeを受け取ったユーザーのID
    --is_active BOOLEAN DEFAULT TRUE, -- likeが有効か（unlikeした場合はfalse）
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (liker_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (liked_id) REFERENCES users(id) ON DELETE CASCADE,
    -- 同じユーザーの組み合わせで複数のlikeを防ぐ
    CONSTRAINT unique_user_like UNIQUE (liker_id, liked_id)
);


-- ユーザーのfriend関係を管理する
CREATE TABLE IF NOT EXISTS user_friends (
    id SERIAL PRIMARY KEY,
    -- user_idの小さい方をuser_id1、大きい方をuser_id2として格納
    user_id1 INTEGER NOT NULL,
    user_id2 INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id1) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id2) REFERENCES users(id) ON DELETE CASCADE,
    -- user_id1 は必ず user_id2 より小さくなるように制約
    CONSTRAINT user_friend_order CHECK (user_id1 < user_id2),
    -- 同じユーザーの組み合わせを防ぐ
    CONSTRAINT unique_user_friend UNIQUE (user_id1, user_id2)
);
-- インデックス
CREATE INDEX idx_user_friends_user1 ON user_friends(user_id1);
CREATE INDEX idx_user_friends_user2 ON user_friends(user_id2);
-- フレンド関係を作成する関数（IDの大小関係を自動的に処理）
CREATE OR REPLACE FUNCTION create_friend_relationship(user_a INTEGER, user_b INTEGER)
RETURNS VOID AS $$
BEGIN
    IF user_a <> user_b THEN
        INSERT INTO user_friends (user_id1, user_id2)
        VALUES (
            LEAST(user_a, user_b),
            GREATEST(user_a, user_b)
        )
        ON CONFLICT (user_id1, user_id2) DO NOTHING;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- ユーザーのblock関係を管理する
CREATE TABLE IF NOT EXISTS user_blocks (
    id SERIAL PRIMARY KEY,
    blocker_id INTEGER NOT NULL, -- blockしたuser
    blocked_id INTEGER NOT NULL, -- blockされたuser
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (blocker_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (blocked_id) REFERENCES users(id) ON DELETE CASCADE,
    -- 同じユーザーの組み合わせを防ぐ
    CONSTRAINT unique_user_block UNIQUE (blocker_id, blocked_id)
);

CREATE INDEX idx_user_blocks_blocker_id ON user_blocks(blocker_id);

CREATE INDEX idx_user_blocks_blocked_id ON user_blocks(blocked_id);

-- 偽アカウントの報告を管理する
CREATE TABLE IF NOT EXISTS report_fake_accounts (
    id SERIAL PRIMARY KEY,
    reporter_id INTEGER NOT NULL, -- 報告したuser
    fake_account_id INTEGER NOT NULL, -- 報告されたuser
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (reporter_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (fake_account_id) REFERENCES users(id) ON DELETE CASCADE,
    -- 同じユーザーの組み合わせを防ぐ
    CONSTRAINT unique_user_fake UNIQUE (reporter_id, fake_account_id)
);
CREATE INDEX idx_fake_accounts_reporter_id ON report_fake_accounts(reporter_id);
CREATE INDEX idx_fake_accounts_fake_id ON report_fake_accounts(fake_account_id);










-- 都道府県の位置情報を返す関数
CREATE OR REPLACE FUNCTION get_prefecture_location(p prefecture) RETURNS geography AS $$
BEGIN
    RETURN CASE p
        WHEN 'Hokkaido' THEN ST_SetSRID(ST_MakePoint(141.3544, 43.0621), 4326)::geography  -- 札幌
        WHEN 'Aomori' THEN ST_SetSRID(ST_MakePoint(140.7475, 40.8244), 4326)::geography    -- 青森
        WHEN 'Iwate' THEN ST_SetSRID(ST_MakePoint(141.1527, 39.7036), 4326)::geography     -- 盛岡
        WHEN 'Miyagi' THEN ST_SetSRID(ST_MakePoint(140.8694, 38.2688), 4326)::geography    -- 仙台
        WHEN 'Akita' THEN ST_SetSRID(ST_MakePoint(140.1025, 39.7186), 4326)::geography     -- 秋田
        WHEN 'Yamagata' THEN ST_SetSRID(ST_MakePoint(140.3636, 38.2404), 4326)::geography  -- 山形
        WHEN 'Fukushima' THEN ST_SetSRID(ST_MakePoint(140.4748, 37.7500), 4326)::geography -- 福島
        WHEN 'Ibaraki' THEN ST_SetSRID(ST_MakePoint(140.4467, 36.3414), 4326)::geography   -- 水戸
        WHEN 'Tochigi' THEN ST_SetSRID(ST_MakePoint(139.8836, 36.5662), 4326)::geography   -- 宇都宮
        WHEN 'Gunma' THEN ST_SetSRID(ST_MakePoint(139.0600, 36.3912), 4326)::geography     -- 前橋
        WHEN 'Saitama' THEN ST_SetSRID(ST_MakePoint(139.6489, 35.8569), 4326)::geography   -- さいたま
        WHEN 'Chiba' THEN ST_SetSRID(ST_MakePoint(140.1233, 35.6047), 4326)::geography     -- 千葉
        WHEN 'Tokyo' THEN ST_SetSRID(ST_MakePoint(139.7454, 35.6586), 4326)::geography     -- 東京
        WHEN 'Kanagawa' THEN ST_SetSRID(ST_MakePoint(139.6425, 35.4478), 4326)::geography  -- 横浜
        WHEN 'Niigata' THEN ST_SetSRID(ST_MakePoint(139.0364, 37.9022), 4326)::geography   -- 新潟
        WHEN 'Toyama' THEN ST_SetSRID(ST_MakePoint(137.2137, 36.6951), 4326)::geography    -- 富山
        WHEN 'Ishikawa' THEN ST_SetSRID(ST_MakePoint(136.6256, 36.5946), 4326)::geography  -- 金沢
        WHEN 'Fukui' THEN ST_SetSRID(ST_MakePoint(136.2196, 36.0652), 4326)::geography     -- 福井
        WHEN 'Yamanashi' THEN ST_SetSRID(ST_MakePoint(138.5683, 35.6645), 4326)::geography -- 甲府
        WHEN 'Nagano' THEN ST_SetSRID(ST_MakePoint(138.1942, 36.6513), 4326)::geography    -- 長野
        WHEN 'Gifu' THEN ST_SetSRID(ST_MakePoint(136.7223, 35.3912), 4326)::geography      -- 岐阜
        WHEN 'Shizuoka' THEN ST_SetSRID(ST_MakePoint(138.3895, 34.9769), 4326)::geography  -- 静岡
        WHEN 'Aichi' THEN ST_SetSRID(ST_MakePoint(136.9066, 35.1802), 4326)::geography     -- 名古屋
        WHEN 'Mie' THEN ST_SetSRID(ST_MakePoint(136.5088, 34.7303), 4326)::geography       -- 津
        WHEN 'Shiga' THEN ST_SetSRID(ST_MakePoint(135.8682, 35.0045), 4326)::geography     -- 大津
        WHEN 'Kyoto' THEN ST_SetSRID(ST_MakePoint(135.7556, 35.0211), 4326)::geography     -- 京都
        WHEN 'Osaka' THEN ST_SetSRID(ST_MakePoint(135.5023, 34.6937), 4326)::geography     -- 大阪
        WHEN 'Hyogo' THEN ST_SetSRID(ST_MakePoint(135.1955, 34.6913), 4326)::geography     -- 神戸
        WHEN 'Nara' THEN ST_SetSRID(ST_MakePoint(135.8328, 34.6851), 4326)::geography      -- 奈良
        WHEN 'Wakayama' THEN ST_SetSRID(ST_MakePoint(135.1675, 34.2261), 4326)::geography  -- 和歌山
        WHEN 'Tottori' THEN ST_SetSRID(ST_MakePoint(134.2383, 35.5039), 4326)::geography   -- 鳥取
        WHEN 'Shimane' THEN ST_SetSRID(ST_MakePoint(133.0505, 35.4722), 4326)::geography   -- 松江
        WHEN 'Okayama' THEN ST_SetSRID(ST_MakePoint(133.9350, 34.6617), 4326)::geography   -- 岡山
        WHEN 'Hiroshima' THEN ST_SetSRID(ST_MakePoint(132.4590, 34.3963), 4326)::geography -- 広島
        WHEN 'Yamaguchi' THEN ST_SetSRID(ST_MakePoint(131.4714, 34.1859), 4326)::geography -- 山口
        WHEN 'Tokushima' THEN ST_SetSRID(ST_MakePoint(134.5594, 34.0657), 4326)::geography -- 徳島
        WHEN 'Kagawa' THEN ST_SetSRID(ST_MakePoint(134.0436, 34.3428), 4326)::geography    -- 高松
        WHEN 'Ehime' THEN ST_SetSRID(ST_MakePoint(132.7658, 33.8416), 4326)::geography     -- 松山
        WHEN 'Kochi' THEN ST_SetSRID(ST_MakePoint(133.5311, 33.5597), 4326)::geography     -- 高知
        WHEN 'Fukuoka' THEN ST_SetSRID(ST_MakePoint(130.4017, 33.6064), 4326)::geography   -- 福岡
        WHEN 'Saga' THEN ST_SetSRID(ST_MakePoint(130.2988, 33.2494), 4326)::geography      -- 佐賀
        WHEN 'Nagasaki' THEN ST_SetSRID(ST_MakePoint(129.8737, 32.7503), 4326)::geography  -- 長崎
        WHEN 'Kumamoto' THEN ST_SetSRID(ST_MakePoint(130.7417, 32.7898), 4326)::geography  -- 熊本
        WHEN 'Oita' THEN ST_SetSRID(ST_MakePoint(131.6088, 33.2382), 4326)::geography      -- 大分
        WHEN 'Miyazaki' THEN ST_SetSRID(ST_MakePoint(131.4202, 31.9111), 4326)::geography  -- 宮崎
        WHEN 'Kagoshima' THEN ST_SetSRID(ST_MakePoint(130.5584, 31.5603), 4326)::geography -- 鹿児島
        WHEN 'Okinawa' THEN ST_SetSRID(ST_MakePoint(127.6809, 26.2124), 4326)::geography   -- 那覇
    END;
END;
$$ LANGUAGE plpgsql;


-- テストユーザー生成の関数定義
CREATE OR REPLACE FUNCTION generate_test_users() RETURNS void AS $$
DECLARE
    created_user_id integer;
    i integer;
    username text;
    email text;
    gender_val gender_type;
    sexuality_val sexuality_type;
    area_val prefecture;
    birth_date date;
    image_path text;
    genders gender_type[] := ARRAY['male'::gender_type, 'female'::gender_type];
    male_names text[] := ARRAY['takashi', 'yutaka', 'koji', 'kai', 'shin'];
    female_names text[] := ARRAY['yui', 'mayumi', 'miyu', 'meiko', 'keiko'];
    sexualities sexuality_type[] := ARRAY['male'::sexuality_type, 'female'::sexuality_type, 'male/female'::sexuality_type];
BEGIN
    FOREACH gender_val IN ARRAY genders LOOP
        FOR i IN 1..2500 LOOP
            -- ユーザー名の生成（性別に応じて）
            IF gender_val = 'male' THEN
                username := male_names[1 + (i % 5)] || i::text;
            ELSE
                username := female_names[1 + (i % 5)] || i::text;
            END IF;

            -- 基本情報の設定
            email := username || i::text || '@ft.com';
            birth_date := '1964-01-01'::date + (random() * ('2006-12-31'::date - '1964-01-01'::date))::integer;
            sexuality_val := sexualities[1 + (floor(random() * 3))::integer];
            area_val := (SELECT enumlabel::prefecture FROM pg_enum WHERE enumtypid = 'prefecture'::regtype ORDER BY random() LIMIT 1);

            -- ユーザーの作成
            INSERT INTO users (
                username,
                email,
                password_hash,
                is_registered,
                is_preparation
            ) VALUES (
                username,
                email,
                '$2a$10$bXCkPbJHvuHQ3UqzNvH9.OYdBhCGVVxAyCPO1rGgE6LL6RbqQHnhK',
                true,
                true
            ) RETURNING id INTO created_user_id;

            -- プロフィール情報の更新
            UPDATE user_info SET
                lastname = 'Kusano',
                firstname = 'Yoshinari',
                birthdate = birth_date,
                gender = gender_val,
                sexuality = sexuality_val,
                area = area_val,
                self_intro = 'Nice to meet you, my name is Satoshi Nakamoto!'
            WHERE user_id = created_user_id;

            -- 選択された都道府県の県庁所在地の位置情報を設定
            UPDATE user_location SET
                location_alternative = get_prefecture_location(area_val)
            WHERE user_id = created_user_id;

            -- プロフィール画像の設定
            image_path := '/home/appuser/uploads/images1/' || 
                         gender_val || 
                         (1 + (random() * 4)::integer)::text || 
                         '.png';
            UPDATE user_image SET
                profile_image_path1 = image_path
            WHERE user_id = created_user_id;
        END LOOP;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- 初期データ投入の実行
DO $$
BEGIN
    -- ユーザーが存在しない場合のみ実行
    IF NOT EXISTS (SELECT 1 FROM users LIMIT 1) THEN
        RAISE NOTICE 'Generating test users...';
        PERFORM generate_test_users();
        RAISE NOTICE 'Test users generation completed.';
    END IF;
END $$;