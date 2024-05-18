-- ORDER_ITEM
-- Tabela que contém dados dos funcionários do restaurantes e parceiros.

-- TABLE

CREATE TABLE IF NOT EXISTS employee (
    employee_id UUID DEFAULT gen_random_uuid()
        constraint employee_pk primary key,
    employee_name varchar(150) not null,
    employee_responsibility varchar(50) not null,
    employee_active bool,
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);

-- TRIGGERS

CREATE OR REPLACE TRIGGER tr_employee_timestamps
    BEFORE UPDATE
    ON public.employee
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();