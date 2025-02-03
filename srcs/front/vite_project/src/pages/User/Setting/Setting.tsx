// NavigationPage.tsx
import React from "react";
import { Link } from "npm:react-router-dom";
import "./Setting.css";

const Setting: React.FC = () => {
  return (
    <div className="navigation-container">
      <div className="button-container">
        <Link to="/change-username" className="nav-button">
          UserName
        </Link>
        <Link to="/change-email" className="nav-button">
          Email
        </Link>
        <Link to="/change-lastname" className="nav-button">
          LastName
        </Link>
        <Link to="/change-firstname" className="nav-button">
          FirstName
        </Link>
        <Link to="/change-birthdate" className="nav-button">
          BirthDate
        </Link>
        <Link to="/change-gender" className="nav-button">
          Gender
        </Link>
        <Link to="/change-sexuality" className="nav-button">
          Sexuality
        </Link>
        <Link to="/change-area" className="nav-button">
          Area
        </Link>
        <Link to="/change-is-gps" className="nav-button">
          IsGps
        </Link>
        <Link to="/change-map" className="nav-button">
          Map
        </Link>
        <Link to="/change-self-intro" className="nav-button">
          SelfIntro
        </Link>
      </div>
    </div>
  );
};

export default Setting;
