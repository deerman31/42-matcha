import { useCallback, useEffect, useRef, useState } from "react";
import "./SingleSlider.css";

interface SliderProps {
  min?: number;
  max?: number;
  defaultValue?: number;
  step?: number;
  onChange?: (value: number) => void;
}

const SingleSlider = ({
  min = 0,
  max = 100,
  defaultValue = 50,
  step = 1,
  onChange,
}: SliderProps) => {
  const [value, setValue] = useState<number>(defaultValue);
  const [dragging, setDragging] = useState<boolean>(false);
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

  const handleMouseDown = (_e: React.MouseEvent<HTMLDivElement>) => {
    setDragging(true);
  };

  const handleMouseMove = useCallback((e: MouseEvent) => {
    if (!dragging) return;

    const newValue = getValueFromPosition(e.clientX);
    const clampedValue = Math.min(Math.max(newValue, min), max);
    setValue(clampedValue);
    // 直接onChangeを呼び出す
     onChange?.(clampedValue);

  }, [dragging, min, max, getValueFromPosition]);

  const handleMouseUp = useCallback(() => {
    setDragging(false);
  }, []);

  useEffect(() => {
    if (dragging) {
      globalThis.addEventListener("mousemove", handleMouseMove);
      globalThis.addEventListener("mouseup", handleMouseUp);
      return () => {
        globalThis.removeEventListener("mousemove", handleMouseMove);
        globalThis.removeEventListener("mouseup", handleMouseUp);
      };
    }
  }, [dragging, handleMouseMove, handleMouseUp]);

  useEffect(() => {
    onChange?.(value);
  }, [value, onChange]);

  return (
    <div className="slider-container" ref={containerRef}>
      <div className="slider-track">
        <div
          className="slider-range"
          style={{
            width: `${getPercentage(value)}%`,
          }}
        />
        <div
          className="slider-thumb"
          style={{
            left: `${getPercentage(value)}%`,
          }}
          onMouseDown={handleMouseDown}
        >
          {dragging && (
            <div className="value-label">
              {value}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default SingleSlider;
