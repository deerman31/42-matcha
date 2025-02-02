import React, { useEffect, useRef } from "react";
import L from "leaflet";
import "leaflet/dist/leaflet.css";
import "./LocationMap.css";

interface LocationMapProps {
  latitude: number;
  longitude: number;
  zoom?: number;
  className?: string;
}

const LocationMap: React.FC<LocationMapProps> = ({
  latitude,
  longitude,
  zoom = 13,
  className = "",
}: LocationMapProps) => {
  const mapRef = useRef<HTMLDivElement>(null);
  const mapInstanceRef = useRef<L.Map | null>(null);

  useEffect(() => {
    if (!mapRef.current) return;

    // カスタムマーカーアイコンの作成
    const customIcon = L.divIcon({
      className: "custom-marker",
      html: '<div class="marker-pin"></div>',
      iconSize: [30, 30],
      iconAnchor: [15, 30],
    });

    // 地図が既に初期化されていない場合のみ初期化
    if (!mapInstanceRef.current) {
      mapInstanceRef.current = L.map(mapRef.current).setView([
        latitude,
        longitude,
      ], zoom);

      // OpenStreetMapのタイルレイヤーを追加
      L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        attribution: "© OpenStreetMap contributors",
        maxZoom: 19,
      }).addTo(mapInstanceRef.current);

      // マーカーを追加
      L.marker([latitude, longitude], { icon: customIcon })
        .addTo(mapInstanceRef.current)
        .bindPopup(`緯度: ${latitude}<br>経度: ${longitude}`)
        .openPopup();
    }

    // クリーンアップ関数
    return () => {
      if (mapInstanceRef.current) {
        mapInstanceRef.current.remove();
        mapInstanceRef.current = null;
      }
    };
  }, [latitude, longitude, zoom]);

  // 位置情報が変更された場合の処理
  useEffect(() => {
    if (mapInstanceRef.current) {
      mapInstanceRef.current.setView([latitude, longitude], zoom);
    }
  }, [latitude, longitude, zoom]);

  return (
    <div
      ref={mapRef}
      className={`location-map ${className}`}
      aria-label={`緯度 ${latitude} 経度 ${longitude} の地図`}
    />
  );
};

export default LocationMap;
