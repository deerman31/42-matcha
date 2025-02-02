// NavigationPage.tsx
import React from 'react';
import { Link } from 'npm:react-router-dom';
import './Setting.css';

const Setting: React.FC = () => {
  return (
    <div className="navigation-container">
      <div className="button-container">
        <Link to="/change-username" className="nav-button">
          UserName
        </Link>
      </div>
    </div>
  );
};

export default Setting;