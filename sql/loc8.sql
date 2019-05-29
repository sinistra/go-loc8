create table if not exists "__EFMigrationsHistory"
(
	"MigrationId" varchar(150) not null
		constraint "PK___EFMigrationsHistory"
			primary key,
	"ProductVersion" varchar(32) not null
);

alter table "__EFMigrationsHistory" owner to postgres;

create table if not exists accounts
(
	id serial not null
		constraint accounts_pkey
			primary key,
	name varchar(150),
	salesrep varchar(20),
	phone varchar(15),
	fax varchar(15),
	home_page varchar(50),
	parent_id integer,
	num_of_employees integer,
	territory_id varchar(10),
	account_manager varchar(50),
	account_manager_support_rep varchar(50),
	abn varchar(15),
	acn varchar(15),
	status varchar(15),
	do_not_contact varchar(1),
	reason varchar(50),
	external_handover_verified integer,
	external_handover_verified_dt timestamp,
	handover_account_manager varchar(50),
	handover_salesrep varchar(50),
	gtm_segment varchar(50),
	ops_segment varchar(50),
	pm_handover_verified integer,
	pm_handover_verified_dt timestamp,
	preferred_csp_rep varchar(50),
	salesforce_id varchar(20),
	editable_in_crisp integer,
	lag_customer integer,
	lag_entity_type varchar(20),
	lag_parent varchar(20),
	lag_contract varchar(20),
	limit_macview_actions integer,
	account_contract_entity varchar(20),
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table accounts owner to postgres;

create index if not exists accounts__updated_index
	on accounts (row_updated_at);

create table if not exists loc8ns
(
	id varchar(50) not null
		constraint loc8ns_pkey
			primary key,
	site_name varchar(255),
	formatted_address_string varchar(255),
	base_address varchar(255),
	sub_address varchar(255),
	address_source varchar(10),
	address_source_id varchar(25),
	street_number varchar(10),
	route varchar(100),
	locality varchar(50),
	state varchar(10),
	postcode varchar(10),
	sub_address_prefix varchar(100),
	sub_address_suffix varchar(100),
	complex_name varchar(100),
	latitude numeric(11,8),
	longitude numeric(11,8),
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP,
	lot_number varchar(20),
	road_name varchar(100),
	road_type_code varchar(10),
	road_suffix_code varchar(5),
	unit_number varchar(10),
	unit_type_code varchar(10),
	level_type_code varchar(50),
	level_number varchar(10),
	distancetoexchange integer,
	exchangecode varchar(50)
);

alter table loc8ns owner to postgres;

create table if not exists nbn_rtc
(
	nbn_location_id varchar(15)
		constraint "nbn_rtc_﻿NBN_LOCATION_ID_uindex"
			unique,
	csa_id varchar(15),
	hfl_effective_date timestamp,
	rollout_region_identifier varchar(30),
	first_serviceable_date date,
	most_recent_date_made_unserviceable date,
	most_recent_date_made_serviceable date
);

alter table nbn_rtc owner to postgres;

create table if not exists products
(
	id serial not null
		constraint products_pkey
			primary key,
	description varchar(50),
	sell_price numeric(10,2),
	cost_price numeric(10,2),
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP,
	prodgroupcode text,
	status integer,
	sunsetdate timestamp,
	startdate timestamp,
	hidefromuser text,
	salesreportcode text,
	bulkimportcode text,
	location text
);

alter table products owner to postgres;

create index if not exists products__updated_index
	on products (row_updated_at);

create table if not exists quote_item_interaction_log
(
	id serial not null
		constraint "PK_quote_item_interaction_log"
			primary key,
	subaddresses json,
	nbnproducts json,
	frontierlinkproducts json,
	buildingproducts json,
	exchangeproducts json,
	cirrusproducts json,
	anywhereproducts json,
	message_id varchar(50) not null,
	quoteitem_id integer not null,
	stepid integer not null,
	invalidatedbymessageid varchar(100)
);

alter table quote_item_interaction_log owner to postgres;

create table if not exists samples
(
	id integer
);

alter table samples owner to postgres;

create table if not exists serviceable_locations
(
	id integer not null
		constraint serviceable_locations_pk
			primary key,
	loc_id text,
	gnaf_pid text,
	latitude numeric(16,6),
	longitude numeric(16,6),
	connected_fibre integer,
	last_modified text,
	deleted integer,
	geometry_point text,
	created_at text,
	updated_at text
);

alter table serviceable_locations owner to postgres;

create table if not exists status
(
	id serial not null
		constraint "PK_status"
			primary key,
	status_id integer not null
		constraint "AK_status_status_id"
			unique,
	description varchar(100),
	stage varchar(50)
);

alter table status owner to postgres;

create table if not exists step
(
	id serial not null
		constraint "PK_step"
			primary key,
	step_id integer not null
		constraint "AK_step_step_id"
			unique,
	description varchar(100),
	status_id integer not null
		constraint step_status_id_fkey
			references status (status_id)
				on delete cascade
);

alter table step owner to postgres;

create index if not exists ix_step_status_id
	on step (status_id);

create table if not exists users
(
	id serial not null
		constraint user_pkey
			primary key,
	user_name varchar(20),
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP,
	user_id varchar(20)
		constraint users_user_id_uindex
			unique,
	token varchar(40)
		constraint users_token_uindex
			unique
);

alter table users owner to postgres;

create index if not exists users__updated_index
	on users (row_updated_at);

create table if not exists account_locations
(
	id serial not null
		constraint account_locations_pkey
			primary key,
	account_id integer
		constraint fk_account_locations_accounts
			references accounts,
	type varchar(20),
	address1 varchar(150),
	address2 varchar(150),
	city varchar(50),
	state varchar(20),
	country varchar(20),
	postcode integer,
	latitude numeric(16,6),
	longitude numeric(16,6),
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP,
	notes text,
	postal_code_id integer,
	attention text,
	change_uid text,
	change_dt text,
	sync_uid text,
	sync_dt text,
	bill_customer_id text,
	dpid text,
	form text,
	qualaddressinfo text,
	qualstreetname text,
	qualstreetno text,
	qualstreetnosuffix text,
	qualstreettype text,
	qualsubaddressno text,
	qualsubaddresstype text,
	qualsuburb text,
	isqualified integer,
	qualsiteid integer,
	localityabbrev text,
	exchangecode text,
	exchangedistance integer,
	addedby text,
	creationdate text,
	qualstreetsuffix text,
	msrepl_tran_version text,
	qualbuildingleveltype text,
	qualbuildinglevelno text,
	qualallotmentlot text,
	qualallotmentno text,
	nbnlocation text,
	telstralocation text
);

alter table account_locations owner to postgres;

create index if not exists account_locations__updated_index
	on account_locations (row_updated_at);

create index if not exists idx_account_locations_account_id
	on account_locations (account_id);

create table if not exists quotes
(
	id serial not null
		constraint quotes_pkey
			primary key,
	user_id integer
		constraint quotes_user_id_fkey
			references users,
	account_id integer
		constraint quotes_account_id_fkey
			references accounts,
	created_by integer not null
		constraint fk_quotes_users
			references users,
	description varchar(100),
	status varchar(20),
	step varchar(20),
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table quotes owner to postgres;

create index if not exists quotes__updated_index
	on quotes (row_updated_at);

create index if not exists idx_quotes_created_by
	on quotes (created_by);

create index if not exists ix_quotes_account_id
	on quotes (account_id);

create index if not exists ix_quotes_user_id
	on quotes (user_id);

create table if not exists quote_loc8ns
(
	id serial not null
		constraint quote_loc8ns_pkey
			primary key,
	quote_id integer not null
		constraint quote_loc8ns_quote_id_fkey
			references quotes
				on delete cascade,
	site_name varchar(255),
	formatted_address_string varchar(255),
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table quote_loc8ns owner to postgres;

create index if not exists quote_loc8ns__quote_id_index
	on quote_loc8ns (quote_id);

create index if not exists quote_loc8ns__sitename_index
	on quote_loc8ns (site_name);

create index if not exists quote_loc8ns__updated_index
	on quote_loc8ns (row_updated_at);

create table if not exists quote_items
(
	id serial not null
		constraint quote_items_pkey
			primary key,
	quote_loc8n_id integer not null
		constraint quote_items_quote_loc8n_id_fkey
			references quote_loc8ns
				on delete cascade,
	parent_id integer,
	product_id integer
		constraint quote_items_product_id_fkey
			references products,
	price numeric(10,2),
	cost numeric(10,2),
	quantity integer,
	site_name varchar(255),
	site_address varchar(255),
	loc8n_id varchar(50)
		constraint quote_items_loc8n_id_fkey
			references loc8ns,
	status varchar(20),
	parameters json,
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP,
	type varchar(50),
	provider varchar(50),
	matched_baseaddress_long_name varchar(255),
	matched_baseaddress_latitude numeric(11,8),
	matched_baseaddress_longitude numeric(11,8),
	matched_baseaddress_score integer,
	matched_baseaddress_msg varchar(255),
	matched_subaddress_long_name varchar(255),
	matched_subaddress_short_name varchar(255),
	matched_subaddress_latitude numeric(11,8),
	matched_subaddress_longitude numeric(11,8),
	matched_subaddress_score integer,
	matched_subaddress_msg varchar(255),
	matched_subaddress_service_class integer,
	matched_subaddress_tech_type varchar(10),
	matched_subaddress_source_name varchar(10),
	matched_subaddress_source_id varchar(50),
	matched_subaddress_mt_id varchar(50),
	preferred_speed varchar(50),
	step_id integer not null
		constraint quote_items_step_id_fkey
			references step (step_id)
				on delete cascade,
	products_picklist json,
	subaddress_picklist json,
	selected_product integer
);

alter table quote_items owner to postgres;

create index if not exists idx_quote_items_quote_loc8n_id
	on quote_items (quote_loc8n_id);

create index if not exists ix_quote_items_loc8n_id
	on quote_items (loc8n_id);

create index if not exists ix_quote_items_product_id
	on quote_items (product_id);

create index if not exists ix_quote_items_step_id
	on quote_items (step_id);

create table if not exists pfl
(
	id integer default nextval('pfl_id_seq'::regclass) not null
		constraint pfl_pkey
			primary key,
	nbn_locid varchar(25),
	gnaf_persistent_identifier varchar(20),
	rollout_region_identifier varchar(25),
	distribution_region_identifier varchar(25),
	formatted_address_string varchar(255),
	service_class integer,
	service_class_description varchar(150),
	service_class_reason varchar(100),
	ready_for_service_date varchar(10),
	disconnection_date varchar(10),
	unit_number varchar(10),
	unit_type_code varchar(10),
	level_number varchar(10),
	level_type_code varchar(3),
	address_site_name varchar(150),
	road_number_1 varchar(20),
	road_number_2 varchar(20),
	lot_number varchar(20),
	road_name varchar(100),
	road_suffix_code varchar(5),
	road_type_code varchar(5),
	locality_name varchar(50),
	secondary_complex_name varchar(255),
	complex_road_number_1 varchar(20),
	complex_road_number_2 varchar(20),
	complex_road_name varchar(50),
	complex_road_type_code varchar(20),
	complex_road_suffix_code varchar(20),
	postcode varchar(10),
	state_territory_code varchar(10),
	latitude numeric(11,8),
	longitude numeric(11,8),
	is_complex_premise_yn varchar(1),
	service_level_region varchar(15),
	service_type varchar(30),
	listing_type varchar(5),
	technology_type varchar(10),
	is_early_access_available varchar(15),
	poi_identifier varchar(5),
	poi_name varchar(30),
	transitional_poi_identifier varchar(5),
	transitional_poi_name varchar(30),
	connectivity_servicing_area_identifier varchar(30),
	connectivity_servicing_area_name varchar(50),
	transitional_connectivity_servicing_area_identifier varchar(30),
	transitional_connectivity_servicing_area_name varchar(50),
	new_develops_charge_applies varchar(15),
	delta_type varchar(15),
	last_updated_timestamp varchar(30),
	coexistence_status varchar(5),
	is_critical_service_flag char,
	row_created_at timestamp with time zone default CURRENT_TIMESTAMP,
	row_updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table pfl owner to postgres;

create index if not exists pfl__nbn_locid_index
	on pfl (nbn_locid);

create index if not exists pfl__updated_index
	on pfl (row_updated_at);

