CREATE TABLE IF NOT EXISTS accounts (
  id UUID PRIMARY KEY,
  business_id UUID,
  email VARCHAR(255),
  password VARCHAR(255),
  role VARCHAR(15),
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
)