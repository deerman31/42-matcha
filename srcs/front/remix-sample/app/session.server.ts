// app/session.server.ts
import { createCookieSessionStorage } from "@remix-run/node";

export const sessionStorage = createCookieSessionStorage({
    cookie: {
        name: "__session", // クッキーの名前
        httpOnly: true,    // JavaScriptからアクセス不可
        path: "/",         // クッキーのパス
        sameSite: "lax",   // CSRF保護
        secrets: ["s3cr3t"], // 本番環境では環境変数から取得すべき
        secure: process.env.NODE_ENV === "production", // HTTPS環境では true
    },
});

// セッション取得のヘルパー関数
export async function getSession(request: Request) {
    const cookie = request.headers.get("Cookie");
    return sessionStorage.getSession(cookie);
}

// トークンの保存と取得のヘルパー関数
export async function setAccessToken(request: Request, token: string) {
    const session = await getSession(request);
    session.set("accessToken", token);
    return sessionStorage.commitSession(session);
}

export async function getAccessToken(request: Request) {
    const session = await getSession(request);
    return session.get("accessToken");
}

// トークンを削除するヘルパー関数を追加
export async function destroyAccessToken(request: Request) {
    const session = await getSession(request);
    session.unset("accessToken"); // トークンを削除
    return sessionStorage.commitSession(session);
}

// もしくは、セッション全体を破棄する場合は以下の関数を使用
export async function destroySession(request: Request) {
    const session = await getSession(request);
    return sessionStorage.destroySession(session);
}