import { getToken } from "../../../utils/auth.ts";
import { ErrorResponse, OtherProfileResponse } from "../../../types/api.ts";
import { useEffect, useState } from "react";
import { useParams } from "npm:react-router-dom";

import "./OtherProfile.css";

const OtherProfile: React.FC = () => {
  // useParamsを使用してURLパラメータからusernameを取得
  const { username } = useParams<{ username: string }>();

  const [profileData, setProfileData] = useState<OtherProfileResponse>(
    null,
  );
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchOtherProfile = async () => {
    try {
      const token = getToken();
      const response = await fetch(`/api/other-users/get/profile/${username}`, {
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
      const data: OtherProfileResponse = await response.json();

      setProfileData(data.other_profile);

    } catch (error) {
      setError(
        error instanceof Error ? error.message : "予期せぬエラーが発生しました",
      );
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchOtherProfile();
  }, [username]);

  if (isLoading) {
    return <div className="loading">読み込み中...</div>;
  }

  if (error || profileData?.error) {
    return <div className="error">{error || profileData?.error}</div>;
  }

  if (!profileData) {
    return <div className="error">プロフィールデータが見つかりません</div>;
  }

  return (
    <div className="profile-container">
      <h1 className="profile-title">プロフィール</h1>

      {/* <AllMyImage /> */}

      <div className="profile-section">
        <h2>基本情報</h2>
        <div className="profile-field">
          <label>ユーザー名:</label>
          <span>{profileData.username}</span>
        </div>
      </div>

      <div className="profile-section">
        <h2>個人情報</h2>
        <div className="profile-field">
          <label>Age:</label>
          <span>{profileData.age}</span>
        </div>
        <div className="profile-field">
          <label>Gender:</label>
          <span>{profileData.gender}</span>
        </div>
        <div className="profile-field">
          <label>Sexuality:</label>
          <span>{profileData.sexuality}</span>
        </div>
      </div>

      <div className="profile-section">
        <h2>地域情報</h2>
        <div className="profile-field">
          <label>エリア:</label>
          <span>{profileData.area}</span>
        </div>

        <div className="profile-section">
          <h2>自己紹介</h2>
          <p className="self-intro">{profileData.self_intro}</p>
        </div>

        <div className="profile-section">
          <h2>Distance</h2>
          <p className="distance">{profileData.distance}</p>
        </div>

        <div className="profile-section">
          <h2>評価</h2>
          <div className="profile-field">
            <label>評価スコア:</label>
            <span>{profileData.fame_rating}</span>
          </div>
        </div>

        <h2>タグ</h2>

        {Array.isArray(profileData.tags) && (
          <div className="tags">
            {profileData.tags.map((tag: string, index: number) => (
              <span key={index} className="tag">
                {tag}
              </span>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default OtherProfile;
