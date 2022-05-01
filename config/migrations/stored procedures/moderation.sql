create procedure savemovement(IN sp_employee_id integer, IN sp_action character varying, IN sp_previous character varying, IN sp_current character varying)
    language plpgsql
as
$$
begin
    IF sp_action = 'CREATE' THEN
        INSERT INTO moderations(employee_id, action, previous_value, current_value, created_at) VALUES(sp_employee_id, sp_action, "no previous data registered", sp_current, current_timestamp);
    END IF;

    IF sp_action = 'DELETE' THEN
        INSERT INTO moderations(employee_id, action, previous_value, current_value, created_at) VALUES(sp_employee_id, sp_action, sp_previous, sp_current, current_timestamp);
    END IF;

    IF sp_action = 'UPDATE' THEN
        INSERT INTO moderations(employee_id, action, previous_value, current_value, created_at) VALUES(sp_employee_id, sp_action, sp_previous, sp_current, current_timestamp);
    END IF;

    commit;
end;
$$;
