import React, { useEffect, useRef, useState, FormEvent } from "react";
import L from "leaflet";
import { ErrorResponse,Response } from "../../../../types/api.ts";
import { getToken } from "../../../../utils/auth.ts";
import "leaflet/dist/leaflet.css";
import "./InteractiveLocationMap.css";

interface LocationFormData {
  latitude: number;
  longitude: number;
}

// 送信状態の型定義
interface SubmitStatus {
  type: "success" | "error" | "";
  message: string;
}

interface InteractiveLocationMapProps {
  initialLatitude?: number;
  initialLongitude?: number;
  zoom?: number;
  className?: string;
}

const InteractiveLocationMap: React.FC<InteractiveLocationMapProps> = ({
  initialLatitude = 35.6812,
  initialLongitude = 139.7671,
  zoom = 13,
  className = "",
}) => {
  const mapRef = useRef<HTMLDivElement>(null);
  const mapInstanceRef = useRef<L.Map | null>(null);
  const markerRef = useRef<L.Marker | null>(null);

  // 状態管理
  const [coordinates, setCoordinates] = useState<LocationFormData>({
    latitude: initialLatitude,
    longitude: initialLongitude,
  });
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [submitStatus, setSubmitStatus] = useState<SubmitStatus>({
    type: "",
    message: "",
  });

  useEffect(() => {
    if (!mapRef.current) return;

    const customIcon = L.divIcon({
      className: "custom-marker",
      html: '<div class="marker-pin"></div>',
      iconSize: [30, 30],
      iconAnchor: [15, 30],
    });

    if (!mapInstanceRef.current) {
      mapInstanceRef.current = L.map(mapRef.current).setView(
        [initialLatitude, initialLongitude],
        zoom
      );

      L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        attribution: "© OpenStreetMap contributors",
        maxZoom: 19,
      }).addTo(mapInstanceRef.current);

      markerRef.current = L.marker([initialLatitude, initialLongitude], {
        icon: customIcon,
        draggable: true,
      })
        .addTo(mapInstanceRef.current)
        .bindPopup("ドラッグして位置を調整できます");

      mapInstanceRef.current.on("click", (e: L.LeafletMouseEvent) => {
        const { lat, lng } = e.latlng;
        updateMarkerPosition(lat, lng);
      });

      markerRef.current.on("dragend", () => {
        const position = markerRef.current?.getLatLng();
        if (position) {
          updateMarkerPosition(position.lat, position.lng);
        }
      });
    }

    return () => {
      if (mapInstanceRef.current) {
        mapInstanceRef.current.remove();
        mapInstanceRef.current = null;
        markerRef.current = null;
      }
    };
  }, [initialLatitude, initialLongitude, zoom]);

  const updateMarkerPosition = (lat: number, lng: number) => {
    setCoordinates({ latitude: lat, longitude: lng });
    
    if (markerRef.current) {
      markerRef.current.setLatLng([lat, lng]);
      markerRef.current
        .bindPopup(`緯度: ${lat.toFixed(6)}<br>経度: ${lng.toFixed(6)}`)
        .openPopup();
    }
  };

  // フォーム送信ハンドラー
  const handleSubmit = async (e: FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setIsSubmitting(true);
    setSubmitStatus({ type: "", message: "" });

    try {
      const token = getToken();
      const response = await fetch("/api/gps/set/location-alternative", {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(coordinates),
      });

      if (!response.ok) {
        const errorData: ErrorResponse = await response.json();
        throw new Error(errorData.error || "位置情報の更新に失敗しました");
      }

      const data: Response = await response.json();
      setSubmitStatus({
        type: "success",
        message: data.message || "位置情報を更新しました！",
      });

    } catch (error) {
      setSubmitStatus({
        type: "error",
        message: error instanceof Error
          ? error.message
          : "エラーが発生しました。もう一度お試しください。",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="interactive-location-container">
      <h2 className="form-title">位置情報の更新</h2>
      <div
        ref={mapRef}
        className={`location-map ${className}`}
        aria-label="インタラクティブ地図"
      />
      <form onSubmit={handleSubmit} className="location-form">
        <div className="form-group">
          <label htmlFor="latitude" className="form-label">緯度:</label>
          <input
            type="number"
            id="latitude"
            name="latitude"
            value={coordinates.latitude}
            onChange={(e) =>
              setCoordinates({
                ...coordinates,
                latitude: parseFloat(e.target.value),
              })
            }
            step="any"
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="longitude" className="form-label">経度:</label>
          <input
            type="number"
            id="longitude"
            name="longitude"
            value={coordinates.longitude}
            onChange={(e) =>
              setCoordinates({
                ...coordinates,
                longitude: parseFloat(e.target.value),
              })
            }
            step="any"
            required
            className="form-input"
          />
        </div>
        <button
          type="submit"
          disabled={isSubmitting}
          className="submit-button"
        >
          {isSubmitting ? "更新中..." : "位置を更新"}
        </button>
      </form>
      {submitStatus.message && (
        <div
          className={`alert ${
            submitStatus.type === "success" ? "alert-success" : "alert-error"
          }`}
          role="alert"
          aria-live="polite"
        >
          {submitStatus.message}
        </div>
      )}
    </div>
  );
};

export default InteractiveLocationMap;