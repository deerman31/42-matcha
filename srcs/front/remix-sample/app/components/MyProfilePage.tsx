import React, { useEffect, useRef } from 'react';

interface MapInstance {
    map: any;
    marker: any;
}

interface MapProps {
    lat?: number;
    lng?: number;
    zoom?: number;
    title?: string;
    markerPopup?: string;
    className?: string; // クラス名を受け取れるように追加
}

const DEFAULT_COORDINATES = {
    lat: 35.6812,
    lng: 139.7671,
    zoom: 13
};

const LeafletMap = ({ 
    lat = DEFAULT_COORDINATES.lat,
    lng = DEFAULT_COORDINATES.lng,
    zoom = DEFAULT_COORDINATES.zoom,
    markerPopup = "現在地",
    className = "" // デフォルトは空文字
}: MapProps) => {
    const mapInstanceRef = useRef<MapInstance>({ map: null, marker: null });
    const isInitializedRef = useRef(false);

    useEffect(() => {
        if (isInitializedRef.current) {
            const { map, marker } = mapInstanceRef.current;
            if (map && marker) {
                map.setView([lat, lng], zoom);
                marker.setLatLng([lat, lng]);
            }
            return;
        }

        const initializeMap = async () => {
            try {
                const L = (await import('leaflet')).default;

                const map = L.map('profile-map', {
                    preferCanvas: true,
                    attributionControl: true,
                    zoomControl: true,
                    closePopupOnClick: true,
                    maxBoundsViscosity: 1.0
                }).setView([lat, lng], zoom);

                L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
                    maxZoom: 19,
                    minZoom: 3,
                    detectRetina: true,
                    noWrap: true
                }).addTo(map);

                const marker = L.marker([lat, lng], {
                    draggable: false,
                    autoPan: true
                }).addTo(map);

                const popupContent = L.Util.template(
                    `<div class="p-2 max-w-sm break-words">${markerPopup}</div>`,
                    { escape: true }
                );
                marker.bindPopup(popupContent);

                mapInstanceRef.current = { map, marker };
                isInitializedRef.current = true;

            } catch (error) {
                console.error('Failed to initialize map:', error);
                const mapContainer = document.getElementById('profile-map');
                if (mapContainer) {
                    const errorElement = document.createElement('div');
                    errorElement.className = 'p-4 text-red-800 bg-red-100 border border-red-200 rounded-lg mt-4';
                    errorElement.textContent = '地図の読み込みに失敗しました。';
                    mapContainer.appendChild(errorElement);
                }
            }
        };

        if (typeof window !== 'undefined') {
            initializeMap();
        }

        return () => {
            if (mapInstanceRef.current.map) {
                if (mapInstanceRef.current.marker) {
                    mapInstanceRef.current.marker.remove();
                }
                mapInstanceRef.current.map.remove();
                mapInstanceRef.current = { map: null, marker: null };
                isInitializedRef.current = false;
            }
        };
    }, [lat, lng, zoom, markerPopup]);

    return (
        <div className={`w-full ${className}`}>
            <div
                id="profile-map"
                className="h-64 w-full relative z-10 rounded-lg shadow-md"
                role="application"
                aria-label="地図"
            />
        </div>
    );
};

export default LeafletMap;