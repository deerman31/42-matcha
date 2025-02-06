import { useEffect, useState } from "react";
import { getToken } from "../../../utils/auth.ts";
import { ErrorResponse, OtherProfileResponse } from "../../../types/api.ts";
import "./ProfileModal.css";
import AllOtherImage from "../../../components/Image/AllOtherImage/AllOtherImage.tsx";

interface ProfileModalProps {
  username: string;
  isOpen: boolean;
  onClose: () => void;
}

const ProfileModal = ({ username, isOpen, onClose }: ProfileModalProps) => {
  const [profileData, setProfileData] = useState<OtherProfileResponse>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchOtherProfile = async () => {
      if (!isOpen) return;

      try {
        const token = getToken();
        const response = await fetch(
          `/api/other-users/get/profile/${username}`,
          {
            method: "GET",
            headers: {
              "Authorization": `Bearer ${token}`,
            },
          },
        );

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
          error instanceof Error
            ? error.message
            : "予期せぬエラーが発生しました",
        );
      } finally {
        setIsLoading(false);
      }
    };

    fetchOtherProfile();
  }, [username, isOpen]);

  if (!isOpen) return null;

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <button
          onClick={onClose}
          className="close-button"
        >
          ✕
        </button>

        {isLoading
          ? <div className="loading">読み込み中...</div>
          : error
          ? <div className="error">{error}</div>
          : profileData
          ? (
            <>
              <h1 className="modal-title">プロフィール</h1>

              <div className="section">
                <h2 className="section-title">画像</h2>
                <AllOtherImage username={profileData.username} />
              </div>

              <div className="section">
                <h2 className="section-title">基本情報</h2>
                <div className="profile-field">
                  <span>ユーザー名: {profileData.username}</span>
                </div>
              </div>

              <div className="section">
                <h2 className="section-title">個人情報</h2>
                <div className="profile-field">
                  <span>Age: {profileData.age}</span>
                </div>
                <div className="profile-field">
                  <span>Gender: {profileData.gender}</span>
                </div>
                <div className="profile-field">
                  <span>Sexuality: {profileData.sexuality}</span>
                </div>
              </div>

              <div className="section">
                <h2 className="section-title">地域情報</h2>
                <div className="profile-field">
                  <span>エリア: {profileData.area}</span>
                </div>
                <div className="profile-field">
                  <span>Distance: {profileData.distance}</span>
                </div>
              </div>

              <div className="section">
                <h2 className="section-title">自己紹介</h2>
                <p>{profileData.self_intro}</p>
              </div>

              <div className="section">
                <h2 className="section-title">評価</h2>
                <div className="profile-field">
                  <span>評価スコア: {profileData.fame_rating}</span>
                </div>
              </div>

              {Array.isArray(profileData.tags) && profileData.tags.length > 0 &&
                (
                  <div className="section">
                    <h2 className="section-title">タグ</h2>
                    <div className="tags-container">
                      {profileData.tags.map((tag: string, index: number) => (
                        <span key={index} className="tag">
                          {tag}
                        </span>
                      ))}
                    </div>
                  </div>
                )}
            </>
          )
          : (
            <div className="error">
              プロフィールデータが見つかりません
            </div>
          )}
      </div>
    </div>
  );
};

export default ProfileModal;
