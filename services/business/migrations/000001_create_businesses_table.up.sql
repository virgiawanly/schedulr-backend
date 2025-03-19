CREATE TABLE IF NOT EXISTS businesses (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  industry VARCHAR(255) NOT NULL,
  company_size VARCHAR(20),
  logo VARCHAR(255),
  email VARCHAR(255),
  phone VARCHAR(25),
  website VARCHAR(255),
  address_1 VARCHAR(255),
  address_2 VARCHAR(255),
  city VARCHAR(150),
  state VARCHAR(150),
  zipcode VARCHAR(20),
  country VARCHAR(100),
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_businesses_name ON businesses (name);
CREATE INDEX IF NOT EXISTS idx_businesses_email ON businesses (email);
CREATE INDEX IF NOT EXISTS idx_businesses_phone ON businesses (phone);