// app/routes/signup.tsx

import type { ActionFunction } from "@remix-run/node";
import { Form, useActionData, useNavigation } from "@remix-run/react";
import { useState } from "react";
import { redirect } from "@remix-run/node"; // redirectをインポート
import { validateEmail, validateUsername, validatePassword, validateRepassword } from "~/utils/validation";


interface ActionData {
  formError?: string;
  fieldErrors?: {
    username?: string;
    email?: string;
    password?: string;
    repassword?: string;
  };
  fields?: {
    username: string;
    email: string;
  };
  backendError?: string;
  success?: boolean;
}

interface FormData {
  username: string;
  email: string;
  password: string;
  repassword: string;
}

export const action: ActionFunction = async ({ request }) => {
  const form = await request.formData();
  const username = form.get("username") as string | null;
  const email = form.get("email") as string | null;
  const password = form.get("password") as string | null;
  const repassword = form.get("repassword") as string | null;

  if (!username || !email || !password || !repassword) {
    return new Response(
      JSON.stringify({ formError: "すべての項目を入力してください。" }),
      {
        status: 400,
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
  }

  const fields = { username, email };
  const fieldErrors = {
    username: validateUsername(username),
    email: validateEmail(email),
    password: validatePassword(password),
    repassword: validateRepassword(password, repassword),
  };

  if (Object.values(fieldErrors).some(Boolean)) {
    return new Response(
      JSON.stringify({ fieldErrors, fields }),
      {
        status: 400,
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
  }

  const formData: FormData = {
    username,
    email,
    password,
    repassword,
  };

  try {
    // サーバーサイドからバックエンドへの直接的なリクエスト
    const response = await fetch("http://back:3000/api/register", {
    //const response = await fetch("/api/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });

    // バックエンドからのレスポンスを取得
    const responseData = await response.json();

    if (!response.ok) {
      return new Response(
        JSON.stringify({
          backendError: responseData.error || "登録中にエラーが発生しました。",
          fields,
        }),
        {
          status: response.status,
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
    }
    // 成功時は/loginにリダイレクト
    return redirect("/login");
  } catch (error) {
    return new Response(
      JSON.stringify({
        formError: "サーバーとの通信中にエラーが発生しました。",
        fields,
      }),
      {
        status: 500,
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
  }
};

export default function Signup() {
  const actionData = useActionData<ActionData>();
  const navigation = useNavigation();
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const isSubmitting = navigation.state === "submitting";

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          アカウント作成
        </h2>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          <Form
            method="post"
            className="space-y-6">
            <div>
              <label
                htmlFor="username"
                className="block text-sm font-medium text-gray-700"
              >
                ユーザー名
              </label>
              <div className="mt-1">
                <input
                  id="username"
                  name="username"
                  type="text"
                  required
                  defaultValue={actionData?.fields?.username}
                  className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                {actionData?.fieldErrors?.username && (
                  <p className="mt-2 text-sm text-red-600">
                    {actionData.fieldErrors.username}
                  </p>
                )}
              </div>
            </div>

            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                メールアドレス
              </label>
              <div className="mt-1">
                <input
                  id="email"
                  name="email"
                  type="email"
                  autoComplete="email"
                  required
                  defaultValue={actionData?.fields?.email}
                  className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                {actionData?.fieldErrors?.email && (
                  <p className="mt-2 text-sm text-red-600">
                    {actionData.fieldErrors.email}
                  </p>
                )}
              </div>
            </div>

            <div>
              <label
                htmlFor="password"
                className="block text-sm font-medium text-gray-700"
              >
                パスワード
              </label>
              <div className="mt-1 relative">
                <input
                  id="password"
                  name="password"
                  type={showPassword ? "text" : "password"}
                  required
                  className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute inset-y-0 right-0 pr-3 flex items-center text-sm leading-5"
                >
                  {showPassword ? "非表示" : "表示"}
                </button>
                {actionData?.fieldErrors?.password && (
                  <p className="mt-2 text-sm text-red-600">
                    {actionData.fieldErrors.password}
                  </p>
                )}
              </div>
            </div>

            <div>
              <label
                htmlFor="repassword"
                className="block text-sm font-medium text-gray-700"
              >
                パスワード(確認)
              </label>
              <div className="mt-1">
                <input
                  id="repassword"
                  name="repassword"
                  type={showPassword ? "text" : "password"}
                  required
                  className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                />

                {/* バックエンドエラーの表示 */}
                {actionData?.backendError && (
                  <div className="rounded-md bg-red-50 p-4">
                    <div className="flex">
                      <div className="ml-3">
                        <p className="text-sm font-medium text-red-800">
                          {actionData.backendError}
                        </p>
                      </div>
                    </div>
                  </div>
                )}

                {/* フォームエラーの表示 */}
                {actionData?.formError && (
                  <div className="rounded-md bg-red-50 p-4">
                    <div className="flex">
                      <div className="ml-3">
                        <p className="text-sm font-medium text-red-800">
                          {actionData.formError}
                        </p>
                      </div>
                    </div>
                  </div>
                )}

              </div>
            </div>
            <div>
              <button
                type="submit"
                disabled={isSubmitting}
                className={`w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 ${isSubmitting ? "opacity-50 cursor-not-allowed" : ""
                  }`}
              >
                {isSubmitting ? "登録中..." : "アカウントを作成"}
              </button>
            </div>
          </Form>
        </div>
      </div>
    </div>
  );
}