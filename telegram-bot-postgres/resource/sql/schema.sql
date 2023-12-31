\c postgres;

CREATE SCHEMA IF NOT EXISTS TELEGRAM_BOT;
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_user WHERE usename = 'root') THEN
            CREATE USER root WITH PASSWORD 'root';
        END IF;
    END
$$;
GRANT ALL PRIVILEGES ON SCHEMA TELEGRAM_BOT TO root;
set SEARCH_PATH = "telegram_bot";

