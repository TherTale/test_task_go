CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE operators (
                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           first_name VARCHAR(255) NOT NULL,
                           last_name VARCHAR(255) NOT NULL,
                           middle_name VARCHAR(255),
                           city VARCHAR(255),
                           phone_number VARCHAR(15) UNIQUE NOT NULL,
                           email VARCHAR(255) UNIQUE NOT NULL,
                           password VARCHAR(255) NOT NULL
);

CREATE TABLE projects (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          name VARCHAR(255) NOT NULL,
                          type VARCHAR(50) NOT NULL
);

CREATE TABLE project_assignments (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
                                     operator_id UUID REFERENCES operators(id) ON DELETE CASCADE
);
