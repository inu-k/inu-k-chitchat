drop table posts;
drop table threads;
drop table sessions;
drop table users;


create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null       
);

create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null  
);

INSERT INTO users (id, uuid, name, email, password, created_at) VALUES (1, '1c0fbae0-fa0f-4220-50ee-679ec3fd25cd', 'test', 'a@a', '7e240de74fb1ed08fa08d38063f6a6a91462a815', '2023-09-18 01:16:53.918359');
INSERT INTO threads (id, uuid, topic, user_id, created_at) VALUES (1, 'bff2553a-5448-4ef9-5ea8-e29580f4df10', 'test topic', 1, '2023-09-18 01:18:17.272956');
INSERT INTO posts (id, uuid, body, user_id, thread_id, created_at) VALUES (1, '660e86d6-89d1-4cc5-5bc8-a77f8b4cd6ac' , 'first post!', 1, 1 , '2023-09-18 01:18:30.870423');
INSERT INTO posts (id, uuid, body, user_id, thread_id, created_at) VALUES (2, '0f1a6f03-acc0-4a9c-4ea3-8affbeda4ad5' , 'second post', 1, 1 , '2023-09-26 12:17:44.395666');
