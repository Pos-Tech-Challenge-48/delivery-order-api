-- CUSTOMER
-- Tabela que contém as referências dos status dos pedidos.

-- TABLE

CREATE TABLE IF NOT EXISTS status
(
    status_id UUID DEFAULT gen_random_uuid()
        constraint status_pk primary key,
    status_name varchar(30) not null,
    status_active bool,
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);

-- INDEXES

CREATE UNIQUE INDEX status_name_idx on status(status_name);

-- TRIGGERS

CREATE OR REPLACE TRIGGER tr_status_timestamps
    BEFORE UPDATE
    ON public.status
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();