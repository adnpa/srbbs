-- public."user" definition

-- Drop table

-- DROP TABLE public."user";

CREATE TABLE public."user" (
	id bigserial NOT NULL,
	user_id int8 NOT NULL,
	username varchar(64) NOT NULL,
	"password" varchar(64) NOT NULL,
	email varchar(64) NULL,
	gender int2 DEFAULT 0 NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT user_pk PRIMARY KEY (id),
	CONSTRAINT user_unique UNIQUE (user_id),
	CONSTRAINT user_unique_1 UNIQUE (username)
);


-- public.community definition

-- Drop table

-- DROP TABLE public.community;

CREATE TABLE public.community (
	id serial4 NOT NULL,
	community_id int4 NOT NULL,
	community_name varchar(128) NOT NULL,
	introduction varchar(256) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	deleted_at timestamp NULL,
	CONSTRAINT community_pk PRIMARY KEY (id),
	CONSTRAINT community_unique UNIQUE (community_id),
	CONSTRAINT community_unique_1 UNIQUE (community_name)
);


-- public.post definition

-- Drop table

-- DROP TABLE public.post;

CREATE TABLE public.post (
	id bigserial NOT NULL,
	post_id int8 NOT NULL,
	title varchar(128) NOT NULL,
	"content" varchar(8192) NOT NULL,
	author_id int8 NOT NULL,
	community_id int4 NOT NULL,
	status int4 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	deleted_at timestamp NULL,
	CONSTRAINT post_pk PRIMARY KEY (id),
	CONSTRAINT post_unique UNIQUE (post_id)
);
CREATE INDEX post_author_id_idx ON public.post USING btree (author_id);
CREATE INDEX post_community_id_idx ON public.post USING btree (community_id);


-- public."comment" definition

-- Drop table

-- DROP TABLE public."comment";

CREATE TABLE public."comment" (
	id bigserial NOT NULL,
	comment_id int8 NOT NULL,
	"content" text NOT NULL,
	post_id int8 NOT NULL,
	author_id int8 NOT NULL,
	parent_id int8 NOT NULL,
	status int4 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	deleted_at timestamp NULL,
	CONSTRAINT comment_pk PRIMARY KEY (id),
	CONSTRAINT comment_unique UNIQUE (comment_id)
);
CREATE INDEX comment_author_id_idx ON public.comment USING btree (author_id);