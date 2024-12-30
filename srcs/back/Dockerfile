FROM node:23-bullseye-slim

WORKDIR /app

COPY ./express_project/package*.json ./

# 依存関係のインストール
RUN npm install

# ソースコードをコピー
COPY ./express_project/ .

RUN npm run build

# Expose the new port the app runs on
EXPOSE 3000

# Define the command to run the app
CMD ["npm", "run", "dev"]