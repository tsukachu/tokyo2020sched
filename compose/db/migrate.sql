CREATE FUNCTION set_updated_at() RETURNS TRIGGER AS $$
  BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TABLE competition (
  id int4 PRIMARY KEY,
  name varchar(16) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER updated_at
  BEFORE UPDATE ON competition
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

CREATE TABLE classification (
  id int4 PRIMARY KEY,
  name varchar(16) NOT NULL,
  competition_id int4 references competition(id),
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER updated_at
  BEFORE UPDATE ON classification
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
