import { useEffect, useRef } from "react";
import type { MetaFunction, LinksFunction } from "@remix-run/node";
import type { Map as LMap, Marker } from "leaflet";

// 定数の外部化
const TOKYO_COORDINATES = {
    lat: 35.6812,
    lng: 139.7671,
    zoom: 13
} as const;

// 型定義
interface MapInstance {
    map: LMap | null;
    marker: Marker | null;
}

// Leafletのスタイルシート定義
export const links: LinksFunction = () => {
    return [
        {
            rel: "stylesheet",
            href: "https://unpkg.com/leaflet@1.9.4/dist/leaflet.css",
            integrity: "sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY=",
            crossOrigin: "anonymous" // または "use-credentials" に変更
        }
    ];
};

export const meta: MetaFunction = () => {
    return [
        { title: "Simple Leaflet Map" },
        { name: "description", content: "A simple map using Leaflet" },
        { name: "viewport", content: "width=device-width, initial-scale=1" }
    ];
};

export default function Map() {
    // useRefを使用してマップインスタンスを保持
    const mapInstanceRef = useRef<MapInstance>({ map: null, marker: null });
    const isInitializedRef = useRef(false);

    useEffect(() => {
        // すでに初期化されている場合は実行しない
        if (isInitializedRef.current) return;

        const mapContainer = document.getElementById("map");
        if (!mapContainer) {
            console.error("Map container not found");
            return;
        }

        // Leafletの初期化関数
        const initializeMap = async () => {
            try {
                // Dynamic import for Leaflet
                const L = (await import("leaflet")).default;

                // マップの初期化
                const map = L.map("map", {
                    // 安全なオプションを設定
                    preferCanvas: true,
                    attributionControl: true,
                    zoomControl: true,
                    closePopupOnClick: true,
                    maxBoundsViscosity: 1.0
                }).setView(
                    [TOKYO_COORDINATES.lat, TOKYO_COORDINATES.lng],
                    TOKYO_COORDINATES.zoom
                );

                // タイルレイヤーの追加（OpenStreetMap）
                L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
                    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
                    maxZoom: 19,
                    minZoom: 3,
                    // XSS対策
                    detectRetina: true,
                    noWrap: true
                }).addTo(map);

                // マーカーの追加（XSS対策済み）
                const marker = L.marker([TOKYO_COORDINATES.lat, TOKYO_COORDINATES.lng], {
                    // マーカーオプションの設定
                    draggable: false,
                    autoPan: true
                }).addTo(map);

                // エスケープ済みのポップアップコンテンツを設定
                const popupContent = L.Util.template(
                    '<div class="popup-content">東京</div>',
                    { escape: true }
                );
                marker.bindPopup(popupContent);

                // インスタンスを保存
                mapInstanceRef.current = { map, marker };
                isInitializedRef.current = true;

            } catch (error) {
                console.error("Failed to initialize map:", error);
                // エラー時のUI表示
                const errorDiv = document.createElement("div");
                errorDiv.className = "map-error";
                errorDiv.textContent = "地図の読み込みに失敗しました。";
                mapContainer.appendChild(errorDiv);
            }
        };

        // クライアントサイドでのみ実行
        if (typeof window !== "undefined") {
            initializeMap();
        }

        // クリーンアップ関数
        return () => {
            if (mapInstanceRef.current.map) {
                // マーカーの削除
                if (mapInstanceRef.current.marker) {
                    mapInstanceRef.current.marker.remove();
                }
                // マップの削除
                mapInstanceRef.current.map.remove();
                // 参照のクリア
                mapInstanceRef.current = { map: null, marker: null };
                isInitializedRef.current = false;
            }
        };
    }, []);

    return (
        <div className="map-container">
            <h1>Leaflet Map Example</h1>
            <div
                id="map"
                style={{
                    height: "300px",
                    width: "100%",
                    position: "relative",
                    zIndex: 1
                }}
                role="application"
                aria-label="地図"
            />
            <style>{`
        .map-container {
          padding: 1rem;
          max-width: 100%;
          overflow: hidden;
        }
        .map-error {
          padding: 1rem;
          color: #721c24;
          background-color: #f8d7da;
          border: 1px solid #f5c6cb;
          border-radius: 4px;
          margin-top: 1rem;
        }
        .popup-content {
          padding: 0.5rem;
          max-width: 200px;
          word-break: break-word;
        }
      `}</style>
        </div>
    );
}