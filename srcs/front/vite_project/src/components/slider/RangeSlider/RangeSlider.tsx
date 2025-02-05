import { useCallback, useEffect, useRef, useState } from "react";
import "./RangeSlider.css";

interface RangeSliderProps {
  min?: number;
  max?: number;
  value: [number, number];
  step?: number;
  onChange?: (value: [number, number]) => void;
}

const RangeSlider = ({
  min = 0,
  max = 100,
  value,
  step = 1,
  onChange,
}: RangeSliderProps) => {
  // const [value, setValue] = useState<[number, number]>(start_value);
  const [dragging, setDragging] = useState<number | null>(null);
  const containerRef = useRef<HTMLDivElement>(null);

  const getPercentage = useCallback((value: number) => {
    return ((value - min) / (max - min)) * 100;
  }, [min, max]);

  const getValueFromPosition = useCallback((position: number) => {
    const containerRect = containerRef.current?.getBoundingClientRect();
    if (!containerRect) return 0;

    const percentage = (position - containerRect.left) / containerRect.width;
    const value = percentage * (max - min) + min;
    return Math.round(value / step) * step;
  }, [min, max, step]);

  const handleMouseDown =
    (index: number) => (_e: React.MouseEvent<HTMLDivElement>) => {
      setDragging(index);
    };

  const handleMouseMove = useCallback((e: MouseEvent) => {
    if (dragging === null) return;

    const newValue = getValueFromPosition(e.clientX);
    const nextValue = [...value] as [number, number];
    nextValue[dragging] = Math.min(Math.max(newValue, min), max);
    onChange?.(nextValue.sort((a, b) => a - b) as [number, number]); // 直接onChangeを呼び出し
  }, [dragging, min, max, getValueFromPosition, onChange, value]);

  const handleMouseUp = useCallback(() => {
    setDragging(null);
  }, []);

  useEffect(() => {
    if (dragging !== null) {
      globalThis.addEventListener("mousemove", handleMouseMove);
      globalThis.addEventListener("mouseup", handleMouseUp);
      return () => {
        globalThis.removeEventListener("mousemove", handleMouseMove);
        globalThis.removeEventListener("mouseup", handleMouseUp);
      };
    }
  }, [dragging, handleMouseMove, handleMouseUp]);

  return (
    <div className="range_slider-container" ref={containerRef}>
      <div className="range_slider-track">
        <div
          className="range_slider-range"
          style={{
            left: `${getPercentage(value[0])}%`,
            width: `${getPercentage(value[1]) - getPercentage(value[0])}%`,
          }}
        />
        {value.map((v: number, i: number) => (
          <div
            key={i}
            className="range_slider-thumb"
            style={{
              left: `${getPercentage(v)}%`,
            }}
            onMouseDown={handleMouseDown(i)}
          >
            {dragging === i && (
              <div className="value-label">
                {v}
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};

export default RangeSlider;
