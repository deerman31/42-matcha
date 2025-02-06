import React, { useEffect, useState } from "react";

// 日付の形式を型で定義
type DateString = `${string}-${string}-${string}`;

// 選択可能な日付の範囲を型で定義
type DateField = {
  year: string;
  month: string;
  day: string;
};

interface DatePickerProps {
  // 初期値がない場合も考慮
  value: DateString | "";
  onChange: (date: DateString) => void;
}

const DatePicker: React.FC<DatePickerProps> = (
  { value, onChange }: DatePickerProps,
) => {
  const currentYear = new Date().getFullYear();

  const [selectedDate, setSelectedDate] = useState<DateField>({
    year: "",
    month: "",
    day: "",
  });

  useEffect(() => {
    if (value && value.includes("-")) {
      const [year, month, day] = value.split("-");
      setSelectedDate({ year, month, day });
    }
  }, [value]);

  const years = Array.from(
    { length: 100 },
    (_, i) => String(currentYear - i),
  );

  const months = Array.from(
    { length: 12 },
    (_, i) => String(i + 1).padStart(2, "0"),
  );

  const getDaysInMonth = (year: string, month: string): string[] => {
    if (!year || !month) return [];
    const daysInMonth = new Date(Number(year), Number(month), 0).getDate();
    return Array.from(
      { length: daysInMonth },
      (_, i) => String(i + 1).padStart(2, "0"),
    );
  };

  const handleDateChange = (field: keyof DateField, value: string) => {
    const newDate = {
      ...selectedDate,
      [field]: value,
    };
    setSelectedDate(newDate);

    if (newDate.year && newDate.month && newDate.day) {
      // 全ての値が存在する場合のみ onChange を呼び出す
      const dateString =
        `${newDate.year}-${newDate.month}-${newDate.day}` as DateString;
      onChange(dateString);
    }
  };

  return (
    <div>
      <select
        value={selectedDate.year}
        onChange={(e) => handleDateChange("year", e.target.value)}
      >
        <option value="">年</option>
        {years.map((year) => <option key={year} value={year}>{year}年</option>)}
      </select>

      <select
        value={selectedDate.month}
        onChange={(e) => handleDateChange("month", e.target.value)}
      >
        <option value="">月</option>
        {months.map((month) => (
          <option key={month} value={month}>{Number(month)}月</option>
        ))}
      </select>

      <select
        value={selectedDate.day}
        onChange={(e) => handleDateChange("day", e.target.value)}
      >
        <option value="">日</option>
        {getDaysInMonth(selectedDate.year, selectedDate.month).map((day) => (
          <option key={day} value={day}>{Number(day)}日</option>
        ))}
      </select>
    </div>
  );
};

export default DatePicker;
