CREATE DEFINER=`root`@`%` PROCEDURE `sp_get_avatar_by_id`(IN id INT)
BEGIN
	SELECT * FROM avatar WHERE avatar_id = id; 
END