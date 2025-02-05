import React from "react";
import "./AuthHeader.css"; // Headerコンポーネントに対応するCSSファイルをインポート
import { Link } from "npm:react-router-dom";

const AuthHeader: React.FC = () => {
  return (
    <header className="auth_header">
      <Link to="/" className="auth_header_logo">
        My Restaurant
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

export default AuthHeader;
