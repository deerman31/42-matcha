import { useState, useEffect } from "react";
import BrowseConditions from "../../../components/BrowseConditions/BrowseConditions.tsx";
import { UserInfo } from "../../../types/api.ts";
import BrowseUsers from "./BrowseUsers/BrowseUsers.tsx";


const Browse = () => {
  const [userInfos, setUserInfos] = useState<UserInfo[]>([]);

  // userInfos の変更を監視
  // useEffect(() => {
  //   console.log("userInfos が更新されました:", userInfos);
  // }, [userInfos]);

  const handleBrowseComplete = async (newUserInfos: UserInfo[]) => {
    if (Array.isArray(newUserInfos)) {
      await setUserInfos([...newUserInfos]); // 新しい配列として設定
    }
  };

  return (
    <div>
      <BrowseConditions onBrowseComplete={handleBrowseComplete} />
      <BrowseUsers userInfos={userInfos} />
    </div>
  );
};

export default Browse;
