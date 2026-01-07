# Quick Start Guide

This guide will help you get the HL4B Portal up and running quickly using Docker Compose.

## Prerequisites

- Docker and Docker Compose installed
- That's it! Docker will handle everything else.

## Step 1: Clone the Repository

```bash
git clone https://github.com/InvertedBit/hl4b-portal.git
cd hl4b-portal
```

## Step 2: Start the Application

```bash
docker compose up -d
```

This will:
- Start a PostgreSQL database
- Create the `auth.users` table automatically
- Build and start the application
- The app will be available at http://localhost:3000

## Step 3: Access the Portal

Open your browser and navigate to:
```
http://localhost:3000
```

You should see the HL4B Portal login page.

## Step 4: Create Your First User

Login with any email address. The system will automatically create:
1. An entry in `auth.users` table
2. A corresponding entry in `portal_users` table

Example:
- Email: `admin@example.com`
- Password: (any password - this is for demo purposes)

## Step 5: Make a User Admin

To promote a user to administrator:

```bash
# Access the database
docker exec -it hl4b-portal-db psql -U hl4b -d hl4b

# Run this SQL (replace email with actual user email)
UPDATE portal_users 
SET is_admin = true 
WHERE user_id = (
    SELECT id FROM auth.users WHERE email = 'admin@example.com'
);

# Verify
SELECT pu.username, pu.is_admin, au.email
FROM portal_users pu
JOIN auth.users au ON pu.user_id = au.id;

# Exit
\q
```

Or use the provided SQL script:
```bash
# Edit set_admin.sql and change the email
docker exec -i hl4b-portal-db psql -U hl4b -d hl4b < set_admin.sql
```

## Step 6: Explore the Features

After logging in, you can:

### As a Regular User:
- View your dashboard
- Upload files
- Manage your uploads
- Download your files

### As an Administrator:
- All user features PLUS:
- Access admin dashboard
- View all users
- Promote/demote users to admin
- Delete users
- View all uploads across all users
- Delete any upload

## Stopping the Application

```bash
docker compose down
```

To remove all data (including uploaded files and database):
```bash
docker compose down -v
```

## Development Mode

If you want to develop and see live changes:

### Install Dependencies
```bash
go mod download
npm install
```

### Start Database Only
```bash
docker compose up -d postgres
```

### Configure Environment
```bash
cp .env.example .env
# Edit .env with your settings
```

### Build CSS
```bash
npm run build:css
# Or for watch mode:
npm run watch:css
```

### Run with Hot Reload (requires air)
```bash
go install github.com/cosmtrek/air@latest
air
```

Or manually:
```bash
go run main.go
```

## Troubleshooting

### Port Already in Use
If port 3000 or 5432 is already in use, edit `docker-compose.yml` and change the ports:
```yaml
ports:
  - "8080:3000"  # Change external port
```

### Database Connection Issues
Check if PostgreSQL is healthy:
```bash
docker compose ps
```

View logs:
```bash
docker compose logs postgres
docker compose logs app
```

### Reset Everything
```bash
docker compose down -v
docker compose up -d
```

## Next Steps

- Read the full [README.md](README.md) for detailed information
- Integrate with your existing authentication system
- Configure production environment variables
- Set up HTTPS for production use
- Add file size limits and validation
- Configure backup strategies for uploads and database
