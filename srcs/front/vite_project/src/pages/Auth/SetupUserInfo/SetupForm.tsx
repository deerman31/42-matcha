import { useNavigate } from "npm:react-router-dom";
import { ChangeEvent, FormEvent, useState } from "react";
import { ErrorResponse, Response } from "../../../types/api.ts";

import DatePicker from "../../../components/DatePicker.tsx";

import GenderPicker from "../../../components/GenderPicker.tsx";

import SexualityPicker from "../../../components/SexualityPicker.tsx";

import AreaPicker from "../../../components/AreaPicker.tsx";

import ImageUploader from "../../../components/ImageUploader.tsx";

import { getToken } from "../../../utils/auth.ts";

interface FormData {
  lastname: string;
  firstname: string;
  birthdate: string;
  gender: string;
  sexuality: string;
  area: string;
  self_intro: string;
}

// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

const SetupForm = () => {
  const navigate = useNavigate();

  // フォームの初期状態
  const initialFormState: FormData = {
    lastname: "",
    firstname: "",
    birthdate: "2000-01-01",
    gender: "male",
    sexuality: "male",
    area: "Hokkaido",
    self_intro: "",
  };
  // 状態管理
  const [formData, setFormData] = useState<FormData>(initialFormState);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [submitStatus, setSubmitStatus] = useState<SubmitStatus>({
    type: "",
    message: "",
  });
  // 画像ファイルの状態
  const [imageFile, setImageFile] = useState<File | null>(null);

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

  // DatePicker コンポーネントに渡すハンドラーを別途定義
  const handleDateChange = (date: string): void => {
    setFormData((prev: FormData) => ({
      ...prev,
      birthdate: date,
    }));
  };

  const handleAreaChange = (area: string): void => {
    setFormData((prev: FormData) => ({
      ...prev,
      area: area,
    }));
  };

  // 画像選択時の処理
  const handleImageSelect = (imageFormData: globalThis.FormData): void => {
    const file = imageFormData.get("image") as File;
    if (file) {
      setImageFile(file);
    }
  };

  // フォーム送信ハンドラー
  const handleSubmit = async (e: FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setIsSubmitting(true);
    setSubmitStatus({ type: "", message: "" });

    if (!imageFile) {
      setSubmitStatus({
        type: "error",
        message: "画像を選択してください",
      });
      setIsSubmitting(false);
      return;
    }

    try {
      // FormDataオブジェクトの作成
      const submitFormData = new FormData();

      // 画像ファイルの追加を最初に行う
      submitFormData.append("image", imageFile);

      // その他のフォームデータの追加
      submitFormData.append("lastname", formData.lastname);
      submitFormData.append("firstname", formData.firstname);
      submitFormData.append("birthdate", formData.birthdate);
      submitFormData.append("gender", formData.gender);
      submitFormData.append("sexuality", formData.sexuality);
      submitFormData.append("area", formData.area);
      submitFormData.append("self_intro", formData.self_intro);

      const token = getToken();

      const response = await fetch("/api/users/set/user-info", {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${token}`,
        },
        body: submitFormData, // Content-Typeヘッダーは自動で設定される
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

      // 成功メッセージを表示した後、短いディレイを設けてからリダイレクト

      setTimeout(() => {
        navigate("/my-profile", {
          // 必要に応じて状態を渡すことができます
          state: { from: "registration", message: data.message },
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
      <h2 className="form-title">Init</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="lastname" className="form-label">
            LastName
          </label>
          <input
            type="text"
            id="lastname"
            name="lastname"
            value={formData.lastname}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>

        <div className="form-group">
          <label htmlFor="firstname" className="form-label">
            FirstName
          </label>
          <input
            type="text"
            id="firstname"
            name="firstname"
            value={formData.firstname}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>

        <div>
          <label htmlFor="birthdate">
            BirthDate
          </label>
          <DatePicker
            value={formData.birthdate}
            onChange={handleDateChange} // 専用のハンドラーを使用
          />
        </div>

        <div>
          <label htmlFor="gender">
            Gender
          </label>
          <GenderPicker value={formData.gender} onChange={handleChange} />
        </div>

        <div>
          <label htmlFor="sexuality">
            Sexuality
          </label>
          <SexualityPicker
            value={formData.sexuality}
            onChange={handleChange}
          />
        </div>

        <div>
          <label htmlFor="area">
            Area
          </label>
          <AreaPicker value={formData.area} onChange={handleAreaChange} />
        </div>

        <div className="form-group">
          <label htmlFor="self_intro" className="form-label">
            SelfIntro
          </label>
          <input
            type="self_intro"
            id="self_intro"
            name="self_intro"
            value={formData.self_intro}
            onChange={handleChange}
            required
            className="form-input"
          />
        </div>

        <div>
          <label htmlFor="image">
            Image
          </label>
          <ImageUploader
            onImageSelect={handleImageSelect}
            maxSizeMB={5}
            acceptedTypes={["image/jpeg", "image/png", "image/gif"]}
          />
          {imageFile && (
            <p className="text-sm text-green-600">
              選択された画像: {imageFile.name}
            </p>
          )}
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

export default SetupForm;
