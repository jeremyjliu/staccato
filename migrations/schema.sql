--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Debian 10.3-1.pgdg90+1)
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: notes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notes (
    id uuid NOT NULL,
    staccato_id uuid NOT NULL,
    user_id uuid NOT NULL,
    songs character varying[] NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.notes OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: staccatos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.staccatos (
    id uuid NOT NULL,
    open boolean NOT NULL,
    time_open timestamp without time zone NOT NULL,
    time_close timestamp without time zone NOT NULL,
    link text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.staccatos OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    provider character varying(255) NOT NULL,
    provider_id character varying(255) NOT NULL,
    last_login timestamp without time zone NOT NULL,
    administrator boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id);


--
-- Name: staccatos staccatos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.staccatos
    ADD CONSTRAINT staccatos_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX version_idx ON public.schema_migration USING btree (version);


--
-- Name: notes notes_staccatos_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_staccatos_id_fk FOREIGN KEY (staccato_id) REFERENCES public.staccatos(id);


--
-- Name: notes notes_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

