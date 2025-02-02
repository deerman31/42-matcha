import React, { ChangeEvent, useEffect, useState } from "react";

// 性別を定義する型
type Gender = "male" | "female";

// コンポーネントのProps型定義
interface GenderPickerProps {
  value: string; // 親コンポーネントに合わせてstring型に
  onChange: (e: ChangeEvent<HTMLInputElement>) => void; // 親コンポーネントの期待する型に合わせる
}

const GenderPicker: React.FC<GenderPickerProps> = ({
  value,
  onChange,
}: GenderPickerProps) => {
  // 選択された性別の状態管理
  const [selectedGender, setSelectedGender] = useState<Gender>(
    value as Gender || "male",
  );

  // valueの変更を監視して内部状態を更新
  useEffect(() => {
    setSelectedGender(value as Gender);
  }, [value]);

  // 性別変更時のハンドラー
  const handleGenderChange = (gender: Gender) => {
    // 親コンポーネントの期待する形式でイベントを作成
    const syntheticEvent = {
      target: {
        name: "gender",
        value: gender,
      },
    } as ChangeEvent<HTMLInputElement>;

    onChange(syntheticEvent);
  };

  return (
    <div className="gender-picker">
      <label className="gender-label">
        <input
          type="radio"
          name="gender"
          value="male"
          checked={selectedGender === "male"}
          onChange={() => handleGenderChange("male")}
        />
        Male
      </label>
      <label className="gender-label">
        <input
          type="radio"
          name="gender"
          value="female"
          checked={selectedGender === "female"}
          onChange={() => handleGenderChange("female")}
        />
        Female
      </label>
    </div>
  );
};

export default GenderPicker;
