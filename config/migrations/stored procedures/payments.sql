create procedure savepayment(IN sp_employee_id integer, IN sp_customer_id integer, IN sp_payment numeric)
    language plpgsql
as
$$
BEGIN
    INSERT INTO payment_logs(employee_id, customer_id, payment, created_at) VALUES(sp_employee_id, sp_customer_id, sp_payment, current_timestamp);
    COMMIT;
END;
$$;