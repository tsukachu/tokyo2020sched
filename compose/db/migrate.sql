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

CREATE TABLE place (
  id int4 PRIMARY KEY,
  name varchar(32) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER updated_at
  BEFORE UPDATE ON place
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

CREATE TABLE olympic_schedule (
  id int4 PRIMARY KEY,
  competition_id int4 references competition(id),
  classification_id int4 references classification(id),
  title varchar(24) NOT NULL,
  "begin" timestamptz NOT NULL,
  "end" timestamptz NOT NULL,
  place_id int4 references place(id),
  content varchar(40),
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER updated_at
  BEFORE UPDATE ON olympic_schedule
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
