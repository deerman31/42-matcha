import React from "react";
import "./UserHeader.css"; // Headerコンポーネントに対応するCSSファイルをインポート
import { Link } from "npm:react-router-dom";

import { Logout } from "../button/LogoutButton.tsx";

const UserHeader: React.FC = () => {
  return (
    <header className="user_header">
      <Link to="/home" className="user_header_logo">
        User
      </Link>
      <nav>
        <ul className="nav-links">
          <li>
            <Link to="/home">Home</Link>
          </li>
          <li>
            <Link to="/my-profile">MyProfile</Link>
          </li>
        </ul>
      </nav>
      <Logout />
    </header>
  );
};

export default UserHeader;
