CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    company_id INT REFERENCES companies(id) ON DELETE CASCADE
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    department_id INT REFERENCES departments(id) ON DELETE SET NULL,
    passport_type VARCHAR(50),
    passport_number VARCHAR(50)
);