import React, { ChangeEvent, useEffect, useState } from "react";

// 性別を定義する型
type SortOptionType = "age" | "distance" | "fame_rating" | "tag";

// コンポーネントのProps型定義
interface SortOptionPickerProps {
  value: string; // 親コンポーネントに合わせてstring型に
  onChange: (e: ChangeEvent<HTMLInputElement>) => void; // 親コンポーネントの期待する型に合わせる
}

const SortOptionPicker: React.FC<SortOptionPickerProps> = ({
  value,
  onChange,
}: SortOptionPickerProps) => {
  // 選択された性別の状態管理
  const [selectedSortOption, setSelectedSortOption] = useState<SortOptionType>(
    value as SortOptionType || value,
  );

  // valueの変更を監視して内部状態を更新
  useEffect(() => {
    setSelectedSortOption(value as SortOptionType);
  }, [value]);

  // 性別変更時のハンドラー
  const handleSortOptionChange = (sort_option: SortOptionType) => {
    // 親コンポーネントの期待する形式でイベントを作成
    const syntheticEvent = {
      target: {
        name: "sort_option",
        value: sort_option,
      },
    } as ChangeEvent<HTMLInputElement>;

    onChange(syntheticEvent);
  };

  return (
    <div className="sort-option-picker">
      <label className="sort-option-label">
        <input
          type="radio"
          name="sort_option"
          value="age"
          checked={selectedSortOption === "age"}
          onChange={() => handleSortOptionChange("age")}
        />
        Age
      </label>

      <label className="sort-option-label">
        <input
          type="radio"
          name="sort_option"
          value="distance"
          checked={selectedSortOption === "distance"}
          onChange={() => handleSortOptionChange("distance")}
        />
        Distance
      </label>

      <label className="sort-option-label">
        <input
          type="radio"
          name="sort_option"
          value="fame_rating"
          checked={selectedSortOption === "fame_rating"}
          onChange={() => handleSortOptionChange("fame_rating")}
        />
        FameRating
      </label>

      <label className="sort-option-label">
        <input
          type="radio"
          name="sort_option"
          value="tag"
          checked={selectedSortOption === "tag"}
          onChange={() => handleSortOptionChange("tag")}
        />
        Tag
      </label>
    </div>
  );
};

export default SortOptionPicker;
