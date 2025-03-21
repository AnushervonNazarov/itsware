-- Active: 1739515049434@@localhost@5432@postgres@public


CREATE DATABASE company;

CREATE SCHEMA hr;

CREATE TABLE test(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name CITEXT
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT,
    body_search TSVECTOR
        GENERATED ALWAYS AS (to_tsvector('english', body)) STORED
);

INSERT INTO posts (title, body)
VALUES
    ('Introduction to PostgreSQL', 'This is an introductory post about PostgreSQL. It covers basic concepts and features.'),
    ('Advanced PostgresSQL Techniques', 'In this post, we delve into advanced PostgreSQL techniques for efficient querying and data manipulation.'),
    ('PostgreSQL Optimization Strategies', 'This post explores various strategies for optimizing PostgreSQL database performance and efficiency.');

SELECT to_tsvector('followers');
SELECT to_tsquery('followers');

INSERT INTO test(name) VALUES ('jack');

SELECT id FROM test;

CREATE EXTENSION "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS pgcrypto;

SELECT uuid_generate_v4();

SELECT * FROM gen_random_uuid();

CREATE EXTENSION citext;

DROP EXTENSION ;