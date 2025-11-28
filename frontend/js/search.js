document.getElementById("search-button").addEventListener("click", async () => {
  const query = document.getElementById("search-input").value.trim();
  const resultsContainer = document.getElementById("results-container");
  const messagesDiv = document.getElementById("messages");

  if (!query) return;

  resultsContainer.innerHTML = "";
  messagesDiv.innerHTML = '<div class="spinner"></div> Searching...';

  try {
    const response = await fetch(`/books/search?q=${encodeURIComponent(query)}`,{
      headers: {
        "token": "authOK"
      }
    });
    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || "Search failed");
    }

    if (data.length === 0) {
      messagesDiv.textContent = "No results found.";
      return;
    }

    messagesDiv.textContent = "";
    data.forEach((book) => {
      const card = document.createElement("div");
      card.className = "card";

      card.innerHTML = `
        <h3>${book.name}</h3>
        <p><strong>Author:</strong> ${book.author}</p>
        <p><strong>Price:</strong> $${book.price}</p>
        <a href="${book.url}" target="_blank">Download PDF</a>
      `;
      resultsContainer.appendChild(card);
    });
  } catch (err) {
    messagesDiv.textContent = `Error: ${err.message}`;
  }
});
