import React from "react";
import "./UserHeader.css"; // Headerコンポーネントに対応するCSSファイルをインポート
import { Link } from "npm:react-router-dom";

import { Logout } from "../../button/LogoutButton.tsx";

const UserHeader: React.FC = () => {
  return (
    <header className="user_header">
      <Link to="/my-profile" className="user_header_logo">
        User
      </Link>
      <nav>
        <ul className="nav-links">
          <li>
            <Link to="/my-profile">MyProfile</Link>
          </li>
          <li>
            <Link to="/setting">Setting</Link>
          </li>
          <li>
            <Link to="/brouse">Brouse</Link>
          </li>
        </ul>
      </nav>
      <Logout />
    </header>
  );
};

export default UserHeader;
