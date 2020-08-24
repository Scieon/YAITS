Use yaits;

Create table `issues` (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
summary varchar (64),
description varchar(256),
priority int not null default 1,
status ENUM('open', 'in progress', 'closed') not null default 'open',
assignee varchar(64) not null default 'unassigned',
reporter varchar(64),
createDate timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
constraint `priority_range` check (priority > 0 and priority < 11)
);
