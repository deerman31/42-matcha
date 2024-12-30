FROM node:23-bullseye-slim

ENV PORT=${FRONT_PORT}

WORKDIR /app

COPY ./vite-react/package*.json ./

# 依存関係のインストール
RUN npm install

# ソースコードをコピー
COPY ./vite-react/ .

RUN npm run build

# Expose the new port the app runs on
#EXPOSE 5173
EXPOSE ${PORT}

# Define the command to run the app
#CMD ["npm", "start"]
CMD ["npm", "run", "dev"]