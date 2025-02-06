import React, { useEffect, useState } from "react";
import { Loader2, Plus, Tag, X } from "npm:lucide-react";
import { getToken } from "../../../../utils/auth.ts";
import "./TagSearchComponent.css";

interface SearchResponse {
  tags: string[];
  error?: string;
}

interface SetTagResponse {
  message?: string;
  error?: string;
}

const TagSearchComponent: React.FC = () => {
  const [searchTerm, setSearchTerm] = useState("");
  const [searchResults, setSearchResults] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  useEffect(() => {
    const searchTags = async () => {
      if (searchTerm.length < 3) {
        setSearchResults([]);
        return;
      }

      setIsLoading(true);
      setError("");

      try {
        const token = getToken();

        const response = await fetch("/api/tag/search", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
          },
          body: JSON.stringify({ tagname: searchTerm }),
        });

        if (!response.ok) {
          throw new Error("Search failed");
        }

        const data: SearchResponse = await response.json();
        if (data.tags) {
          setSearchResults(data.tags);
        }
      } catch (err) {
        setError("Failed to search tags");
        setSearchResults([]);
      } finally {
        setIsLoading(false);
      }
    };

    const debounceTimeout = setTimeout(searchTags, 300);
    return () => clearTimeout(debounceTimeout);
  }, [searchTerm]);

  const handleTagSelect = async (tag: string) => {
    setIsLoading(true);
    setError("");
    setSuccess("");

    try {
      const token = getToken();

      const response = await fetch("/api/tag/set", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({ tag }),
      });

      if (!response.ok) {
        throw new Error("Failed to set tag");
      }

      setSuccess(`Tag "${tag}" has been added successfully`);
      setSearchTerm("");
      setSearchResults([]);
    } catch (err) {
      setError("Failed to add tag");
    } finally {
      setIsLoading(false);
    }
  };

  const handleCreateNewTag = () => {
    if (searchTerm.length < 3) return;
    handleTagSelect(searchTerm);
  };

  return (
    <div className="tag-search">
      <div className="input-area">
        <div className="search-input">
          <Tag className="search-icon" />
          <input
            type="text"
            placeholder="Search tags (minimum 3 characters)"
            value={searchTerm}
            onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
              setSearchTerm(e.target.value)}
            disabled={isLoading}
            className="tag-input"
          />
          {searchTerm && (
            <button
              className="clear-btn"
              onClick={() => setSearchTerm("")}
              type="button"
            >
              <X className="clear-icon" />
            </button>
          )}
        </div>
      </div>

      {isLoading && (
        <div className="loader">
          <Loader2 className="loader-icon" />
        </div>
      )}

      {error && (
        <div className="alert error">
          <p>{error}</p>
        </div>
      )}

      {success && (
        <div className="alert success">
          <p>{success}</p>
        </div>
      )}

      {searchResults.length > 0
        ? (
          <div className="results-card">
            <div className="results-content">
              <div className="tag-grid">
                {searchResults.map((tag: string, index: number) => (
                  <button
                    key={index}
                    onClick={() => handleTagSelect(tag)}
                    disabled={isLoading}
                    className="tag-btn"
                  >
                    {tag}
                  </button>
                ))}
              </div>
            </div>
          </div>
        )
        : (
          searchTerm.length >= 3 && !isLoading && !error && (
            <div className="create-new-tag">
              <p>No matching tags found</p>
              <button
                onClick={handleCreateNewTag}
                className="create-tag-btn"
                disabled={isLoading}
              >
                <Plus className="plus-icon" />
                Create new tag "{searchTerm}"
              </button>
            </div>
          )
        )}
    </div>
  );
};

export default TagSearchComponent;
