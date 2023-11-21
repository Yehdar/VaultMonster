import React, { useState } from "react";
import "./App.css";

const ENDPOINT = "http://localhost:4000";

function App() {
  const [file, setFile] = useState<File | null>(null);
  const [uploadedFile, setUploadedFile] = useState<string | null>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = event.target.files?.[0] || null;
    setFile(selectedFile);
  };

  const handleUpload = async () => {
    if (file) {
      const formData = new FormData();
      formData.append("file", file);

      try {
        const response = await fetch(`${ENDPOINT}/api/upload`, {
          method: "POST",
          body: formData,
        });

        if (response.ok) {
          const result = await response.json();
          setUploadedFile(result.fileName);
        } else {
          console.error("File upload failed");
        }
      } catch (error) {
        console.error("Error uploading file:", error);
      }
    }
  };

  const handleDownload = async () => {
    if (uploadedFile) {
      try {
        const response = await fetch(
          `${ENDPOINT}/api/download/${uploadedFile}`
        );
        if (response.ok) {
          const blob = await response.blob();
          const url = URL.createObjectURL(blob);
          const link = document.createElement("a");
          link.href = url;
          link.download = uploadedFile;
          document.body.appendChild(link);
          link.click();
          document.body.removeChild(link);
        } else {
          console.error("File download failed");
        }
      } catch (error) {
        console.error("Error downloading file:", error);
      }
    }
  };

  return (
    <div>
      <h1>Vault Monster</h1>
      <div>
        <input type="file" onChange={handleFileChange} />
        <button onClick={handleUpload}>Upload File</button>
      </div>
      {uploadedFile && (
        <div>
          <p>Uploaded File: {uploadedFile}</p>
          <button onClick={handleDownload}>Download File</button>
        </div>
      )}
    </div>
  );
}

export default App;
