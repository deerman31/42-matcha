# FROM postgres:latest

# ENV PORT=${POSTGRES_PORT}

# COPY ./init.sql ./docker-entrypoint-initdb.d/init.sql

# EXPOSE ${PORT}

# バージョンを固定した方が安全です
FROM postgres:16  

# PostgreSQLの標準環境変数を設定
ENV POSTGRES_DB=${POSTGRES_DB}
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
# PostgreSQLの標準的なポート環境変数
ENV PGPORT=${POSTGRES_PORT} 

# パスを絶対パスに修正
COPY ./init.sql /docker-entrypoint-initdb.d/  

EXPOSE ${POSTGRES_PORT}