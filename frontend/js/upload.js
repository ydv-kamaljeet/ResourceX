document.getElementById("upload-form").addEventListener("submit", async (e) => {
  e.preventDefault();

  const statusDiv = document.getElementById("status");
  statusDiv.innerHTML = '<div class="spinner"></div> Uploading...';

  const formData = new FormData(e.target);

  try {
    const response = await fetch("/books/upload", {
      method: "POST",
      headers: {
        "token": "authOK"
      },
      body: formData,
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || "Failed to upload");
    }

    statusDiv.textContent = data.message || "Book uploaded successfully!";
    e.target.reset();
  } catch (err) {
    statusDiv.textContent = `Error: ${err.message}`;
  }
});
