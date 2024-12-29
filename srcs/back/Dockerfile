# ベースイメージ
FROM node:20-slim

# 作業ディレクトリを設定
WORKDIR /usr/src/app

# パッケージ関連ファイルをコピー
COPY package*.json ./

# 依存関係のインストール
RUN npm install

# TypeScriptの設定ファイルをコピー
COPY tsconfig.json ./

# ソースコードをコピー
COPY src/ ./src/

# TypeScriptをコンパイル
RUN npm run build

# ポート設定
EXPOSE ${BACK_PORT}

# アプリケーションの起動
CMD ["npm", "start"]