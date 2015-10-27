ALTER TABLE `profile`
	ADD COLUMN `access_token` VARCHAR(64) NULL DEFAULT NULL AFTER `is_blocked`,
	ADD COLUMN `access_token_datetime` INT(11) NULL DEFAULT NULL AFTER `access_token`;