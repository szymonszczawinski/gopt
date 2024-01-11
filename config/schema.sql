CREATE TABLE lifecyclestate (
	id bigserial NOT NULL,
	name varchar NULL,
	CONSTRAINT lifecyclestate_pkey PRIMARY KEY (id)
);

CREATE TABLE lifecycle (
	id bigserial NOT NULL,
	name varchar NULL,
	start_state_id int8 NULL,
	CONSTRAINT lifecycle_pkey PRIMARY KEY (id)
);
CREATE TABLE project (
	id bigserial NOT NULL,
	created timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	name varchar NULL,
	project_key varchar NULL,
	description varchar NULL,
	state_id int8 NULL,
	lifecycle_id int8 NULL,
	created_by_id int8 NULL,
	owner_id int8 NULL,
	CONSTRAINT project_pkey PRIMARY KEY (id)
);

CREATE TABLE statetransition (
	lifecycle_id int8 NULL,
	from_state_id int8 NULL,
	to_state_id int8 NULL
);

CREATE TABLE users (
	id bigserial NOT NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	email varchar NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
