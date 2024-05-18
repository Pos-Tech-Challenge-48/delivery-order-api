-- CUSTOMER
-- Tabela que contém dados dos clientes que realizam pedidos.

-- TABLE

CREATE TABLE IF NOT EXISTS customer
(
    customer_id UUID DEFAULT gen_random_uuid()
        constraint customer_pk primary key,
    customer_document varchar(14),
    customer_name varchar(150),
    customer_email varchar(80),
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);

-- INDEXES

CREATE INDEX IF NOT EXISTS customer_document_idx ON customer(customer_document);

CREATE INDEX IF NOT EXISTS customer_document_idx ON customer(customer_email);

-- TRIGGERS

CREATE OR REPLACE TRIGGER tr_customer_timestamps
    BEFORE UPDATE
    ON public.customer
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();
