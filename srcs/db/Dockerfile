FROM postgres:15

# init.sql をコンテナの初期化スクリプトディレクトリにコピー
COPY init.sql /docker-entrypoint-initdb.d/

CMD ["postgres"]