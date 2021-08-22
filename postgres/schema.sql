
CREATE TABLE car_user (
  id int(11) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  account_id INT(11) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE car_user_account (
  id int(11) PRIMARY KEY,
  username VARCHAR(32) NOT NULL,
  password VARCHAR(32) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
