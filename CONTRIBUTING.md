# Contributing to HL4B Portal

Thank you for your interest in contributing to HL4B Portal! This document provides guidelines and instructions for contributing to the project.

## Development Setup

### Prerequisites
- Go 1.21 or higher
- Node.js 16+ and npm
- Docker and Docker Compose
- Git

### Setup Steps

1. **Fork and Clone**
   ```bash
   git clone https://github.com/YOUR_USERNAME/hl4b-portal.git
   cd hl4b-portal
   ```

2. **Install Dependencies**
   ```bash
   make install
   # Or manually:
   go mod download
   npm install
   ```

3. **Start PostgreSQL**
   ```bash
   docker compose up -d postgres
   ```

4. **Configure Environment**
   ```bash
   cp .env.example .env
   # Edit .env with your settings
   ```

5. **Build CSS**
   ```bash
   make css
   # Or for watch mode in a separate terminal:
   make watch-css
   ```

6. **Run the Application**
   ```bash
   make run
   # Or with hot reload:
   make dev
   ```

## Project Structure

```
hl4b-portal/
├── config/          # Application configuration
├── database/        # Database connection and migrations
├── handlers/        # HTTP request handlers
│   ├── handlers.go  # Public and user handlers
│   └── admin.go     # Admin-specific handlers
├── middleware/      # HTTP middleware (auth, admin checks)
├── models/          # Database models (GORM)
├── static/          # Static assets (CSS, JS)
│   └── css/
│       ├── input.css   # Tailwind source
│       └── output.css  # Generated (gitignored)
├── templates/       # Gomponents HTML templates
│   ├── components.go   # Reusable UI components
│   ├── pages.go        # Public pages
│   ├── user.go         # User pages
│   └── admin.go        # Admin pages
├── main.go          # Application entry point
└── README.md        # Project documentation
```

## Code Style

### Go Code
- Follow standard Go conventions
- Run `go fmt` before committing
- Use meaningful variable names
- Add comments for exported functions
- Keep functions focused and small

### Templates (Gomponents)
- Use semantic HTML
- Keep components small and reusable
- Use Tailwind classes for styling
- Follow existing naming conventions

### CSS (Tailwind)
- Prefer Tailwind utility classes
- Only add custom CSS when absolutely necessary
- Keep `input.css` minimal

## Making Changes

### 1. Create a Branch
```bash
git checkout -b feature/your-feature-name
```

### 2. Make Your Changes
- Write clean, readable code
- Follow the existing code style
- Add comments where necessary
- Test your changes locally

### 3. Test Your Changes
```bash
# Build the application
make build

# Run the application
make run

# Test manually at http://localhost:3000
```

### 4. Commit Your Changes
```bash
git add .
git commit -m "Add feature: description of your changes"
```

### 5. Push and Create PR
```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub.

## Commit Message Guidelines

- Use clear, descriptive commit messages
- Start with a verb (Add, Fix, Update, Remove, etc.)
- Keep the first line under 50 characters
- Add details in the body if needed

Examples:
```
Add user profile page
Fix login redirect issue
Update README with deployment instructions
Remove deprecated authentication method
```

## Areas for Contribution

### High Priority
- [ ] Add unit tests for handlers
- [ ] Add integration tests
- [ ] Improve error handling
- [ ] Add rate limiting for login attempts
- [ ] Add CSRF protection
- [ ] Add file type validation for uploads
- [ ] Add file size limits

### Features
- [ ] User profile editing
- [ ] Password reset functionality
- [ ] Email notifications
- [ ] Advanced file search
- [ ] File sharing between users
- [ ] Upload history and analytics
- [ ] Dark mode support

### Documentation
- [ ] API documentation
- [ ] Deployment guides (AWS, GCP, DigitalOcean)
- [ ] Security best practices
- [ ] Backup and restore procedures

### Infrastructure
- [ ] CI/CD pipeline
- [ ] Kubernetes manifests
- [ ] Terraform modules
- [ ] Monitoring and logging setup

## Testing

Currently, the project doesn't have a comprehensive test suite. Adding tests is a great way to contribute!

### Writing Tests
- Place tests in `*_test.go` files
- Use table-driven tests where appropriate
- Mock external dependencies
- Aim for good coverage of critical paths

Example test structure:
```go
func TestLoginHandler(t *testing.T) {
    // Setup
    app := fiber.New()
    // ... register routes
    
    // Test
    req := httptest.NewRequest("POST", "/login", nil)
    resp, _ := app.Test(req)
    
    // Assert
    if resp.StatusCode != 302 {
        t.Errorf("Expected 302, got %d", resp.StatusCode)
    }
}
```

## Code Review

All submissions require review. We use GitHub pull requests for this purpose.

### What We Look For
- Code quality and readability
- Adherence to project conventions
- Test coverage
- Documentation updates
- Security considerations

## Security

If you discover a security vulnerability, please email security@example.com instead of creating a public issue.

## Questions?

Feel free to open an issue for:
- Questions about the codebase
- Feature discussions
- Bug reports
- General feedback

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (MIT).
