📚 ResourceX — Book Upload & Search Platform

ResourceX is a simple, full-stack project that allows users to upload book PDFs and search for books stored in a cloud system.
It uses Go (Gin) for the backend, Supabase PostgreSQL for the database, Azure Blob Storage for file storage, and a clean HTML/CSS/JS frontend hosted separately.

🚀 Features
✔ Upload Books

Users can upload a PDF file along with metadata:

Book Name

Author

Price

The PDF is stored in Azure Blob Storage.

Only the blob file name is stored in the database.

✔ Search Books

Users can search by book name.

Matching records are fetched from PostgreSQL.

A fresh SAS URL is generated for each PDF so users can download securely.

✔ Modern UI

Neon-themed futuristic look

Responsive upload/search pages

Hosted via GitHub Pages

🏗 Project Structure
ResourceX
│
├── cmd/app/main.go              # App entry point
├── internal/
│   ├── db/db.go                 # Database connection (Supabase PostgreSQL)
│   ├── handlers/handlers.go     # Upload & Search logic
│   ├── models/store.go          # Book model
│   ├── routes/routes.go         # Route definitions
│   └── storage/azure.go         # Azure Blob upload + SAS generation
│
├── frontend/                    # Frontend hosted via GitHub Pages
│   ├── index.html
│   ├── upload.html
│   ├── search.html
│   ├── css/style.css
│   └── js/
│       ├── upload.js
│       └── search.js
│
└── README.md

🗄 Database Schema (Supabase PostgreSQL)

Table: books

Column	Type	Description
id	bigserial	Primary key
name	text	Book name
author	text	Author name
price	int	Price of the book
file_url	text	Blob file name in Azure
created_at	timestamptz	Auto-generated
☁ Azure Storage (Blob)

When a book is uploaded:

The PDF is uploaded to Azure Blob Storage inside container books

Filename is transformed to:

<originalName>_<timestamp>.pdf


When searching:

A 24-hour SAS URL is generated so users can download the PDF securely.

🌐 Backend Deployment (Render)

The backend is deployed on Render at:

https://resourcex-tcem.onrender.com


Important routes:

➤ Upload book

POST /books/upload
Form-data fields:

name: string  
author: string  
price: number  
file: PDF file

➤ Search books

GET /books/search?q=bookName

🎨 Frontend Deployment (GitHub Pages)

Frontend repo:

https://github.com/Kamaljeet-01/ResourceX-Frontend


Live site:

https://kamaljeet-01.github.io/ResourceX-Frontend


Frontend is connected to backend via:

const API_BASE = "https://resourcex-tcem.onrender.com";

⚠ Notes

Only a few books exist in the database currently.
More will appear once users upload new books.

The project uses free-tier hosting, so loading may take a few seconds.

🛠 Tech Stack
Backend

Go (Gin)

Supabase PostgreSQL

Azure Blob Storage

Render (deployment)

Frontend

HTML, CSS (neon theme), JavaScript

GitHub Pages (hosting)

🧪 How to Run Locally
1. Clone backend repository
git clone https://github.com/Kamaljeet-01/ResourceX.git

2. Set environment variables

Create .env:

DB_USER=your-user
DB_PASSWORD=your-password
DB_HOST=your-host
DB_PORT=5432 or your pooler port
DB_NAME=postgres
SSLMODE=require

AZURE_STORAGE_ACCOUNT=xxx
AZURE_STORAGE_KEY=xxx
AZURE_STORAGE_CONTAINER=books

3. Run backend
go mod tidy
go run cmd/app/main.go

4. Open frontend

Open index.html from:

ResourceX-Frontend/

📌 Future Improvements

User authentication

Categories / tags

Book previews

Bulk upload

Admin dashboard
