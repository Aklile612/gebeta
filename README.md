```markdown
# 🍽️ Recipe Sharing Platform

A full-stack web application where users can discover, share, and purchase food recipes. Built with Hasura, Go, PostgreSQL, Nuxt 3, and integrated with Cloudinary for image storage and Chapa for payments.

---

## ✨ Features

### 🧑‍🍳 Users
- Sign up, log in (JWT-based authentication)
- Browse public and premium recipes
- Create, edit, and delete personal recipes
- Comment, like, bookmark, and rate others' recipes

### 🍱 Recipes
- Filter by category, preparation time, difficulty, and ingredients
- Upload multiple images (with featured image)
- Add dynamic steps and ingredients
- Support for free and premium (paid) recipes

### 🛒 Premium
- Purchase premium recipes securely with Chapa
- Access purchased recipes in your dashboard

---

## 🛠 Tech Stack

| Layer     | Technology                         |
|-----------|-------------------------------------|
| Frontend  | Nuxt 3, Vue 3, TailwindCSS, Pinia   |
| Backend   | Hasura + Go (custom GraphQL handlers) |
| Database  | PostgreSQL                          |
| Auth      | JWT (via Hasura permissions)        |
| Storage   | Cloudinary (for images)             |
| Payments  | Chapa (for ETH payments)            |
| DevOps    | Docker, Render                      |

---

## 📁 Project Structure

```

recipe-app/
├── frontend/         # Nuxt 3 project
├── backend/          # Go backend for Hasura Actions & Events
│   └── cmd/
│       └── server/   # main.go entry point
├── hasura/           # Hasura migrations & metadata
└── docker/           # Dockerfiles and deployment setup

````

---

## 🚀 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Go ≥ 1.24](https://golang.org/)
- [Node.js ≥ 18](https://nodejs.org/)
- [Hasura CLI](https://hasura.io/docs/latest/hasura-cli/)
- Cloudinary Account
- Chapa Account

---

### 🔧 Environment Variables

Create a `.env` file in both `backend/` and `frontend/` with the necessary variables.

#### Backend `.env`:

```env
HASURA_GRAPHQL_ADMIN_SECRET=your_admin_secret
HASURA_URL=https://your-hasura-instance.io/v1/graphql
CLOUDINARY_CLOUD_NAME=your_cloud
CLOUDINARY_API_KEY=your_key
CLOUDINARY_API_SECRET=your_secret
JWT_SECRET=your_jwt_secret
CHAPA_SECRET_KEY=your_chapa_key
````

---

## 🐳 Docker Setup

### Build & Run Backend

```bash
docker build -t recipe-backend -f backend/Dockerfile .
docker run -p 8080:8080 --env-file=backend/.env recipe-backend
```

### Start Hasura (if using Docker Compose)

```bash
docker-compose up -d hasura postgres
```

---

## 💰 Payment Integration

Chapa is integrated in the backend using Go. When a user clicks "Buy", the backend:

* Validates JWT token for user info
* Sends payment request to Chapa API
* Handles success/cancel callbacks via Hasura event trigger

---

## 📦 Deployment

### Frontend (Nuxt)

Deploy to platforms like **Vercel**, **Netlify**, or **Render**

### Backend (Go)

Deploy to **Render**, **Railway**, or **Fly.io**

Ensure you add all the required env vars in your dashboard.

---

## 📸 Screenshots

| Home                     | Recipe Detail | Create Recipe |
| ------------------------ | ------------- | ------------- |
 home page
| <img width="1352" height="621" alt="image" src="https://github.com/user-attachments/assets/4596e540-0855-4a54-a4ec-af7b66177605" />

<img width="1356" height="612" alt="image" src="https://github.com/user-attachments/assets/8e035a81-d643-4470-a393-50223c9fda25" />
  Login page
<img width="1366" height="619" alt="image" src="https://github.com/user-attachments/assets/b52b8fbf-1296-4b1c-bb44-86270964cc00" />
register page
<img width="1365" height="619" alt="image" src="https://github.com/user-attachments/assets/585a03a1-5483-4bd6-a89b-fca92b9e0e86" />

---

## 🤝 Contributing

Pull requests are welcome. For major changes, open an issue first to discuss what you would like to change.

---

## 📄 License

[MIT](LICENSE)

---

## ✍️ Author

**Aklile Ansa**
[GitHub](https://github.com/aklile) • [LinkedIn](https://linkedin.com/in/aklile-ansa)

```

