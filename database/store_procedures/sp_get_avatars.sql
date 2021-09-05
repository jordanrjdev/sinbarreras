CREATE DEFINER=`root`@`%` PROCEDURE `sp_get_avatars`()
BEGIN
	SELECT * FROM avatar; 
END