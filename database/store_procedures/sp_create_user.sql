CREATE DEFINER=`root`@`%` PROCEDURE `sp_create_user`(in username varchar(50), in avatar_id int, out error varchar(255))
BEGIN
DECLARE EXIT HANDLER FOR SQLEXCEPTION
	BEGIN
		set error = 'ERROR';
        RESIGNAL;
	END;
    
    START TRANSACTION; 
	insert into `funsiba`.`user`(`username`, `avatar_id`) values(username, avatar_id); 
    set error = 'OK'; 
	COMMIT; 
END