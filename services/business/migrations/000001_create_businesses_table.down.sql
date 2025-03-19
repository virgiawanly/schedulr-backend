-- Drop the indexes
DROP INDEX IF EXISTS idx_businesses_name;
DROP INDEX IF EXISTS idx_businesses_email;
DROP INDEX IF EXISTS idx_businesses_phone;

-- Drop the table
DROP TABLE IF EXISTS businesses;