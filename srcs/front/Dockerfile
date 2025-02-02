# FROM node:23-bullseye-slim
FROM denoland/deno:debian

WORKDIR /app

# パーミッションを設定
ENV DENO_DIR=/deno-dir
ENV DENO_INSTALL_ROOT=/usr/local

# キャッシュを活用するため、依存関係を先にコピー
COPY ./vite_project/deno.json ./vite_project/deno.lock* ./

# 依存関係をキャッシュ
RUN deno cache --reload deps.ts || true

# ソースコードをコピー
COPY ./vite_project/ .

# ビルド実行
RUN deno task build

# Expose the new port the app runs on
EXPOSE 5173


# 開発サーバー起動
CMD ["deno", "task", "dev"]