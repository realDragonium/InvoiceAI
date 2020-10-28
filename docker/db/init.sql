-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.3-beta1
-- PostgreSQL version: 13.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- object: public.user | type: TABLE --
-- DROP TABLE IF EXISTS public."user" CASCADE;
CREATE TABLE public.user (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	verified boolean DEFAULT TRUE,
	email text UNIQUE,
	password text,
	name text,
	CONSTRAINT user_pk PRIMARY KEY (id)
);

-- ddl-end --
ALTER TABLE public.user OWNER TO invoice_ai;
-- ddl-end --

-- object: public.company | type: TABLE --
-- DROP TABLE IF EXISTS public.company CASCADE;
CREATE TABLE public.company (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	name text NOT NULL,
	CONSTRAINT company_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.company OWNER TO invoice_ai;
-- ddl-end --

-- object: public.customer | type: TABLE --
-- DROP TABLE IF EXISTS public.customer CASCADE;
CREATE TABLE public.customer (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	name text NOT NULL,
	info json,
	id_company uuid NOT NULL,
	CONSTRAINT customer_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.customer OWNER TO invoice_ai;
-- ddl-end --

-- object: public.many_company_has_many_user | type: TABLE --
-- DROP TABLE IF EXISTS public.many_company_has_many_user CASCADE;
CREATE TABLE public.many_company_has_many_user (
	id_company uuid NOT NULL,
	id_user uuid NOT NULL,
	CONSTRAINT many_company_has_many_user_pk PRIMARY KEY (id_company,id_user)

);
-- ddl-end --

-- object: company_fk | type: CONSTRAINT --
-- ALTER TABLE public.many_company_has_many_user DROP CONSTRAINT IF EXISTS company_fk CASCADE;
ALTER TABLE public.many_company_has_many_user ADD CONSTRAINT company_fk FOREIGN KEY (id_company)
REFERENCES public.company (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: user_fk | type: CONSTRAINT --
-- ALTER TABLE public.many_company_has_many_user DROP CONSTRAINT IF EXISTS user_fk CASCADE;
ALTER TABLE public.many_company_has_many_user ADD CONSTRAINT user_fk FOREIGN KEY (id_user)
REFERENCES public."user" (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: company_fk | type: CONSTRAINT --
-- ALTER TABLE public.customer DROP CONSTRAINT IF EXISTS company_fk CASCADE;
ALTER TABLE public.customer ADD CONSTRAINT company_fk FOREIGN KEY (id_company)
REFERENCES public.company (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --


