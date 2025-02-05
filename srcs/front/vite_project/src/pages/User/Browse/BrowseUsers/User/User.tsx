import { getToken } from "../../../../../utils/auth.ts";
import {
  BrowseUserImageResponse,
  ErrorResponse,
} from "../../../../../types/api.ts";
import { useEffect, useState } from "react";
import "./User.css";

interface BrowseUserProps {
  username: string;
  age: number;
  distance_km: number;
  common_tag_count: number;
  fame_rating: number;
  image_path: string;
}

const BrowseUser: React.FC<BrowseUserProps> = (
  { username, age, distance_km, common_tag_count, fame_rating, image_path }:
    BrowseUserProps,
) => {
  const [image, setImage] = useState<string | null>(null);

  useEffect(() => {
    const fetchImage = async () => {
      try {
        const token = getToken();
        const response = await fetch("/api/other-users/get-image", {
          method: "POST",
          headers: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ image_path: image_path }),
        });
        if (!response.ok) {
          const errorData: ErrorResponse = await response.json();
          throw new Error(errorData.error);
        }
        const data: BrowseUserImageResponse = await response.json();
        setImage(data.image);
      } catch (error) {
        console.error("Error Other Image:", error);
      }
    };
    fetchImage();
  }, [image_path]);

  return (
    <div className="user-card">
      <img src={image} alt={`${username}'s profile`} />
      <h2>{username}</h2>
      <p>Age: {age}</p>
      <p>Distance: {distance_km} km</p>
      <p>Common Tags: {common_tag_count}</p>
      <p>Fame Rating: {fame_rating}</p>
    </div>
  );
};

export default BrowseUser;
