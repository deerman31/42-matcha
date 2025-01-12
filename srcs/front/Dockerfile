FROM node:23-bullseye-slim

WORKDIR /app


COPY ./remix-sample/package*.json ./

# 依存関係のインストール
RUN npm install

# ソースコードをコピー
COPY ./remix-sample/ .

RUN npm run build

# Expose the new port the app runs on
EXPOSE 5173

# Define the command to run the app
CMD ["npm", "run", "dev"]
