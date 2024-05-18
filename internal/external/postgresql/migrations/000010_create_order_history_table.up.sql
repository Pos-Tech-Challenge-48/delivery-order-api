-- ORDER_HISTORY
-- Tabela que contém o histórico de mudança dos status dos pedidos ao longo
-- de seu ciclo de vida.

-- TABLE

CREATE TABLE IF NOT EXISTS order_history (
    order_history_id UUID DEFAULT gen_random_uuid()
        constraint order_history_pk primary key,
    order_history_order_id UUID
        constraint order_history_order_id_fk
            references restaurant_order(restaurant_order_id),
    order_history_status_id UUID
        constraint order_history_status_id_fk
            references status(status_id),
    order_responsible_employee UUID
        constraint order_responsible_employee_fk
            references employee(employee_id),
    created_date_db timestamptz default now() not null
);


-- TRIGGERS

CREATE OR REPLACE TRIGGER tr_order_history_timestamps
    BEFORE UPDATE
    ON public.order_history
    FOR EACH ROW
EXECUTE PROCEDURE public.fn_update_last_modified_date_db();