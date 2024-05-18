-- PUBLIC FUNCTIONS AND TRIGGERS

CREATE OR REPLACE FUNCTION public.fn_update_last_modified_date_db()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS $$
    BEGIN
        NEW.last_modified_date_db = NOW();
        RETURN NEW;
    END;
    $$;

