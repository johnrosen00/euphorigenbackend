create table if not exists passwords (
    tpass varchar(255) not null,
    ntpass varchar(255) not null
);

create table if not exists playersession (
    playerid int not null auto_increment primary key,
    lastpuzzleid int not null
);

create table if not exists events (
    eventid int not null auto_increment primary key,
    playersessionid int not null,
    puzzleid int not null,
    metrictypeid varchar(255) not null,
    timeinitiated datetime not null,
    logtext varchar(255)
);