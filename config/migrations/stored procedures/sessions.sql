create procedure savesession(IN sp_employee_id integer)
    language plpgsql
as
$$
BEGIN
    INSERT INTO moderations(employee_id, created_at) VALUES(sp_employee_id, current_timestamp);
END;
$$;