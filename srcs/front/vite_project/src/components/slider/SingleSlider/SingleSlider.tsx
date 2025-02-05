import { useCallback, useEffect, useRef, useState } from "react";
import "./SingleSlider.css";

interface SliderProps {
  min?: number;
  max?: number;
  value: number;
  step?: number;
  onChange?: (value: number) => void;
}

const SingleSlider = ({
  min = 0,
  max = 100,
  value,
  step = 1,
  onChange,
}: SliderProps) => {
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
    onChange?.(clampedValue); // setValueの代わりにonChangeを直接呼び出し
  }, [dragging, min, max, getValueFromPosition, onChange]);

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

  return (
    <div className="single_slider-container" ref={containerRef}>
      <div className="single_slider-track">
        <div
          className="single_slider-range"
          style={{
            width: `${getPercentage(value)}%`,
          }}
        />
        <div
          className="single_slider-thumb"
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
