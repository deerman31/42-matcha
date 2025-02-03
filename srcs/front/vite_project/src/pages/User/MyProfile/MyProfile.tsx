import React, { useEffect, useState } from "react";
import "./MyProfile.css";
import { ErrorResponse, MyProfileResponse } from "../../../types/api.ts";
import { getToken } from "../../../utils/auth.ts";
import ShowLocationMap from "../../../components/LocationMap/ShowLocationMap.tsx";
import { useLocationContext } from "../../../components/LocationService/LocationContextType.tsx";

import ShowTags from "./ShowTags/ShowTags.tsx";

const MyProfile: React.FC = () => {
  const [profileData, setProfileData] = useState<MyProfileResponse | null>(
    null,
  );
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const { setTrackingEnabled } = useLocationContext();

  const token = getToken();


    const fetchProfile = async () => {
      try {
        const response = await fetch("/api/my-profile", {
          method: "GET",
          headers: {
            "Authorization": `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          const errorData: ErrorResponse = await response.json();
          throw new Error(
            errorData.error || "プロフィールの取得に失敗しました",
          );
        }
        const data: MyProfileResponse = await response.json();
        setProfileData(data);

        setTrackingEnabled(data.my_info.is_gps);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "予期せぬエラーが発生しました",
        );
      } finally {
        setIsLoading(false);
      }
    };

  useEffect(() => {

    fetchProfile();
  }, []);

  if (isLoading) {
    return <div className="loading">読み込み中...</div>;
  }

  if (error || profileData?.error) {
    return <div className="error">{error || profileData?.error}</div>;
  }

  if (!profileData?.my_info) {
    return <div className="error">プロフィールデータが見つかりません</div>;
  }

  const { my_info: profile } = profileData;

  // setTrackingEnabled(profile.is_gps);

  const is_gps_string = profile.is_gps ? "OK" : "NG";

  return (
    <div className="profile-container">
      <h1 className="profile-title">プロフィール</h1>

      <div className="profile-section">
        <h2>基本情報</h2>
        <div className="profile-field">
          <label>ユーザー名:</label>
          <span>{profile.username}</span>
        </div>
        <div className="profile-field">
          <label>メールアドレス:</label>
          <span>{profile.email}</span>
        </div>
      </div>

      <div className="profile-section">
        <h2>個人情報</h2>
        <div className="profile-field">
          <label>氏名:</label>
          <span>{profile.lastname} {profile.firstname}</span>
        </div>
        <div className="profile-field">
          <label>生年月日:</label>
          <span>{profile.birthdate}</span>
        </div>
        <div className="profile-field">
          <label>性別:</label>
          <span>{profile.gender}</span>
        </div>
        <div className="profile-field">
          <label>セクシュアリティ:</label>
          <span>{profile.sexuality}</span>
        </div>
      </div>

      <div className="profile-section">
        <h2>地域情報</h2>
        <div className="profile-field">
          <label>エリア:</label>
          <span>{profile.area}</span>
        </div>
        <div className="profile-field">
          <label>位置情報許可</label>
          <span>{is_gps_string}</span>
        </div>

        <div className="profile-field">
          <label>位置情報:</label>
          <span>緯度: {profile.latitude}, 経度: {profile.longitude}</span>
        </div>
        <ShowLocationMap
          latitude={profile.latitude}
          longitude={profile.longitude}
          zoom={14}
          className="profile-map"
        />
      </div>

      <div className="profile-section">
        <h2>自己紹介</h2>
        <p className="self-intro">{profile.self_intro}</p>
      </div>

      <div className="profile-section">
        <ShowTags
        tags={profile.tags}
        refetchTags={fetchProfile} />
      </div>

      <div className="profile-section">
        <h2>評価</h2>
        <div className="profile-field">
          <label>評価スコア:</label>
          <span>{profile.fame_rating}</span>
        </div>
      </div>
    </div>
  );
};

export default MyProfile;
