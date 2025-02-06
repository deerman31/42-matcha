import React, { ChangeEvent, useEffect, useState } from "react";

// 性別を定義する型
type SortOrder = 0 | 1;

// コンポーネントのProps型定義
interface SortOrderPickerProps {
  value: number; // 親コンポーネントに合わせてstring型に
  onChange: (e: ChangeEvent<HTMLInputElement>) => void; // 親コンポーネントの期待する型に合わせる
}

const SortOrderPicker: React.FC<SortOrderPickerProps> = ({
  value,
  onChange,
}: SortOrderPickerProps) => {
  // 選択された性別の状態管理
  const [selectedSortOrder, setSelectedSortOrder] = useState<SortOrder>(
    value as SortOrder || value,
  );

  // valueの変更を監視して内部状態を更新
  useEffect(() => {
    setSelectedSortOrder(value as SortOrder);
  }, [value]);

  const handleSortOrderChange = (sort_order: SortOrder) => {
    // 親コンポーネントの期待する形式でイベントを作成
    const syntheticEvent = {
      target: {
        name: "sort_order",
        value: sort_order,
      },
    } as ChangeEvent<HTMLInputElement>;

    onChange(syntheticEvent);
  };

  return (
    <div className="sort-order-picker">
      <label className="sort-order-label">
        <input
          type="radio"
          name="sort_order"
          value="descending"
          checked={selectedSortOrder === 0}
          onChange={() => handleSortOrderChange(0)}
        />
        Descending
      </label>

      <label className="sort-order-label">
        <input
          type="radio"
          name="sort_order"
          value="ascending"
          checked={selectedSortOrder === 1}
          onChange={() => handleSortOrderChange(1)}
        />
        Ascending
      </label>
    </div>
  );
};

export default SortOrderPicker;
