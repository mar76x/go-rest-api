CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE branches (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE areas (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE departments (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

