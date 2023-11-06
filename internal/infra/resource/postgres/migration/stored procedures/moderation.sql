create procedure savemovement(IN sp_employee_id integer, IN sp_action character varying, IN sp_previous text, IN sp_current text)
    language plpgsql
as
$$
begin
    IF sp_action = 'INSERT' THEN
        INSERT INTO moderation_logs(triggered_by, action, previous_value, current_value, created_at) VALUES(sp_employee_id, sp_action, CAST(sp_previous AS TEXT), CAST( sp_current AS TEXT), CURRENT_TIMESTAMP);
    END IF;

    IF sp_action = 'DELETE' THEN
        INSERT INTO moderation_logs(triggered_by, action, previous_value, current_value, created_at) VALUES(sp_employee_id, sp_action, CAST(sp_previous AS TEXT), CAST( sp_current AS TEXT), CURRENT_TIMESTAMP);
    END IF;

    IF sp_action = 'UPDATE' THEN
        INSERT INTO moderation_logs(triggered_by, action, previous_value, current_value, created_at) VALUES(sp_employee_id, sp_action, CAST(sp_previous AS TEXT), CAST( sp_current AS TEXT), CURRENT_TIMESTAMP);
    END IF;

    commit;
end;
$$;
