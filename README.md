# gin-vue-login

A simple full-stack authentication demo using:
- **Backend**: Gin (Go), MongoDB, JWT
- **Frontend**: Vue 3, Vite, Pinia, Axios, Tailwind CSS

This project provides basic auth flow:
- Register user
- Login user
- Store JWT token on frontend
- Access protected profile endpoint

## How to run

### Backend (`gin-login`)
1. `cd gin-login`
2. Make sure MongoDB is running on `mongodb://localhost:27017`
3. Configure `.env` (for example: `SECRET_KEY=your-secret`)
4. Run: `go run ./cmd/main.go`

Backend runs on `http://localhost:8000`.

### Frontend (`vue-login`)
1. `cd vue-login`
2. Install dependencies: `npm install`
3. Run dev server: `npm run dev`

Frontend runs on `http://localhost:5173`.
