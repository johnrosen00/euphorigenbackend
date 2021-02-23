create table if not exists users (
    passhash binary(72) not null
);

create table if not exists management (
    passhash binary(72) not null,
    ntpasshash binary(72) not null
);

create table if not exists playersession (
    playersessionid int not null auto_increment primary key,
    lastpuzzleid int not null
);

create table if not exists events (
    eventid int not null auto_increment primary key,
    playersessionid int not null,
    puzzleid int not null,
    metrictypeid int not null,
    timeinitiated datetime not null,
    logtext varchar(255)
);

create table if not exists metrictype (
    metrictypeid int not null auto_increment primary key,
    metrictypename varchar(255) not null
);