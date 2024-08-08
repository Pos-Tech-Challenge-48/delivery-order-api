-- CUSTOMER_DATA_REMOVAL_REQUESTS
-- Tabela que contém os dados de pedidos de remocão de dados dos clientes.

-- TABLE

CREATE TABLE IF NOT EXISTS customer_data_removal_requests (
    request_id UUID DEFAULT gen_random_uuid(),
    requester_name varchar(200) not null,
    requester_address varchar(300) not null,
    requester_phonenumber varchar(20) not null,
    created_date_db timestamptz default now() not null,
    last_modified_date_db timestamptz default now() not null
);

-- TRIGGERS

CREATE OR REPLACE TRIGGER tr_customer_data_removal_requests_timestamps
    BEFORE UPDATE
    ON public.customer_data_removal_requests
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();