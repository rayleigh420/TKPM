import React from "react";
import axios from "axios";

// PdfDownloader is a component that allows the user to input and download a pdf
function PdfDownloader() {
  const handleClick = () => {
    const url = `http://localhost:8080/inphieumuon/WuoItYMArWFD`;

    axios
      .get(url, { responseType: "blob" })
      .then((res) => {
        const blob = new Blob([res.data], { type: "application/pdf" });

        const link = document.createElement("a");
        link.href = URL.createObjectURL(blob);
        link.download = ".pdf";

        document.body.appendChild(link);
        link.click();

        // Remove the link from the document body
        document.body.removeChild(link);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <div>
      <h1>Tải về PDF</h1>
      <br />
      <button onClick={handleClick}>Tải về PDF</button>
    </div>
  );
}

export default PdfDownloader;
