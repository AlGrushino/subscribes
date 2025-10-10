CREATE TABLE subscribes (
	id SERIAL PRIMARY KEY,
	service_name VARCHAR(50) NOT NULL,
	price INTEGER NOT NULL CHECK (price > 0),
	user_id UUID NOT NULL,
	start_date TIMESTAMP NOT NULL,
	end_date TIMESTAMP,

	CONSTRAINT fk_subscribe_user
		FOREIGN KEY (user_id)
		REFERENCES users(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);