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
    is_gps BOOLEAN DEFAULT FALSE, -- 位置情報を利用するか
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
    location GEOGRAPHY(POINT, 4326), -- WGS84形式の位置情報
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