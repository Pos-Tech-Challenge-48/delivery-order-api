-- CATEGORY
-- Tabela que contém as referências das categorias dos produtos.

-- TABLE

CREATE TABLE IF NOT EXISTS category (
    category_id UUID DEFAULT gen_random_uuid()
        constraint category_pk primary key,
    category_name varchar(100) not null
);

-- INDEXES

CREATE UNIQUE INDEX category_name_idx on category(category_name);