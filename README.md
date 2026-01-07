# hl4b-portal
A homelab portal site for HL4B

## Overview

This is a web portal built with:
- **Go** - Backend language
- **Gofiber** - Web framework
- **Gomponents** - HTML components in Go
- **GORM** - ORM for database operations
- **htmx** - Dynamic HTML without JavaScript
- **Tailwind CSS** - Utility-first CSS framework
- **PostgreSQL** - Database

## Features

- **Authentication** - Login/logout functionality
- **User Dashboard** - Personal dashboard for users
- **File Uploads** - Users can upload and manage their files
- **Admin Area** - Administrators can manage users and all uploads
- **Role-based Access Control** - Admin vs regular user permissions

## Database Structure

The application uses two tables:

1. **auth.users** (external, pre-existing)
   - Contains authentication data
   - Primary key: `id` (UUID)
   - Contains: `email`, `created_at`, `updated_at`

2. **portalusers** (managed by this app)
   - Links to `auth.users` via `user_id` (UUID foreign key)
   - Contains: `username`, `is_admin`, etc.

3. **uploads** (managed by this app)
   - Stores file upload metadata
   - Links to `auth.users` via `user_id` (UUID foreign key)

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database with `auth.users` table
- Node.js and npm (for Tailwind CSS)

## Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/InvertedBit/hl4b-portal.git
   cd hl4b-portal
   ```

2. **Install Go dependencies**
   ```bash
   go mod download
   ```

3. **Install Node dependencies (for Tailwind CSS)**
   ```bash
   npm install
   ```

4. **Build Tailwind CSS**
   ```bash
   npm run build:css
   ```

5. **Configure environment**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` with your settings:
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
   SESSION_SECRET=your-secure-random-string
   PORT=3000
   UPLOADS_DIR=./uploads
   ```

6. **Ensure PostgreSQL database has auth.users table**
   ```sql
   CREATE SCHEMA IF NOT EXISTS auth;
   
   CREATE TABLE IF NOT EXISTS auth.users (
       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       email VARCHAR(255) UNIQUE NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

7. **Run the application**
   ```bash
   go run main.go
   ```

   The application will:
   - Connect to the database
   - Auto-migrate the `portalusers` and `uploads` tables
   - Start the web server on the configured port (default: 3000)

8. **Access the portal**
   - Open your browser to `http://localhost:3000`
   - Login with any email (for demo purposes, users are auto-created)
   - First user can be manually set as admin in the database

## Development

### Watch Tailwind CSS changes
```bash
npm run watch:css
```

### Run the application with auto-reload
```bash
# Install air for live reload
go install github.com/cosmtrek/air@latest

# Run with air
air
```

## Project Structure

```
hl4b-portal/
├── config/          # Configuration management
├── database/        # Database connection and migrations
├── handlers/        # HTTP handlers
├── middleware/      # Authentication and authorization middleware
├── models/          # Database models
├── static/          # Static files (CSS, JS)
├── templates/       # Gomponents templates
├── uploads/         # User uploaded files (created at runtime)
├── main.go          # Application entry point
├── go.mod           # Go dependencies
├── package.json     # Node dependencies (Tailwind)
└── .env             # Environment configuration
```

## Routes

### Public Routes
- `GET /` - Home page
- `GET /login` - Login page
- `POST /login` - Login handler
- `POST /logout` - Logout handler

### User Routes (requires authentication)
- `GET /user/dashboard` - User dashboard
- `GET /user/uploads` - User's uploads page
- `POST /user/upload` - Upload file
- `DELETE /user/upload/:id` - Delete own upload

### Admin Routes (requires admin role)
- `GET /admin/dashboard` - Admin dashboard with statistics
- `GET /admin/users` - Manage users
- `POST /admin/users/:id/role` - Update user role
- `DELETE /admin/users/:id` - Delete user
- `GET /admin/uploads` - View all uploads
- `DELETE /admin/uploads/:id` - Delete any upload

## Security Notes

- Change `SESSION_SECRET` in production
- The demo login creates users automatically - integrate with your auth system in production
- File uploads should have size limits and type validation in production
- Add HTTPS in production
- Consider adding CSRF protection
- Add rate limiting for login attempts

## License

MIT
