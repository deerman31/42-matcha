import { ChangeEvent, FormEvent, useState } from "react";
import "./LoginForm.css";

import { ErrorResponse, LoginResponse } from "../../../types/api.ts";

import { saveToken } from "../../../utils/auth.ts";

import { useNavigate } from "npm:react-router-dom";

interface FormData {
  username: string;
  password: string;
}
// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

const LoginForm = () => {
  const navigate = useNavigate();

  // フォームの初期状態
  const initialFormState: FormData = {
    username: "",
    password: "",
  };
  // 状態管理
  const [formData, setFormData] = useState<FormData>(initialFormState);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [submitStatus, setSubmitStatus] = useState<SubmitStatus>({
    type: "",
    message: "",
  });

  // 入力変更ハンドラー
  const handleChange = (
    e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ): void => {
    const { name, value } = e.target;
    setFormData((prev: FormData) => ({
      ...prev,
      [name]: value,
    }));
  };

  // フォーム送信ハンドラー
  const handleSubmit = async (e: FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setIsSubmitting(true);
    setSubmitStatus({ type: "", message: "" });

    try {
      const response = await fetch("/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        const errorData: ErrorResponse = await response.json();
        throw new Error(errorData.error || "送信に失敗しました");
      }
      const data: LoginResponse = await response.json();

      setSubmitStatus({
        type: "success",
        message: "送信が完了しました！",
      });

      // フォームをリセット
      setFormData(initialFormState);

      // tokenをsave
      saveToken(data.access_token);

      const redirectURL = data.is_preparation ? "/home" : "/setup-user-info";

      // 成功メッセージを表示した後、短いディレイを設けてからリダイレクト
      setTimeout(() => {
        navigate(redirectURL, {
          // 必要に応じて状態を渡すことができます
          state: { from: "registration", message: "Login success" },
        });
      }, 1000); // 1.0秒後にリダイレクト
    } catch (error) {
      setSubmitStatus({
        type: "error",
        message: error instanceof Error
          ? error.message
          : "エラーが発生しました。もう一度お試しください。",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="form-container">
      <h2 className="form-title">Login</h2>

      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="username" className="form-label">
            UserName
          </label>
          <input
            type="text"
            id="username"
            name="username"
            value={formData.username}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>

        <div className="form-group">
          <label htmlFor="password" className="form-label">
            Password
          </label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>

        <button
          type="submit"
          disabled={isSubmitting}
          className="submit-button"
        >
          {isSubmitting ? "Login中..." : "Login"}
        </button>
      </form>

      {submitStatus.message && (
        <div
          className={`alert ${
            submitStatus.type === "success" ? "alert-success" : "alert-error"
          }`}
          role="alert"
          aria-live="polite"
        >
          {submitStatus.message}
        </div>
      )}
    </div>
  );
};

export default LoginForm;
