import { Heart, MessageCircle, Search, Users } from "npm:lucide-react";
import React from "react";
import "./Landing.css";
import { Link } from "npm:react-router-dom";

const Landing: React.FC = () => {
  const features = [
    {
      icon: <Users size={24} />,
      title: "マッチング",
      description: "あなたに最適な相手をご紹介",
    },
    {
      icon: <MessageCircle size={24} />,
      title: "リアルタイムチャット",
      description: "マッチした相手とすぐに会話",
    },
    {
      icon: <Search size={24} />,
      title: "詳細検索",
      description: "好みの条件で相手を探せる",
    },
  ];
  return (
    <div className="landing">
      <div className="container">
        <div className="hero">
          <h1 className="hero__title">
            Matcha
            <Heart className="text-red-500" size={32} />
          </h1>
          <p className="hero__description">
            あなたの運命の出会いがここに
          </p>
          <div className="button-group">
            <Link to="register" className="button button--primary">
              Register
            </Link>
            <Link to="login" className="button button--secondary">
              Login
            </Link>
          </div>
        </div>

        <div className="features">
          {features.map((feature, index) => (
            <div key={index} className="feature-card">
              <div className="feature-card__icon">
                {feature.icon}
              </div>
              <h3 className="feature-card__title">
                {feature.title}
              </h3>
              <p className="feature-card__description">
                {feature.description}
              </p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Landing;
