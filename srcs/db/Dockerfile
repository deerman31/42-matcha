FROM postgres:15

# システムパッケージの更新とPostGISのインストール
RUN apt-get update \
    && apt-get install -y postgis postgresql-15-postgis-3 \
    && rm -rf /var/lib/apt/lists/*

# init.sql をコンテナの初期化スクリプトディレクトリにコピー
COPY init.sql /docker-entrypoint-initdb.d/

CMD ["postgres"]