-- SQL Script for setting up the auth.users table in PostgreSQL
-- This script should be run before starting the application

-- Create the auth schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS auth;

-- Create the auth.users table
-- This table is managed externally and should match your authentication system
CREATE TABLE IF NOT EXISTS auth.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on email for faster lookups
CREATE INDEX IF NOT EXISTS idx_auth_users_email ON auth.users(email);

-- Optional: Create a trigger to update updated_at automatically
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_auth_users_updated_at
    BEFORE UPDATE ON auth.users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Optional: Insert a sample user for testing (remove in production)
-- The application will create portal_users entry automatically on first login
INSERT INTO auth.users (id, email, created_at, updated_at)
VALUES 
    ('550e8400-e29b-41d4-a716-446655440000'::uuid, 'admin@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('550e8400-e29b-41d4-a716-446655440001'::uuid, 'user@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (email) DO NOTHING;

-- Note: The portalusers and uploads tables will be created automatically
-- by the application using GORM AutoMigrate when the application starts
