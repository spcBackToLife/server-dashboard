# Wallet Management Application

This project is a personal/family wallet management application designed to help track finances, manage expenses, and potentially handle loan repayments. It features a React (TypeScript) frontend and a Go backend.

## Prerequisites

Before you begin, ensure you have the following installed:

*   **Node.js:** Version 18.x or later (for the frontend). You can download it from [nodejs.org](https://nodejs.org/).
*   **Go:** Version 1.20 or later (for the backend). You can download it from [go.dev](https://go.dev/dl/).
*   **npm:** (Usually comes with Node.js) Version 8.x or later.

## Project Structure

\`\`\`
/
├── backend/         # Go backend source code
├── frontend/        # React + Vite + TailwindCSS frontend source code
├── product_features.md
├── ui_design_concepts.md
└── README.md
\`\`\`

## Backend Setup (Go)

1.  **Navigate to the backend directory:**
    \`\`\`bash
    cd backend
    \`\`\`

2.  **Install dependencies:**
    Go modules are typically downloaded automatically when you build or run the project. If you need to explicitly tidy the modules:
    \`\`\`bash
    go mod tidy
    \`\`\`

3.  **Run the backend server:**
    \`\`\`bash
    go run main.go
    \`\`\`
    The backend server will start on \`http://localhost:8080\`.

### Backend Available Scripts

*   \`go run main.go\`: Starts the backend development server.
*   \`go test ./...\`: Runs all unit tests in the backend.
*   \`go build -o wallet-backend main.go\`: Compiles the backend application into an executable named \`wallet-backend\`.

## Frontend Setup (React + Vite)

1.  **Navigate to the frontend directory:**
    \`\`\`bash
    cd frontend
    \`\`\`

2.  **Install dependencies:**
    \`\`\`bash
    npm install
    \`\`\`

3.  **Run the frontend development server:**
    \`\`\`bash
    npm run dev
    \`\`\`
    The frontend development server will start, typically on \`http://localhost:5173\` (Vite will show the exact port). API requests from the frontend to \`/api/...\` will be proxied to the backend server running on \`http://localhost:8080\` (as configured in \`vite.config.ts\`).

### Frontend Available Scripts

*   \`npm run dev\`: Starts the frontend development server with Hot Module Replacement (HMR).
*   \`npm run build\`: Compiles the frontend application for production into the \`dist\` directory.
*   \`npm run preview\`: Serves the production build locally for preview.
*   \`npm run lint\`: (If ESLint is configured) Lints the codebase.

## Development Workflow

1.  Start the backend server from the \`backend\` directory.
2.  Start the frontend development server from the \`frontend\` directory.
3.  Open your browser to the address provided by the Vite dev server (usually \`http://localhost:5173\`).

EOL
