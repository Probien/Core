CREATE PROCEDURE update_orders()
language plpgsql
AS 
$$
BEGIN

    /*update all orders where extension date is overdue today*/
	UPDATE pawn_orders SET status_id = 2 WHERE extension_date = CURRENT_DATE AND status_id <> 3;

	/*update all orders with order status lost, where its extension date > today*/
    UPDATE pawn_orders SET status_id = 4 WHERE extension_date > CURRENT_DATE; 

END
$$;