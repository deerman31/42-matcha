import { useEffect } from "react";
import { getToken } from "../../utils/auth.ts";
import { useLocationContext } from "./LocationContextType.tsx";

const LocationService = () => {
  const token = getToken();
  const { isTrackingEnabled } = useLocationContext();

  useEffect(() => {

    if (!isTrackingEnabled) {
      return; // トラッキングが無効な場合は何もしない
    }

    // 位置情報をサーバーに送信
    const sendLocationToServer = async (
      latitude: number,
      longitude: number,
    ) => {
      try {
        const response = await fetch("/api/gps/set/location", {
          method: "PATCH",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
          },
          credentials: "include", // cookieを送信するために必要
          body: JSON.stringify({ latitude, longitude }),
        });

        if (!response.ok) {
          throw new Error("Failed to update location");
        }
      } catch (error) {
        console.error("Error sending location:", error);
      }
    };

    // 位置情報を取得
    const updateLocation = () => {
      if (!navigator.geolocation) {
        console.error("Geolocation is not supported");
        return;
      }

      // 位置情報の取得オプション
      const options = {
        enableHighAccuracy: true, // 高精度の位置情報を取得
        timeout: 5000, // 5秒でタイムアウト
        maximumAge: 0, // キャッシュした位置情報を使用しない
      };

      navigator.geolocation.getCurrentPosition(
        // 成功時のコールバック
        (position) => {
          const { latitude, longitude } = position.coords;
          sendLocationToServer(latitude, longitude);
        },
        // エラー時のコールバック
        (error) => {
          switch (error.code) {
            case error.PERMISSION_DENIED:
              console.error("User denied the request for Geolocation.");
              break;
            case error.POSITION_UNAVAILABLE:
              console.error("Location information is unavailable.");
              break;
            case error.TIMEOUT:
              console.error("The request to get user location timed out.");
              break;
            default:
              console.error("An unknown error occurred.");
          }
        },
        options,
      );
    };

    // 初回の位置情報取得
    updateLocation();

    // 5分ごとに位置情報を更新
    //const intervalId = setInterval(updateLocation, 5 * 60 * 1000);
    const intervalId = setInterval(updateLocation, 1 * 60 * 1000);


    // クリーンアップ関数
    return () => {
      clearInterval(intervalId);
    };
  }, [isTrackingEnabled, token]); // 空の依存配列でマウント時のみ実行

  // このコンポーネントはUIを持たない
  return null;
};

export default LocationService;
