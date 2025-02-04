import { useCallback, useEffect, useRef, useState } from "react";
import "./RangeSlider.css";

interface RangeSliderProps {
  min?: number;
  max?: number;
  defaultValue?: [number, number];
  step?: number;
  onChange?: (value: [number, number]) => void;
}

export default function RangeSlider({
  min = 0,
  max = 100,
  defaultValue = [20, 37],
  step = 1,
  onChange,
}: RangeSliderProps) {
  const [value, setValue] = useState<[number, number]>(defaultValue);
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
    setValue((prev: number[]) => {
      const next = [...prev] as [number, number];
      next[dragging] = Math.min(Math.max(newValue, min), max);
      return next.sort((a, b) => a - b) as [number, number];
    });
  }, [dragging, min, max, getValueFromPosition]);

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

  useEffect(() => {
    onChange?.(value);
  }, [value, onChange]);

  return (
    <div className="slider-container" ref={containerRef}>
      <div className="slider-track">
        <div
          className="slider-range"
          style={{
            left: `${getPercentage(value[0])}%`,
            width: `${getPercentage(value[1]) - getPercentage(value[0])}%`,
          }}
        />
        {value.map((v: number, i: number) => (
          <div
            key={i}
            className="slider-thumb"
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
}
