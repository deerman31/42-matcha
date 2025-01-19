# 公式のNginxイメージをベースとして使用
FROM nginx:alpine

# タイムゾーンを日本時間に設定
ENV TZ=Asia/Tokyo

# 必要なツールをインストール
RUN apk add --no-cache tzdata openssl

# Nginxの設定ファイルをコピー
COPY ./tool/nginx.conf /etc/nginx/nginx.conf

COPY ./tool/generate-ssl.sh /generate-ssl.sh

# 実行権限を付与
RUN chmod +x /generate-ssl.sh

# デフォルトのサーバーブロックを削除
RUN rm /etc/nginx/conf.d/default.conf

# 静的ファイル用のディレクトリを作成
RUN mkdir -p /usr/share/nginx/html

# Nginxユーザーに所有権を変更
RUN chown -R nginx:nginx /usr/share/nginx/html && \
    chmod -R 755 /usr/share/nginx/html

# Nginxのデフォルトポートを公開
EXPOSE 443

# ヘルスチェック用のファイルを作成
RUN echo "OK" > /usr/share/nginx/html/health.txt

# 起動時にSSL証明書を生成し、Nginxを起動
CMD ["/bin/sh", "-c", "/generate-ssl.sh && nginx -g 'daemon off;'"]