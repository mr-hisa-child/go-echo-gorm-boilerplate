--ログインロール作成
CREATE ROLE gouser WITH LOGIN PASSWORD 'gopass';

-- データベース作成
CREATE DATABASE godb OWNER gouser ENCODING 'UTF8';

\c godb

\connect - gouser