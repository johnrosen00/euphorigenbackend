create table if not exists passwords (
    tpass varchar(255) not null,
    ntpass varchar(255) not null
);

create table if not exists playersession (
    playerid int not null auto_increment primary key,
    lastpuzzleid int not null
);

create table if not exists metrics (
    metricid int not null auto_increment primary key,
    playerid int not null,
    puzzleid int not null,
    timeinitiated datetime not null,
    metrictype varchar(255) not null,
    info varchar(255)
);