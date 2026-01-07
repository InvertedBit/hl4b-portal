-- Script to set a user as administrator
-- Replace the email with the actual user's email

-- First, find the user's UUID from auth.users
-- Then update the portal user to be an admin

-- Example: Make user with email 'admin@example.com' an admin
UPDATE portalusers 
SET is_admin = true 
WHERE user_id = (
    SELECT id FROM auth.users WHERE email = 'admin@example.com'
);

-- Verify the change
SELECT 
    pu.id,
    pu.username,
    pu.is_admin,
    au.email
FROM portalusers pu
JOIN auth.users au ON pu.user_id = au.id
WHERE pu.is_admin = true;
