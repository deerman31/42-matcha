import React from "react";
import "./UserHeader.css"; // Headerコンポーネントに対応するCSSファイルをインポート
import { Link } from "npm:react-router-dom";

const UserHeader: React.FC = () => {
  return (
    <header className="user_header">
      <Link to="/" className="user_header_logo">
        User
      </Link>
      <nav>
        <ul className="nav-links">
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/register">Register</Link>
          </li>
          <li>
            <Link to="/login">Login</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default UserHeader;
