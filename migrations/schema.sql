--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Ubuntu 14.4-1.pgdg22.04+1)
-- Dumped by pg_dump version 14.4 (Ubuntu 14.4-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: imdb_movies; Type: TABLE; Schema: public; Owner: kullop
--

CREATE TABLE public.imdb_movies (
    id integer NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    poster character varying(255) DEFAULT ''::character varying NOT NULL,
    genre character varying(255) DEFAULT ''::character varying NOT NULL,
    year integer DEFAULT 0 NOT NULL,
    rating numeric DEFAULT '0'::numeric NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.imdb_movies OWNER TO kullop;

--
-- Name: imdb_movies_id_seq; Type: SEQUENCE; Schema: public; Owner: kullop
--

CREATE SEQUENCE public.imdb_movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.imdb_movies_id_seq OWNER TO kullop;

--
-- Name: imdb_movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kullop
--

ALTER SEQUENCE public.imdb_movies_id_seq OWNED BY public.imdb_movies.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: kullop
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO kullop;

--
-- Name: imdb_movies id; Type: DEFAULT; Schema: public; Owner: kullop
--

ALTER TABLE ONLY public.imdb_movies ALTER COLUMN id SET DEFAULT nextval('public.imdb_movies_id_seq'::regclass);


--
-- Name: imdb_movies imdb_movies_pkey; Type: CONSTRAINT; Schema: public; Owner: kullop
--

ALTER TABLE ONLY public.imdb_movies
    ADD CONSTRAINT imdb_movies_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: kullop
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

