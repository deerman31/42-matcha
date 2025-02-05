import { UserInfo } from "../../../../types/api.ts";
import BrowseUser from "./User/User.tsx";
import "./BrowseUsers.css";

interface BrowseUsersProps {
  userInfos: UserInfo[];
}

const BrowseUsers = ({ userInfos }: BrowseUsersProps) => {
  if (!userInfos || userInfos.length === 0) {
    return (
      <div className="no-users-message">
        表示できるユーザーが見つかりませんでした。
      </div>
    );
  }

  return (
    <div className="user-cards-grid">
      {userInfos.map((user) => (
        <BrowseUser
          key={user.username}
          username={user.username}
          age={user.age}
          distance_km={user.distance_km}
          common_tag_count={user.common_tag_count}
          fame_rating={user.fame_rating}
          image_path={user.image_path}
        />
      ))}
    </div>
  );
};

export default BrowseUsers;