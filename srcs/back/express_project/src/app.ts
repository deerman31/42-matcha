// src/app.ts
import express from "express";
import { Database } from "./db";
import dotenv from "dotenv";
import { setupRoutes } from "./routes";

// 環境変数を読み込む
dotenv.config();

const app = express();
const port = 3000;

// データベース接続の初期化
let db: Database;


function routing() {
  app.get("/", (req, res) => {
    res.send("Hello World!");
  });
}

async function initializeApp() {
  try {
    db = new Database();

    // 接続テスト
    await db.testConnection();
    console.log("Database connection successful");

    // ルーティングのセットアップ
    app.use('/api', setupRoutes(db));

    app.listen(port, () => {
      console.log(`Example app listening on port ${port}`);
    });

  } catch (error) {
    console.error("Failed to initialize application:", error);
    process.exit(1);
  }
}

initializeApp();

// Graceful shutdown
process.on('SIGTERM', async () => {
  console.log('SIGTERM received. Closing database connections...');
  if (db) {
    await db.end();
  }
  process.exit(0);
});
