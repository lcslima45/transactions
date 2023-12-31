PGDMP     $                    {            Teste    15.3    15.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    24576    Teste    DATABASE     ~   CREATE DATABASE "Teste" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Portuguese_Angola.1252';
    DROP DATABASE "Teste";
                postgres    false            �            1259    24578    article    TABLE     j   CREATE TABLE public.article (
    id integer NOT NULL,
    title character varying(255),
    body text
);
    DROP TABLE public.article;
       public         heap    postgres    false            �            1259    24577    article_id_seq    SEQUENCE     �   CREATE SEQUENCE public.article_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.article_id_seq;
       public          postgres    false    215                       0    0    article_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.article_id_seq OWNED BY public.article.id;
          public          postgres    false    214            �            1259    24587    transactions    TABLE     �   CREATE TABLE public.transactions (
    id integer NOT NULL,
    cardholder character varying(255),
    merchant character varying(255),
    acquirer character varying(255),
    brand character varying(255),
    issuer character varying(255)
);
     DROP TABLE public.transactions;
       public         heap    postgres    false            �            1259    24586    transactions_id_seq    SEQUENCE     �   CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.transactions_id_seq;
       public          postgres    false    217            	           0    0    transactions_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;
          public          postgres    false    216            j           2604    24581 
   article id    DEFAULT     h   ALTER TABLE ONLY public.article ALTER COLUMN id SET DEFAULT nextval('public.article_id_seq'::regclass);
 9   ALTER TABLE public.article ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215            k           2604    24590    transactions id    DEFAULT     r   ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);
 >   ALTER TABLE public.transactions ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    217    217            �          0    24578    article 
   TABLE DATA           2   COPY public.article (id, title, body) FROM stdin;
    public          postgres    false    215   y                 0    24587    transactions 
   TABLE DATA           Y   COPY public.transactions (id, cardholder, merchant, acquirer, brand, issuer) FROM stdin;
    public          postgres    false    217   "       
           0    0    article_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.article_id_seq', 1, false);
          public          postgres    false    214                       0    0    transactions_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.transactions_id_seq', 36, true);
          public          postgres    false    216            m           2606    24585    article article_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.article DROP CONSTRAINT article_pkey;
       public            postgres    false    215            o           2606    24594    transactions transactions_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.transactions DROP CONSTRAINT transactions_pkey;
       public            postgres    false    217            �   �   x��̻
1��:y�����ڈ��lk3���$2/oo�g�;����;J�<c���͕���������)d��#��@1�HA��&&�u*����n�	*\��1����!)��cF&k
�U*4�[12N�qlJ�BK~�G{�G����,=�         �  x��V�V�8]�_�ݬ&�c�HCc��3a�&�*#ɀ�������r1zfћ�����UŶ8��$�y�̜��V�Y[��j�Nו8VF�A��@ٹ��^�+�%#�hesJvġr��7� ��e��v$2��6v�| 7T�X���(�'ʹF^r�1%�5r��p[NK�'���i;���T�t��=.y/Nx��L����$��	�J3�9v�΀|���E�^N�Wf%/6�����7�đ�qk6�֋��%�s�#wۈ�.
C�SKȚ�#2�|���d@�B�}ae�Z�i��!��3� u�� ��I���P�BLƻ2��<KA�� 3r�:�Nz�INԟ�C��q,�VdC�$g8��6@i�?@$�PGD�s0���Q��Lϑ	[yH� �=���l�$�TL9� �\)	��k��E�t���	��˽:��9��O�[�Bq�*��/�
>�wa)�����i�C��Ԭ���-��9X�_�+�J\ŮY�7�z��^���=E�y9�:h�M�;(����@��k�į~yC�Eg�3�Yp��w7;t�^�t�Ce����h���[��o?�NV�]7�!/s��mX����{�L�=��Q�.t�1|��y\��	dYzp_�m$N�St�GV���+��T0 �8���*�����"6<�sFf�k
M�~#us*�!9�˗�MW�<㴌�����OҾ�Bm�rv��5奍/5�~�n����D��G��6ISq�j<>U�%�~ʑE��iG�RU�tV?u�cKM�g �������(I��Ud�@�f<���筸H�q�Z��C��@dLT���c�s�r�Z���$݅4����@�8g�݌���n�*��5.{q�S��֏���֊X"��׼v�S��=�Z� z��Ę��*P!��ܲ���y9O�}x�S�N]��H��F]���<b6���7���6�k��I�x�uc���i�F7��3V�}��\}߈��#�F�c�n��R�QҶ)�^�SE��w(�	3�}���3 <[F]*5��*/#�
���>.|{M�x<׵6H;oշ�Q���F2��������q���� �?!�KlXqN|����E��J�T'���$ɿ�&\�     