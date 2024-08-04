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
                                     project_id UUID NOT NULL,
                                     operator_id UUID NOT NULL,
                                     FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
                                     FOREIGN KEY (operator_id) REFERENCES operators(id) ON DELETE CASCADE,
                                     CONSTRAINT unique_project_operator UNIQUE (project_id, operator_id) -- Уникальное ограничение
);
