import { ChangeEvent, FormEvent, useState } from "react";
import { ErrorResponse, Response } from "../../../../types/api.ts";
import { getToken } from "../../../../utils/auth.ts";

import "./ChangeSelfIntro.css";

interface FormData {
  self_intro: string;
}

// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

const ChangeSelfIntro = () => {
  // フォームの初期状態
  const initialFormState: FormData = {
    self_intro: "",
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
      const token = getToken();

      const response = await fetch("/api/users/set/self-intro", {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({ value: formData.self_intro }),
      });

      if (!response.ok) {
        const errorData: ErrorResponse = await response.json();
        throw new Error(errorData.error || "送信に失敗しました");
      }

      const data: Response = await response.json();

      setSubmitStatus({
        type: "success",
        message: data.message || "登録が完了しました！",
      });

      // フォームをリセット
      setFormData(initialFormState);
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
      <h2 className="form-title">Change SelfIntro</h2>

      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="self_intro" className="form-label">
            SelfIntro
          </label>
          <input
            type="text"
            id="self_intro"
            name="self_intro"
            value={formData.self_intro}
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
          {isSubmitting ? "更新中..." : "更新"}
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

export default ChangeSelfIntro;
