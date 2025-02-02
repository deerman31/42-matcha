import React, { ChangeEvent, useEffect, useState } from "react";

// 性別を定義する型
type Gender = "male" | "female" | "male/female";

// コンポーネントのProps型定義
interface SexualityPickerProps {
  value: string; // 親コンポーネントに合わせてstring型に
  onChange: (e: ChangeEvent<HTMLInputElement>) => void; // 親コンポーネントの期待する型に合わせる
}

const SexualityPicker: React.FC<SexualityPickerProps> = ({
  value,
  onChange,
}: SexualityPickerProps) => {
  // 選択された性別の状態管理
  const [selectedGender, setSelectedGender] = useState<Gender>(
    value as Gender || "male",
  );

  // valueの変更を監視して内部状態を更新
  useEffect(() => {
    setSelectedGender(value as Gender);
  }, [value]);

  // 性別変更時のハンドラー
  const handleGenderChange = (sexuality: Gender) => {
    // 親コンポーネントの期待する形式でイベントを作成
    const syntheticEvent = {
      target: {
        name: "sexuality",
        value: sexuality,
      },
    } as ChangeEvent<HTMLInputElement>;

    onChange(syntheticEvent);
  };

  return (
    <div className="sexuality-picker">
      <label className="sexuality-label">
        <input
          type="radio"
          name="sexuality"
          value="male"
          checked={selectedGender === "male"}
          onChange={() => handleGenderChange("male")}
        />
        Male
      </label>

      <label className="sexuality-label">
        <input
          type="radio"
          name="sexuality"
          value="female"
          checked={selectedGender === "female"}
          onChange={() => handleGenderChange("female")}
        />
        Female
      </label>

      <label className="sexuality-label">
        <input
          type="radio"
          name="sexuality"
          value="male/female"
          checked={selectedGender === "male/female"}
          onChange={() => handleGenderChange("male/female")}
        />
        Male/Female
      </label>
    </div>
  );
};

export default SexualityPicker;
