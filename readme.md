# 🛠️ Environment Setup

## 📂 Backend – `app.env`

Create a file named `app.env` inside the `backend/` folder and add the environment variables given in example-env.md.

Use "go run main.go" or "air" command to run the server.

# frontend/.env

Create a file named .env and add this vaiable to it.

VITE_BACKEND_URL=http://localhost:8000

Use "npm run dev" to start dev server for frontend.

# Database

If not available you can use "docker-compose up -d" command to run a MySQL container. envs for this would be picked up from backend folder.