import { useState } from "@types/react";
import RangeSlider from "../slider/RangeSlider.tsx";
import SingleSlider from "../slider/SingleSlider.tsx";

interface FormData {
    age_range: {
        min:number;
        max:number;
    },
    distance_range:{
        min:number;
        max:number;
    },
    min_common_tags:number;
    min_fame_rating:number;
    sort_option:number;
    sort_order:number;
};
// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

const BrowseConditions = () => {
  // フォームの初期状態
  const initialFormState: FormData = {
    age_range:{
        min:20,
        max:60,
    },
    distance_range:{
        min:1,
        max:100,
    },
    min_common_tags:0,
    min_fame_rating:0,
    sort_option:0,
    sort_order:0,
  };

  // 状態管理
  const [formData, setFormData] = useState<FormData>(initialFormState);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [submitStatus, setSubmitStatus] = useState<SubmitStatus>({
    type: "",
    message: "",
  });

  return (
    <>
      <div>
        <RangeSlider
          min={20}
          max={60}
          defaultValue={[20, 37]}
          step={1}
          onChange={(value) => console.log("Value changed:", value)}
        />
      </div>

      <div>
        <RangeSlider
          min={1}
          max={100}
          defaultValue={[1, 100]}
          step={1}
          onChange={(value) => console.log("Value changed:", value)}
        />
      </div>

      <div>
        <SingleSlider
          min={0}
          max={5}
          defaultValue={0}
          step={1}
          onChange={(value) => console.log(value)}
        />
      </div>

      <div>
        <SingleSlider
          min={0}
          max={5}
          defaultValue={0}
          step={1}
          onChange={(value) => console.log(value)}
        />
      </div>

    </>
  );
};

export default BrowseConditions;
