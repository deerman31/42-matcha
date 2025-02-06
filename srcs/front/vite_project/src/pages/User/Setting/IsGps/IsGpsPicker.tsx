import React, { ChangeEvent, useEffect, useState } from "react";

// 性別を定義する型
type IsGps = boolean;

// コンポーネントのProps型定義
interface IsGpsPickerProps {
  value: boolean; // 親コンポーネントに合わせてstring型に
  onChange: (e: ChangeEvent<HTMLInputElement>) => void; // 親コンポーネントの期待する型に合わせる
}

const IsGpsPicker: React.FC<IsGpsPickerProps> = ({
  value,
  onChange,
}: IsGpsPickerProps) => {
  // 選択された性別の状態管理
  const [selectedIsGps, setSelectedIsGps] = useState<IsGps>(
    value as IsGps,
  );

  // valueの変更を監視して内部状態を更新
  useEffect(() => {
    setSelectedIsGps(value as IsGps);
  }, [value]);

  // 性別変更時のハンドラー
  const handleIsGpsChange = (IsGps: IsGps) => {
    // 親コンポーネントの期待する形式でイベントを作成
    const syntheticEvent = {
      target: {
        name: "isGpsEnabled",
        value: IsGps,
      },
    } as ChangeEvent<HTMLInputElement>;

    onChange(syntheticEvent);
  };

  return (
    <div className="gender-picker">
      <label className="gender-label">
        <input
          type="radio"
          name="IsGps"
          value="true"
          checked={selectedIsGps === true}
          onChange={() => handleIsGpsChange(true)}
        />
        True
      </label>
      <label className="gender-label">
        <input
          type="radio"
          name="IsGps"
          value="false"
          checked={selectedIsGps === false}
          onChange={() => handleIsGpsChange(false)}
        />
        False
      </label>
    </div>
  );
};

export default IsGpsPicker;
