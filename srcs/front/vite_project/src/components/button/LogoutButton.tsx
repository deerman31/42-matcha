import React, { useState } from "react";
import "./LogoutButton.css";
import { getToken, removeToken } from "../../utils/auth.ts";
import { ErrorResponse, Response } from "../../types/api.ts";

import { useNavigate } from "npm:react-router-dom";

export const Logout: React.FC = () => {
  const navigate = useNavigate();

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [response, setResponse] = useState<string>("");
  const [error, setError] = useState<string>("");

  const handleClick = async (): Promise<void> => {
    setIsLoading(true);
    setError("");

    const token = getToken();

    try {
      const response = await fetch("/api/logout", {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorData: ErrorResponse = await response.json();
        throw new Error(errorData.error || "送信に失敗しました");
      }

      const data: Response = await response.json();
      setResponse(data.message);

      removeToken();
      navigate("/login");
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "An unknown error occurred",
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="logout-button-container">
      <button
        className="logout-submit-button"
        onClick={handleClick}
        disabled={isLoading}
      >
        {isLoading ? "Sending..." : "Logout"}
      </button>

      {response && (
        <div className="logout-response-message">
          Response: {response}
        </div>
      )}

      {error && (
        <div className="logout-error-message">
          Error: {error}
        </div>
      )}
    </div>
  );
};
