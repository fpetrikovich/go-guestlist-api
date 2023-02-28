DROP TABLE IF EXISTS `event_table`;
DROP TABLE IF EXISTS `guest`;
DROP TABLE IF EXISTS `seating`;
DROP VIEW IF EXISTS `seating_usage`;

CREATE TABLE `event_table` (
  `table_id` INT NOT NULL auto_increment, 
  `capacity` INT UNSIGNED,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`table_id`)
);

CREATE TABLE `guest` (
  `guest_id` INT NOT NULL auto_increment, 
  `name` CHAR(100) NOT NULL UNIQUE, 
  `entourage` INT UNSIGNED DEFAULT 0,
  `arrival_status` ENUM('not_arrived', 'arrived', 'left', 'rejected', 'allocate') DEFAULT 'not_arrived',
  `arrived_at` TIMESTAMP NULL DEFAULT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`guest_id`)
);

CREATE TABLE `seating` (
  `guest_id` INT NOT NULL,
  `table_id` INT NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`guest_id`),
  CONSTRAINT `FK_guest_id` FOREIGN KEY (`guest_id`) REFERENCES `guest` (`guest_id`) ON DELETE CASCADE,
  CONSTRAINT `FK_table_id` FOREIGN KEY (`table_id`) REFERENCES `event_table` (`table_id`) ON DELETE CASCADE
);

CREATE VIEW `seating_usage` AS (
  SELECT tab.table_id, 
         tab.capacity, 
         (tab.capacity - (IFNULL(SUM(filtered_guest.entourage), 0) + IFNULL(COUNT(filtered_guest.entourage), 0))) as `free_seats`
  FROM event_table as tab
  LEFT JOIN (
              SELECT guest.entourage as entourage, 
                     seating.table_id as table_id 
              FROM `guest` 
              JOIN `seating` ON guest.guest_id=seating.guest_id 
              WHERE FIELD(guest.arrival_status, "not_arrived", "arrived")
  ) as `filtered_guest` ON tab.table_id=filtered_guest.table_id
  GROUP BY tab.table_id
);