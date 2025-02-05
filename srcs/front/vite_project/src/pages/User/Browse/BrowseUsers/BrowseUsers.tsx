import { UserInfo } from "../../../../types/api.ts";
import BrowseUser from "./User/User.tsx";
import { useState,useEffect } from "react";
import "./BrowseUsers.css";

interface BrowseUsersProps {
  userInfos: UserInfo[];
}

const BrowseUsers = ({ userInfos }: BrowseUsersProps) => {
  // コンポーネントの初期マウント時にのみユーザー配列を初期化
  const [remainingUsers, setRemainingUsers] = useState<UserInfo[]>([]);
  
  useEffect(() => {
    setRemainingUsers(userInfos);
  }, [userInfos]);

  if (!remainingUsers || remainingUsers.length === 0) {
    return (
      <div className="no-users-message">
        表示できるユーザーが見つかりませんでした。
      </div>
    );
  }

  const handleResponse = (accepted: boolean) => {
    setRemainingUsers(prev => prev.slice(1));
  };

  // 常に配列の最初のユーザーを表示
  const currentUser = remainingUsers[0];

  return (
    <div className="user-card-container">
      <div className="user-card">
        <BrowseUser
          username={currentUser.username}
          age={currentUser.age}
          distance_km={currentUser.distance_km}
          common_tag_count={currentUser.common_tag_count}
          fame_rating={currentUser.fame_rating}
          image_path={currentUser.image_path}
        />
      </div>
      <div className="action-buttons">
        <button 
          className="action-button reject"
          onClick={() => handleResponse(false)}
        >
          ✕
        </button>
        <button 
          className="action-button accept"
          onClick={() => handleResponse(true)}
        >
          ○
        </button>
      </div>
    </div>
  );

};

export default BrowseUsers;