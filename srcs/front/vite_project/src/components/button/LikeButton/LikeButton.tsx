import { ErrorResponse } from "../../../types/api.ts";
import { getToken } from "../../../utils/auth.ts";
import "./LikeButton.css";

interface LikeButtonProps {
  username: string;
  onClick: () => void;
}

const LikeButton: React.FC<LikeButtonProps> = (
  { username, onClick }: LikeButtonProps,
) => {
  const handleClick = async (): Promise<void> => {
    const token = getToken();

    try {
      const response = await fetch("/api/friend/like/do-like", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({ username: username }),
      });
      if (!response.ok) {
        const errorData: ErrorResponse = await response.json();
        throw new Error(errorData.error || "送信に失敗しました");
      }
      onClick();
      // const data:Response = await response.json();
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <button
      className="like-button"
      onClick={handleClick}
      aria-label="いいね"
    >
      <span className="like-icon">♥</span>
      <span className="like-text">いいね</span>
    </button>
  );
};

export default LikeButton;
