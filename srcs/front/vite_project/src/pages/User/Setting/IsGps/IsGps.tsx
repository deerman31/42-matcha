import { ChangeEvent, FormEvent, useState } from "react";

import { ErrorResponse, Response } from "../../../../types/api.ts";

import { getToken } from "../../../../utils/auth.ts";

import IsGpsPicker from "../IsGps/IsGpsPicker.tsx";

import { useLocationContext } from "../../../../components/LocationService/LocationContextType.tsx";

interface FormData {
  isGpsEnabled: boolean;
}

// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

const ChangeIsGps = () => {
  const {  setTrackingEnabled } = useLocationContext();

  // フォームの初期状態
  const initialFormState: FormData = {
    isGpsEnabled: true,
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
    //   [name]: value,
      //[name]: value === 'true', // 文字列を boolean に変換
      [name]: Boolean(value),  // 単純にBooleanに変換
    }));
  };

  // フォーム送信ハンドラー
  const handleSubmit = async (e: FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setIsSubmitting(true);
    setSubmitStatus({ type: "", message: "" });


    try {
      const token = getToken();

      const response = await fetch("/api/gps/set/is-gps", {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(formData),
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

      setTrackingEnabled(formData.isGpsEnabled);


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
      <h2 className="form-title">Init</h2>
      <form onSubmit={handleSubmit}>


        <div>
          <label htmlFor="isgps">
            IsGps
          </label>
          <IsGpsPicker value={formData.isGpsEnabled} onChange={handleChange} />
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
export default ChangeIsGps;