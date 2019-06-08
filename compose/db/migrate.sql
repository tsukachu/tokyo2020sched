CREATE FUNCTION set_updated_at() RETURNS TRIGGER AS $$
  BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TABLE competitions (
  id int4 PRIMARY KEY,
  name varchar(16) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER updated_at
  BEFORE UPDATE ON competitions
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
