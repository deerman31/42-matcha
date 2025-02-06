import { FormEvent, useState } from "react";
import { ErrorResponse, Response } from "../../../../types/api.ts";

import AreaPicker from "../../../../components/AreaPicker.tsx";

import { getToken } from "../../../../utils/auth.ts";

interface FormData {
  area: string;
}

// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

const ChangeArea = () => {
  // フォームの初期状態
  const initialFormState: FormData = {
    area: "Hokkaido",
  };
  // 状態管理
  const [formData, setFormData] = useState<FormData>(initialFormState);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [submitStatus, setSubmitStatus] = useState<SubmitStatus>({
    type: "",
    message: "",
  });

  // フォーム送信ハンドラー
  const handleSubmit = async (e: FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setIsSubmitting(true);
    setSubmitStatus({ type: "", message: "" });

    try {
      const token = getToken();

      const response = await fetch("/api/users/set/area", {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({ value: formData.area }),
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

  const handleAreaChange = (area: string): void => {
    setFormData((prev: FormData) => ({
      ...prev,
      area: area,
    }));
  };

  return (
    <div className="form-container">
      <h2 className="form-title">Change Area</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="area">
            Area
          </label>
          <AreaPicker
            value={formData.area}
            onChange={handleAreaChange} // 専用のハンドラーを使用
          />
        </div>

        <button
          type="submit"
          disabled={isSubmitting}
          className="submit-button"
        >
          {isSubmitting ? "送信中..." : "送信"}
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

export default ChangeArea;
