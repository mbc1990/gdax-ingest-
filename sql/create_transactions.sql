CREATE TABLE transactions(
  txn_id serial primary key not null,
  txn_side varchar(1024) not null,
  price decimal(12, 2) not null,
  timestamp timestamp
);
