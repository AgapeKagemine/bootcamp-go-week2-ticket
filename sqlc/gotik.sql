drop table if exists users, events, tickets, transaction_details, events_tickets, transaction_details_events_users cascade;

create table users (
	id BIGSERIAL primary key,
	username VARCHAR(255),
	phone VARCHAR(255),
	email VARCHAR(255),
	balance NUMERIC
);

create table tickets (
	id BIGSERIAL primary key,
	stock INT,
	type VARCHAR(255),
	price NUMERIC
);

create table events (
	id BIGSERIAL primary key,
	name VARCHAR(255),
	date VARCHAR(255),
	description VARCHAR(255),
	location VARCHAR(255)
);

create table transaction_details (
	id BIGSERIAL primary key,
	time VARCHAR(255),
	status VARCHAR(255),
	total_payment numeric
);

create table events_tickets (
	event_id BIGINT,
	ticket_id BIGINT,
	constraint fk_event_id
		foreign key (event_id)
			references events(id),
	constraint fk_ticket_id
		foreign key (ticket_id)
			references tickets(id)
);

create table transaction_details_events_users (
	transaction_detail_id BIGINT,
	event_id BIGINT,
	user_id BIGINT,
	constraint fk_td_id
		foreign key (transaction_detail_id)
			references transaction_details(id),
	constraint fk_event_id
		foreign key (event_id)
			references events(id),
	constraint fk_user_id
		foreign key (user_id)
			references users(id)
);
