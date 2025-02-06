import React, { useRef, useState } from "react";
import { Upload } from "npm:lucide-react";
import "./ImageUploader.css";

interface ImageUploaderProps {
  onImageSelect: (formData: FormData) => void;
  maxSizeMB?: number;
  acceptedTypes?: string[];
}

const ImageUploader: React.FC<ImageUploaderProps> = ({
  onImageSelect,
  maxSizeMB = 5,
  acceptedTypes = ["image/jpeg", "image/png", "image/gif"],
}: ImageUploaderProps) => {
  const [preview, setPreview] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    setError(null);

    if (!file) return;

    // ファイルタイプの検証
    if (!acceptedTypes.includes(file.type)) {
      setError(`対応している画像形式: ${acceptedTypes.join(", ")}`);
      return;
    }

    // ファイルサイズの検証
    const maxSize = maxSizeMB * 1024 * 1024; // MB to bytes
    if (file.size > maxSize) {
      setError(`ファイルサイズは${maxSizeMB}MB以下にしてください`);
      return;
    }

    // プレビューの生成
    const reader = new FileReader();
    reader.onloadend = () => {
      setPreview(reader.result as string);
    };
    reader.readAsDataURL(file);

    // FormDataの作成と親コンポーネントへの送信
    const formData = new FormData();
    formData.append("image", file);
    onImageSelect(formData);
  };

  const handleClick = () => {
    fileInputRef.current?.click();
  };

  return (
    <div className="image-uploader">
      <div
        className="upload-area"
        onClick={handleClick}
      >
        {preview
          ? <img src={preview} alt="プレビュー" className="preview-image" />
          : (
            <div className="upload-placeholder">
              <Upload className="upload-icon" />
              <p>クリックして画像を選択</p>
              <p className="upload-hint">
                {acceptedTypes.join(", ")} 形式、{maxSizeMB}MB以下
              </p>
            </div>
          )}
        <input
          type="file"
          ref={fileInputRef}
          accept={acceptedTypes.join(",")}
          onChange={handleFileSelect}
          className="file-input"
        />
      </div>
      {error && <p className="error-message">{error}</p>}
    </div>
  );
};

export default ImageUploader;
