import React, { useEffect, useState } from "react";

// 都道府県を表す型
type Area =
  | "Hokkaido"
  | "Aomori"
  | "Iwate"
  | "Miyagi"
  | "Akita"
  | "Yamagata"
  | "Fukushima"
  | "Ibaraki"
  | "Tochigi"
  | "Gunma"
  | "Saitama"
  | "Chiba"
  | "Tokyo"
  | "Kanagawa"
  | "Niigata"
  | "Toyama"
  | "Ishikawa"
  | "Fukui"
  | "Yamanashi"
  | "Nagano"
  | "Gifu"
  | "Shizuoka"
  | "Aichi"
  | "Mie"
  | "Shiga"
  | "Kyoto"
  | "Osaka"
  | "Hyogo"
  | "Nara"
  | "Wakayama"
  | "Tottori"
  | "Shimane"
  | "Okayama"
  | "Hiroshima"
  | "Yamaguchi"
  | "Tokushima"
  | "Kagawa"
  | "Ehime"
  | "Kochi"
  | "Fukuoka"
  | "Saga"
  | "Nagasaki"
  | "Kumamoto"
  | "Oita"
  | "Miyazaki"
  | "Kagoshima"
  | "Okinawa";

// コンポーネントのProps型定義
interface AreaPickerProps {
  value: string;
  onChange: (value: string) => void; // 親コンポーネントの期待する型に合わせる
}

const AreaPicker: React.FC<AreaPickerProps> = ({
  value,
  onChange,
}: AreaPickerProps) => {
  const [selectedArea, setSelectedArea] = useState<Area>(
    value as Area || "",
  );
  // valueの変更を監視して内部状態を更新
  useEffect(() => {
    setSelectedArea(value as Area);
  }, [value]);

  const handleAreaChange = (area: Area) => {
    onChange(area);
  };

  return (
    <div className="area-picker">
      <select
        value={selectedArea}
        onChange={(e) => handleAreaChange(e.target.value as Area)}
        className="area-select"
      >
        <option value="Hokkaido">Hokkaido</option>
        <option value="Aomori">Aomori</option>
        <option value="Iwate">Iwate</option>
        <option value="Miyagi">Miyagi</option>
        <option value="Akita">Akita</option>
        <option value="Yamagata">Yamagata</option>
        <option value="Fukushima">Fukushima</option>
        <option value="Ibaraki">Ibaraki</option>
        <option value="Tochigi">Tochigi</option>
        <option value="Gunma">Gunma</option>
        <option value="Saitama">Saitama</option>
        <option value="Chiba">Chiba</option>
        <option value="Tokyo">Tokyo</option>
        <option value="Kanagawa">Kanagawa</option>
        <option value="Niigata">Niigata</option>
        <option value="Toyama">Toyama</option>
        <option value="Ishikawa">Ishikawa</option>
        <option value="Fukui">Fukui</option>
        <option value="Yamanashi">Yamanashi</option>
        <option value="Nagano">Nagano</option>
        <option value="Gifu">Gifu</option>
        <option value="Shizuoka">Shizuoka</option>
        <option value="Aichi">Aichi</option>
        <option value="Mie">Mie</option>
        <option value="Shiga">Shiga</option>
        <option value="Kyoto">Kyoto</option>
        <option value="Osaka">Osaka</option>
        <option value="Hyogo">Hyogo</option>
        <option value="Nara">Nara</option>
        <option value="Wakayama">Wakayama</option>
        <option value="Tottori">Tottori</option>
        <option value="Shimane">Shimane</option>
        <option value="Okayama">Okayama</option>
        <option value="Hiroshima">Hiroshima</option>
        <option value="Yamaguchi">Yamaguchi</option>
        <option value="Tokushima">Tokushima</option>
        <option value="Kagawa">Kagawa</option>
        <option value="Ehime">Ehime</option>
        <option value="Kochi">Kochi</option>
        <option value="Fukuoka">Fukuoka</option>
        <option value="Saga">Saga</option>
        <option value="Nagasaki">Nagasaki</option>
        <option value="Kumamoto">Kumamoto</option>
        <option value="Oita">Oita</option>
        <option value="Miyazaki">Miyazaki</option>
        <option value="Kagoshima">Kagoshima</option>
        <option value="Okinawa">Okinawa</option>
      </select>
    </div>
  );
};

export default AreaPicker;
