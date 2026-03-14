DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
-- TODO make user dynamic
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;
