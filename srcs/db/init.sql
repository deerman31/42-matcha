CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,  -- AUTO_INCREMENT の代わりに SERIAL
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    is_online BOOLEAN DEFAULT FALSE,
    is_registered BOOLEAN DEFAULT FALSE
);

CREATE INDEX idx_username ON users(username);
CREATE INDEX idx_email ON users(email);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_info (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    firstname VARCHAR(50) NOT NULL,
    birthdate DATE NOT NULL,
    is_gps BOOLEAN DEFAULT FALSE,
    gender VARCHAR(6) CHECK (gender IN ('male', 'female')) DEFAULT 'male',  -- ENUMの代わりにCHECK制約
    sexual_orientation VARCHAR(12) CHECK (sexual_orientation IN ('heterosexual', 'homosexual', 'bisexual')) DEFAULT 'bisexual',
    eria VARCHAR(10) CHECK (eria IN (
        'Hokkaido', 'Aomori', 'Iwate', 'Miyagi', 'Akita', 'Yamagata', 'Fukushima',
        'Ibaraki', 'Tochigi', 'Gunma', 'Saitama', 'Chiba', 'Tokyo', 'Kanagawa',
        'Niigata', 'Toyama', 'Ishikawa', 'Fukui', 'Yamanashi', 'Nagano',
        'Gifu', 'Shizuoka', 'Aichi', 'Mie',
        'Shiga', 'Kyoto', 'Osaka', 'Hyogo', 'Nara', 'Wakayama',
        'Tottori', 'Shimane', 'Okayama', 'Hiroshima', 'Yamaguchi',
        'Tokushima', 'Kagawa', 'Ehime', 'Kochi',
        'Fukuoka', 'Saga', 'Nagasaki', 'Kumamoto', 'Oita', 'Miyazaki', 'Kagoshima', 'Okinawa'
    )) DEFAULT 'Tokyo',
    self_intro VARCHAR(300) NOT NULL DEFAULT '',
    profile_image_path1 VARCHAR(255) DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);