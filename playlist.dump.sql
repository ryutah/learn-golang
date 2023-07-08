--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

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

--
-- Name: name; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.name AS text
	CONSTRAINT name_check CHECK ((length(VALUE) <= 50));


ALTER DOMAIN public.name OWNER TO postgres;

--
-- Name: short_text; Type: DOMAIN; Schema: public; Owner: postgres
--

CREATE DOMAIN public.short_text AS text
	CONSTRAINT short_text_check CHECK ((length(VALUE) <= 256));


ALTER DOMAIN public.short_text OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: playlist; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.playlist (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title public.short_text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.playlist OWNER TO postgres;

--
-- Name: TABLE playlist; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.playlist IS 'プレイリスト';


--
-- Name: COLUMN playlist.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.playlist.title IS 'タイトル';


--
-- Name: COLUMN playlist.user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.playlist.user_id IS 'ユーザID外部キー';


--
-- Name: playlist_song_stores; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.playlist_song_stores (
    playlists_id uuid NOT NULL,
    songs_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.playlist_song_stores OWNER TO postgres;

--
-- Name: TABLE playlist_song_stores; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.playlist_song_stores IS 'プレイリストに保存しているソング';


--
-- Name: songs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.songs (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title public.short_text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.songs OWNER TO postgres;

--
-- Name: TABLE songs; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.songs IS 'ソング';


--
-- Name: COLUMN songs.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.songs.title IS 'タイトル';


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name public.short_text NOT NULL,
    birthday date,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: TABLE users; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.users IS 'ユーザ';


--
-- Name: COLUMN users.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.users.name IS 'ユーザ名';


--
-- Name: COLUMN users.birthday; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.users.birthday IS '誕生日';


--
-- Name: COLUMN users.created_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.users.created_at IS 'データ作成日時';


--
-- Name: COLUMN users.updated_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.users.updated_at IS 'データ更新日時';


--
-- Data for Name: playlist; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.playlist (id, title, created_at, updated_at, user_id) FROM stdin;
\.


--
-- Data for Name: playlist_song_stores; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.playlist_song_stores (playlists_id, songs_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: songs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.songs (id, title, created_at, updated_at) FROM stdin;
61951d4b-ac47-4fb0-960f-933486014b73	ソング1	2023-06-18 05:18:38.063795+00	2023-06-18 05:18:38.063795+00
f96db4be-33a4-44da-99ba-2cd94093a36f	ソング2	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
bdeea59e-e8fe-426f-a8c1-ae306eace029	ソング3	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
ece4eeef-2ee5-4189-b310-3b17dd1da431	ソング4	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
77d69867-d0f1-4fd3-9574-ed2eb3821dbd	ソング5	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
12488a6e-7c50-447e-9f9d-29189586445f	ソング6	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
7b92f593-dbc4-4a36-b724-4cbd9f347ba6	ソング7	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
50edd792-9f53-4cc7-bc2a-ea9f1c7e790a	ソング8	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
0fa96701-3640-497e-9884-341a067aecb2	ソング9	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
2ce51241-a765-4b74-9529-1cccb283a37e	ソング10	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
754a6126-133d-402e-8f4a-bb599d358dab	ソング11	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
ab3b7bed-8711-4cbc-9a06-87051966537e	ソング12	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
b2f2f5f9-28ea-4560-bf02-7b397cbe10fc	ソング13	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
a911b02a-4129-49bd-968b-9c6ddb82df68	ソング14	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
ddd49d27-f168-460f-81b5-718bf413d134	ソング15	2023-06-18 05:19:22.189909+00	2023-06-18 05:19:22.189909+00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, birthday, created_at, updated_at) FROM stdin;
e874d694-421e-4bd8-89c7-1b2f173d9862	平田 哲	1970-08-26	2023-06-18 05:05:09.350521+00	2023-06-18 05:05:09.350521+00
00863110-dd46-44ab-b2ac-2c01a629b11c	井村 栄美	1982-01-08	2023-06-18 05:06:36.275332+00	2023-06-18 05:06:36.275332+00
3f0e5559-55d5-413a-9d56-18832f37e373	松尾 尚子	1979-12-09	2023-06-18 05:06:36.275332+00	2023-06-18 05:06:36.275332+00
34647ffe-f787-4530-ad02-cbc57af81627	古本 豊治	2006-03-01	2023-06-18 05:06:36.275332+00	2023-06-18 05:06:36.275332+00
0ed8e576-3e4b-4afb-b9d1-ab666df688e6	橋田 凛	1975-05-23	2023-06-18 05:06:36.275332+00	2023-06-18 05:06:36.275332+00
44bcd2d9-2e2f-495d-8f97-0717438ba2c6	奥山 葉菜	1993-05-24	2023-06-18 05:06:36.275332+00	2023-06-18 05:06:36.275332+00
4d906681-bbfb-4674-bdf1-1e91c5bf2814	仁科 唯衣	2003-04-12	2023-06-18 05:07:41.458484+00	2023-06-18 05:07:41.458484+00
38f02d3a-737e-444a-aacb-75398641d716	坂根 高志	1987-12-24	2023-06-18 05:07:41.458484+00	2023-06-18 05:07:41.458484+00
33412c01-6164-40f5-a66a-ddd47ec9017a	高木 研治	2009-08-30	2023-06-18 05:07:41.458484+00	2023-06-18 05:07:41.458484+00
ee39ab9b-d6f9-4cd9-a6da-1dfbedfd480f	泉 詩織	2000-12-13	2023-06-18 05:07:41.458484+00	2023-06-18 05:07:41.458484+00
\.


--
-- Name: playlist playlist_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlist
    ADD CONSTRAINT playlist_pkey PRIMARY KEY (id);


--
-- Name: playlist_song_stores playlist_song_stores_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlist_song_stores
    ADD CONSTRAINT playlist_song_stores_pkey PRIMARY KEY (playlists_id, songs_id);


--
-- Name: songs songs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: playlists_id_1687073098579_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX playlists_id_1687073098579_index ON public.playlist_song_stores USING btree (playlists_id);


--
-- Name: songs_id_1687073103018_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX songs_id_1687073103018_index ON public.playlist_song_stores USING btree (songs_id);


--
-- Name: user_id_1687069787196_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_id_1687069787196_index ON public.playlist USING btree (user_id);


--
-- Name: playlist_song_stores playlist_song_stores_playlists_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlist_song_stores
    ADD CONSTRAINT playlist_song_stores_playlists_id_fkey FOREIGN KEY (playlists_id) REFERENCES public.playlist(id) ON DELETE CASCADE;


--
-- Name: playlist_song_stores playlist_song_stores_songs_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlist_song_stores
    ADD CONSTRAINT playlist_song_stores_songs_id_fkey FOREIGN KEY (songs_id) REFERENCES public.songs(id) ON DELETE CASCADE;


--
-- Name: playlist playlist_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlist
    ADD CONSTRAINT playlist_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

