import { getToken } from "../../../../utils/auth.ts";
import "./ShowTags.css";

import React from "react";
import { X } from 'npm:lucide-react'; // バツアイコンのインポート

interface ShowTagsProps {
  tags: string[];
  refetchTags: () => void;  // 追加
}

const ShowTags: React.FC<ShowTagsProps> = (
  { tags,refetchTags }: ShowTagsProps,
) => {
  const deleteTag = async (tag: string) => {
    const token = getToken();
    try {
      const response = await fetch("/api/tag/delete", {
        method: "DELETE",
        headers: {
          "Authorization": `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ tag }),
      });

      if (!response.ok) {
        throw new Error("Failed to delete tag");
      }

      return await response.json();
    } catch (error) {
      console.error("Error deleting tag:", error);
      throw error;
    }
  };

  const handleDelete = async (tag: string) => {
    await deleteTag(tag);
    // 削除後にデータを再取得
    refetchTags();
  };

  return (
    <>
      <h2>タグ</h2>
      <div className="tags">
        {tags.map((tag: string, index: number) => (
          <span key={index} className="tag">
            {tag}
            <button
              onClick={() =>
                handleDelete(tag)}
              className="delete-button"
              aria-label={`Delete ${tag} tag`}
            >
              <X size={14} />
            </button>
          </span>
        ))}
      </div>
    </>
  );
};

export default ShowTags;
