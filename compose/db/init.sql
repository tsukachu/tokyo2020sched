CREATE ROLE user1 WITH LOGIN PASSWORD 'password123' CREATEDB;
ALTER ROLE user1 SET client_encoding TO 'UTF8';
ALTER ROLE user1 SET default_transaction_isolation TO 'read committed';
ALTER ROLE user1 SET timezone TO 'UTC';

CREATE DATABASE tokyo2020_sch OWNER user1;
