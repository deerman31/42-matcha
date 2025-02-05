import {
  ChangeEvent,
  FormEvent,
  useCallback,
  useEffect,
  useState,
} from "react";
import RangeSlider from "../slider/RangeSlider/RangeSlider.tsx";
import SingleSlider from "../slider/SingleSlider/SingleSlider.tsx";

import SortOptionPicker from "./SortOptionPicker/SortOptionPicker.tsx";
import SortOrderPicker from "./SortOrderPicker/SortOrderPicker.tsx";
import { getToken } from "../../utils/auth.ts";

import { BrowseResponse, ErrorResponse } from "../../types/api.ts";

//age distance fame_rating tag

enum SortOptionType {
  Age = "age",
  Distance = "distance",
  FameRating = "fame_rating",
  Tag = "tag",
}

enum SortOrder {
  Descending = 0,
  Ascending,
}

interface FormData {
  age_range: {
    min: number;
    max: number;
  };
  distance_range: {
    min: number;
    max: number;
  };
  min_common_tags: number;
  min_fame_rating: number;
  sort_option: SortOptionType;
  sort_order: SortOrder;
}
// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

const BrowseConditions = () => {
  // フォームの初期状態
  const initialFormState: FormData = {
    age_range: {
      min: 20,
      max: 60,
    },
    distance_range: {
      min: 1,
      max: 100,
    },
    min_common_tags: 0,
    min_fame_rating: 0,
    sort_option: SortOptionType.Age,
    sort_order: SortOrder.Descending,
  };

  // 状態管理
  const [formData, setFormData] = useState<FormData>(initialFormState);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [submitStatus, setSubmitStatus] = useState<SubmitStatus>({
    type: "",
    message: "",
  });

  // 年齢範囲の更新ハンドラー
  // 各ハンドラーはuseCallbackでメモ化
  const handleAgeRangeChange = useCallback((value: [number, number]) => {
    setFormData((prev: FormData) => ({
      ...prev,
      age_range: {
        min: value[0],
        max: value[1],
      },
    }));
  }, []);

  // 距離範囲の更新ハンドラー
  const handleDistanceRangeChange = useCallback((value: [number, number]) => {
    setFormData((prev: FormData) => ({
      ...prev,
      distance_range: {
        min: value[0],
        max: value[1],
      },
    }));
  }, []);

  // 共通タグ数の更新ハンドラー
  const handleCommonTagsChange = useCallback((value: number) => {
    setFormData((prev: FormData) => ({
      ...prev,
      min_common_tags: value,
    }));
  }, []);

  // 評価の更新ハンドラー
  const handleFameRatingChange = useCallback((value: number) => {
    setFormData((prev: FormData) => ({
      ...prev,
      min_fame_rating: value,
    }));
  }, []);

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

  const handleSubmit = async (e: FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setIsSubmitting(true);
    setSubmitStatus({ type: "", message: "" });

    try {
      const token = getToken();

      const response = await fetch("/api/browse", {
        method: "POST",
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
      const data: BrowseResponse = await response.json();

      console.log(data.user_infos);

      setSubmitStatus({
        type: "success",
        message: "送信が完了しました！",
      });
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

  // useEffect(() => {
  //   console.log('Form data updated:', formData);
  // }, [formData]);

  return (
    <div className="browse-form-container">
      <h2 className="browse-form-title">YKUSANO</h2>

      <form onSubmit={handleSubmit}>
        <div>
          <RangeSlider
            min={20}
            max={60}
            value={[formData.age_range.min, formData.age_range.max]}
            step={1}
            onChange={handleAgeRangeChange}
          />
        </div>

        <div>
          <RangeSlider
            min={1}
            max={100}
            value={[
              formData.distance_range.min,
              formData.distance_range.max,
            ]}
            step={1}
            onChange={handleDistanceRangeChange}
          />
        </div>

        <div>
          <SingleSlider
            min={0}
            max={5}
            value={formData.min_common_tags}
            step={1}
            onChange={handleCommonTagsChange}
          />
        </div>

        <div>
          <SingleSlider
            min={0}
            max={5}
            value={formData.min_fame_rating}
            step={1}
            onChange={handleFameRatingChange}
          />
        </div>

        <div>
          <SortOptionPicker
            value={formData.sort_option}
            onChange={handleChange}
          />
        </div>

        <div>
          <SortOrderPicker
            value={formData.sort_order}
            onChange={handleChange}
          />
        </div>

        <button
          type="submit"
          disabled={isSubmitting}
          className="browse-submit-button"
        >
          {isSubmitting ? "Ykusano中..." : "Ykusano"}
        </button>
      </form>

      {submitStatus.message && (
        <div
          className={`browse-alert ${
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

export default BrowseConditions;
