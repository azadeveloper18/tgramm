-- Modified by Kgramm, 2026-10-31.
-- Derived from teamgram-server SQL, which remains under the Apache License 2.0.
-- MySQL dialect (ENGINE, AUTO_INCREMENT counters, charset, ON DUPLICATE KEY) is not preserved.

CREATE TABLE IF NOT EXISTS auths (
  id bigint NOT NULL,
  auth_key_id bigint NOT NULL,
  layer integer NOT NULL DEFAULT 0,
  api_id integer NOT NULL,
  device_model varchar(255) NOT NULL DEFAULT '',
  system_version varchar(255) NOT NULL DEFAULT '',
  app_version varchar(255) NOT NULL DEFAULT '',
  system_lang_code varchar(255) NOT NULL DEFAULT '',
  lang_pack varchar(255) NOT NULL DEFAULT '',
  lang_code varchar(255) NOT NULL DEFAULT '',
  system_code varchar(256) NOT NULL DEFAULT '',
  proxy varchar(512) NOT NULL DEFAULT '',
  params jsonb NOT NULL,
  client_ip varchar(32) NOT NULL DEFAULT '',
  date_active bigint NOT NULL DEFAULT 0,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS auth_keys (
  id bigint NOT NULL,
  auth_key_id bigint NOT NULL,
  body varchar(512) NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS auth_key_infos (
  id bigint NOT NULL,
  auth_key_id bigint NOT NULL,
  auth_key_type integer NOT NULL,
  perm_auth_key_id bigint NOT NULL DEFAULT 0,
  temp_auth_key_id bigint NOT NULL DEFAULT 0,
  media_temp_auth_key_id bigint NOT NULL DEFAULT 0,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS auth_seq_updates (
  id bigint NOT NULL,
  auth_id bigint NOT NULL,
  user_id bigint NOT NULL,
  seq integer NOT NULL DEFAULT 0,
  update_type integer NOT NULL DEFAULT 0,
  update_data jsonb NOT NULL,
  date2 bigint NOT NULL DEFAULT 0,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS auth_users (
  id bigint NOT NULL,
  auth_key_id bigint NOT NULL,
  user_id bigint NOT NULL DEFAULT 0,
  hash bigint NOT NULL DEFAULT 0,
  layer integer NOT NULL DEFAULT 0,
  device_model varchar(128) NOT NULL DEFAULT '',
  platform varchar(64) NOT NULL DEFAULT '',
  system_version varchar(64) NOT NULL DEFAULT '',
  api_id integer NOT NULL DEFAULT 0,
  app_name varchar(64) NOT NULL DEFAULT '',
  app_version varchar(64) NOT NULL DEFAULT '',
  date_created bigint NOT NULL DEFAULT 0,
  date_actived bigint NOT NULL DEFAULT 0,
  ip varchar(64) NOT NULL DEFAULT '',
  country varchar(64) NOT NULL DEFAULT '',
  region varchar(64) NOT NULL DEFAULT '',
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS bots (
  id bigint NOT NULL,
  bot_id bigint NOT NULL,
  bot_type integer NOT NULL DEFAULT 0,
  creator_user_id bigint NOT NULL DEFAULT 0,
  token varchar(128) NOT NULL DEFAULT '',
  description varchar(10240) NOT NULL DEFAULT '',
  bot_chat_history boolean NOT NULL DEFAULT false,
  bot_nochats boolean NOT NULL DEFAULT true,
  verified boolean NOT NULL DEFAULT false,
  bot_inline_geo boolean NOT NULL DEFAULT false,
  bot_info_version integer NOT NULL DEFAULT 1,
  bot_inline_placeholder varchar(128) NOT NULL DEFAULT '',
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS bot_commands (
  id bigint NOT NULL,
  bot_id bigint NOT NULL,
  command varchar(128) NOT NULL DEFAULT '',
  description varchar(10240) NOT NULL DEFAULT '',
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS chats (
  id bigint NOT NULL,
  creator_user_id bigint NOT NULL,
  access_hash bigint NOT NULL,
  random_id bigint NOT NULL,
  participant_count integer NOT NULL,
  title varchar(255) NOT NULL,
  about varchar(255) NOT NULL DEFAULT '',
  photo_id bigint NOT NULL DEFAULT 0,
  default_banned_rights bigint NOT NULL DEFAULT 0,
  migrated_to_id bigint NOT NULL DEFAULT 0,
  migrated_to_access_hash bigint NOT NULL DEFAULT 0,
  deactivated boolean NOT NULL DEFAULT false,
  version integer NOT NULL DEFAULT 1,
  "date" bigint NOT NULL DEFAULT 0,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS chat_invites (
  id bigint NOT NULL,
  chat_id bigint NOT NULL,
  admin_id bigint NOT NULL,
  migrated_to_id bigint NOT NULL DEFAULT 0,
  link varchar(64) NOT NULL,
  permanent boolean NOT NULL DEFAULT false,
  revoked boolean NOT NULL DEFAULT false,
  request_needed boolean NOT NULL DEFAULT false,
  start_date bigint NOT NULL DEFAULT 0,
  expire_date bigint NOT NULL DEFAULT 0,
  usage_limit integer NOT NULL DEFAULT 0,
  usage2 integer NOT NULL DEFAULT 0,
  requested integer NOT NULL DEFAULT 0,
  title varchar(64) NOT NULL DEFAULT '',
  date2 bigint NOT NULL,
  state integer NOT NULL DEFAULT 0,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS chat_invite_participants (
  id bigint NOT NULL,
  link varchar(32) NOT NULL,
  user_id bigint NOT NULL,
  date2 bigint NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS chat_participants (
  id bigint NOT NULL,
  chat_id bigint NOT NULL,
  user_id bigint NOT NULL,
  participant_type integer DEFAULT 0,
  link varchar(64) NOT NULL DEFAULT '',
  usage2 integer NOT NULL DEFAULT 0,
  admin_rights integer NOT NULL DEFAULT 0,
  inviter_user_id bigint NOT NULL DEFAULT 0,
  invited_at bigint NOT NULL DEFAULT 0,
  kicked_at bigint NOT NULL DEFAULT 0,
  left_at bigint NOT NULL DEFAULT 0,
  state integer NOT NULL DEFAULT 0,
  date2 bigint NOT NULL DEFAULT 0,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS devices (
  id bigint NOT NULL,
  auth_key_id bigint NOT NULL,
  user_id bigint NOT NULL,
  token_type integer NOT NULL,
  token varchar(512) NOT NULL,
  no_muted boolean NOT NULL DEFAULT false,
  locked_period integer NOT NULL DEFAULT 0,
  app_sandbox boolean NOT NULL DEFAULT false,
  secret varchar(1024) NOT NULL DEFAULT '',
  other_uids varchar(1024) NOT NULL DEFAULT '',
  state boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS dialogs (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  peer_type integer NOT NULL DEFAULT 2,
  peer_id bigint NOT NULL,
  peer_dialog_id bigint NOT NULL,
  pinned bigint NOT NULL DEFAULT 0,
  top_message integer NOT NULL DEFAULT 0,
  pinned_msg_id integer NOT NULL DEFAULT 0,
  read_inbox_max_id integer NOT NULL DEFAULT 0,
  read_outbox_max_id integer NOT NULL DEFAULT 0,
  unread_count integer NOT NULL DEFAULT 0,
  unread_mentions_count integer NOT NULL DEFAULT 0,
  unread_mark boolean NOT NULL DEFAULT false,
  draft_type integer NOT NULL DEFAULT 0,
  draft_message_data jsonb NOT NULL,
  folder_id integer NOT NULL DEFAULT 0,
  folder_pinned bigint NOT NULL DEFAULT 0,
  has_scheduled boolean NOT NULL DEFAULT false,
  date2 bigint NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS dialog_filters (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  dialog_filter_id integer NOT NULL,
  dialog_filter jsonb NOT NULL,
  order_value bigint NOT NULL DEFAULT 0,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS documents (
  id bigint NOT NULL,
  document_id bigint NOT NULL,
  access_hash bigint NOT NULL,
  dc_id integer NOT NULL,
  file_path varchar(255) NOT NULL,
  file_size integer NOT NULL,
  uploaded_file_name varchar(255) NOT NULL DEFAULT '',
  ext varchar(32) NOT NULL DEFAULT '',
  mime_type varchar(128) NOT NULL DEFAULT '',
  thumb_id bigint NOT NULL DEFAULT 0,
  video_thumb_id bigint NOT NULL DEFAULT 0,
  version integer NOT NULL DEFAULT 0,
  attributes jsonb NOT NULL,
  date2 bigint NOT NULL DEFAULT 0,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS encrypted_files (
  id bigint NOT NULL,
  encrypted_file_id bigint NOT NULL,
  access_hash bigint NOT NULL,
  dc_id integer NOT NULL,
  file_size integer NOT NULL,
  key_fingerprint integer NOT NULL,
  md5_checksum varchar(32) NOT NULL DEFAULT '',
  file_path varchar(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS hash_tags (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  peer_type integer NOT NULL,
  peer_id bigint NOT NULL,
  hash_tag varchar(128) NOT NULL,
  hash_tag_message_id integer NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS imported_contacts (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  imported_user_id bigint NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS messages (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  user_message_box_id integer NOT NULL,
  dialog_id1 bigint NOT NULL,
  dialog_id2 bigint NOT NULL,
  dialog_message_id bigint NOT NULL,
  sender_user_id bigint NOT NULL,
  peer_type integer NOT NULL,
  peer_id bigint NOT NULL,
  random_id bigint NOT NULL DEFAULT 0,
  message_filter_type integer NOT NULL DEFAULT 0,
  message_data jsonb DEFAULT NULL,
  message varchar(6000) NOT NULL DEFAULT '',
  mentioned boolean NOT NULL DEFAULT false,
  media_unread boolean NOT NULL DEFAULT false,
  pinned boolean NOT NULL DEFAULT false,
  date2 bigint NOT NULL DEFAULT 0,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS phone_books (
  id bigint NOT NULL,
  user_id bigint NOT NULL DEFAULT 0,
  auth_key_id bigint NOT NULL,
  client_id bigint NOT NULL,
  phone varchar(32) NOT NULL,
  first_name varchar(64) NOT NULL,
  last_name varchar(64) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS photos (
  id bigint NOT NULL,
  photo_id bigint NOT NULL,
  access_hash bigint NOT NULL,
  has_stickers boolean NOT NULL DEFAULT false,
  dc_id integer NOT NULL DEFAULT 2,
  date2 bigint NOT NULL DEFAULT 0,
  has_video boolean NOT NULL DEFAULT false,
  size_id bigint NOT NULL DEFAULT 0,
  video_size_id bigint NOT NULL DEFAULT 0,
  input_file_name varchar(128) NOT NULL DEFAULT '',
  ext varchar(32) NOT NULL DEFAULT '',
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS photo_sizes (
  id bigint NOT NULL,
  photo_size_id bigint NOT NULL,
  size_type char(1) NOT NULL DEFAULT '',
  volume_id bigint NOT NULL,
  local_id integer NOT NULL,
  secret bigint NOT NULL,
  width integer NOT NULL,
  height integer NOT NULL,
  file_size integer NOT NULL,
  file_path varchar(255) NOT NULL,
  has_stripped boolean NOT NULL DEFAULT false,
  stripped_bytes varchar(4096) NOT NULL DEFAULT '',
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS popular_contacts (
  id bigint NOT NULL,
  phone varchar(32) NOT NULL,
  importers integer NOT NULL DEFAULT 1,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS predefined_users (
  id bigint NOT NULL,
  phone varchar(32) NOT NULL,
  first_name varchar(255) NOT NULL DEFAULT '',
  last_name varchar(255) NOT NULL DEFAULT '',
  username citext NOT NULL,
  code varchar(32) NOT NULL,
  verified boolean NOT NULL DEFAULT false,
  registered_user_id bigint NOT NULL DEFAULT 0,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS unregistered_contacts (
  id bigint NOT NULL,
  phone varchar(32) NOT NULL,
  importer_user_id bigint NOT NULL,
  import_first_name varchar(64) NOT NULL,
  import_last_name varchar(64) NOT NULL,
  imported boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS username (
  id bigint NOT NULL,
  username citext NOT NULL,
  peer_type integer NOT NULL DEFAULT 0,
  peer_id bigint NOT NULL DEFAULT 0,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS users (
  id bigint NOT NULL,
  user_type integer NOT NULL DEFAULT 2,
  access_hash bigint NOT NULL,
  secret_key_id bigint NOT NULL DEFAULT 0,
  first_name varchar(64) NOT NULL DEFAULT '',
  last_name varchar(64) NOT NULL DEFAULT '',
  username citext NOT NULL DEFAULT '',
  phone varchar(32) NOT NULL,
  country_code varchar(3) NOT NULL,
  verified boolean NOT NULL DEFAULT false,
  support boolean NOT NULL DEFAULT false,
  scam boolean NOT NULL DEFAULT false,
  fake boolean NOT NULL DEFAULT false,
  about varchar(128) NOT NULL DEFAULT '',
  state integer NOT NULL DEFAULT 0,
  is_bot boolean NOT NULL DEFAULT false,
  account_days_ttl integer NOT NULL DEFAULT 180,
  photo_id bigint NOT NULL DEFAULT 0,
  restricted boolean NOT NULL DEFAULT false,
  restriction_reason varchar(128) NOT NULL DEFAULT '',
  archive_and_mute_new_noncontact_peers boolean NOT NULL DEFAULT false,
  deleted boolean NOT NULL DEFAULT false,
  delete_reason varchar(128) NOT NULL DEFAULT '',
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_contacts (
  id bigint NOT NULL,
  owner_user_id bigint NOT NULL,
  contact_user_id bigint NOT NULL,
  contact_phone varchar(32) NOT NULL DEFAULT '',
  contact_first_name varchar(255) NOT NULL DEFAULT '',
  contact_last_name varchar(255) NOT NULL DEFAULT '',
  mutual boolean NOT NULL DEFAULT false,
  is_deleted boolean NOT NULL DEFAULT false,
  date2 bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_global_privacy_settings (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  archive_and_mute_new_noncontact_peers boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_notify_settings (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  peer_type integer NOT NULL,
  peer_id bigint NOT NULL,
  show_previews integer NOT NULL DEFAULT -1,
  silent integer NOT NULL DEFAULT -1,
  mute_until integer NOT NULL DEFAULT -1,
  sound varchar(255) NOT NULL DEFAULT 'default',
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_peer_blocks (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  peer_type integer NOT NULL,
  peer_id bigint NOT NULL,
  "date" bigint NOT NULL DEFAULT 0,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_peer_settings (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  peer_type integer NOT NULL,
  peer_id bigint NOT NULL,
  hide boolean NOT NULL DEFAULT false,
  report_spam boolean NOT NULL DEFAULT false,
  add_contact boolean NOT NULL DEFAULT false,
  block_contact boolean NOT NULL DEFAULT false,
  share_contact boolean NOT NULL DEFAULT false,
  need_contacts_exception boolean NOT NULL DEFAULT false,
  report_geo boolean NOT NULL DEFAULT false,
  autoarchived boolean NOT NULL DEFAULT false,
  invite_members boolean NOT NULL DEFAULT false,
  geo_distance integer NOT NULL DEFAULT 0,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_presences (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  last_seen_at bigint NOT NULL,
  expires integer NOT NULL DEFAULT 0,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_privacies (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  key_type integer NOT NULL DEFAULT 0,
  rules jsonb NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_profile_photos (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  photo_id bigint NOT NULL,
  date2 bigint NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_pts_updates (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  pts integer NOT NULL,
  pts_count integer NOT NULL,
  update_type smallint NOT NULL DEFAULT 0,
  update_data jsonb NOT NULL,
  date2 bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS user_settings (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  key2 varchar(64) NOT NULL,
  value varchar(512) NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS video_sizes (
  id bigint NOT NULL,
  video_size_id bigint NOT NULL,
  size_type char(1) NOT NULL DEFAULT '',
  volume_id bigint NOT NULL,
  local_id integer NOT NULL,
  secret bigint NOT NULL,
  width integer NOT NULL,
  height integer NOT NULL,
  file_size integer NOT NULL DEFAULT 0,
  video_start_ts double precision NOT NULL DEFAULT 0,
  file_path varchar(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.auths'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE auths ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS auths_auth_key_id_key ON auths (auth_key_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.auth_keys'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE auth_keys ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS auth_keys_auth_key_id_key ON auth_keys (auth_key_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.auth_key_infos'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE auth_key_infos ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS auth_key_infos_auth_key_id_key ON auth_key_infos (auth_key_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.auth_seq_updates'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE auth_seq_updates ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS auth_seq_updates_auth_id_key ON auth_seq_updates (auth_id, user_id, seq);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.auth_users'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE auth_users ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS auth_users_auth_key_id_key ON auth_users (auth_key_id, user_id);
CREATE INDEX IF NOT EXISTS auth_users_auth_key_id_2_idx ON auth_users (auth_key_id, user_id, deleted);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.bots'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE bots ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS bots_bot_id_key ON bots (bot_id);
CREATE UNIQUE INDEX IF NOT EXISTS bots_token_key ON bots (token);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.bot_commands'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE bot_commands ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE INDEX IF NOT EXISTS bot_commands_bot_id_idx ON bot_commands (bot_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.chats'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE chats ADD PRIMARY KEY (id);
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.chat_invites'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE chat_invites ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS chat_invites_link_key ON chat_invites (link);
CREATE INDEX IF NOT EXISTS chat_invites_chat_id_idx ON chat_invites (chat_id, permanent, admin_id);
CREATE INDEX IF NOT EXISTS chat_invites_chat_id_2_idx ON chat_invites (chat_id, admin_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.chat_invite_participants'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE chat_invite_participants ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS chat_invite_participants_link_key ON chat_invite_participants (link, user_id);
CREATE INDEX IF NOT EXISTS chat_invite_participants_link_2_idx ON chat_invite_participants (link);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.chat_participants'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE chat_participants ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS chat_participants_chat_id_2_key ON chat_participants (chat_id, user_id);
CREATE INDEX IF NOT EXISTS chat_participants_chat_id_idx ON chat_participants (chat_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.devices'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE devices ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS devices_auth_key_id_key ON devices (auth_key_id, user_id, token_type);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.dialogs'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE dialogs ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS dialogs_user_id_key ON dialogs (user_id, peer_type, peer_id);
CREATE UNIQUE INDEX IF NOT EXISTS dialogs_user_id_2_key ON dialogs (user_id, peer_dialog_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.dialog_filters'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE dialog_filters ADD PRIMARY KEY (id);
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.documents'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE documents ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE INDEX IF NOT EXISTS documents_document_id_idx ON documents (document_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.encrypted_files'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE encrypted_files ADD PRIMARY KEY (id);
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.hash_tags'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE hash_tags ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS hash_tags_user_id_4_key ON hash_tags (user_id, hash_tag, hash_tag_message_id);
CREATE INDEX IF NOT EXISTS hash_tags_user_id_idx ON hash_tags (user_id, hash_tag);
CREATE INDEX IF NOT EXISTS hash_tags_user_id_2_idx ON hash_tags (user_id, peer_type, peer_id, hash_tag);
CREATE INDEX IF NOT EXISTS hash_tags_user_id_3_idx ON hash_tags (user_id, hash_tag_message_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.imported_contacts'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE imported_contacts ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS imported_contacts_user_id_key ON imported_contacts (user_id);
CREATE UNIQUE INDEX IF NOT EXISTS imported_contacts_user_id_2_key ON imported_contacts (user_id, imported_user_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.messages'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE messages ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS messages_user_id_2_key ON messages (user_id, user_message_box_id);
CREATE INDEX IF NOT EXISTS messages_user_id_idx ON messages (user_id, dialog_id1, dialog_id2);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.phone_books'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE phone_books ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS phone_books_auth_key_id_key ON phone_books (auth_key_id, client_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.photos'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE photos ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS photos_photo_id_key ON photos (photo_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.photo_sizes'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE photo_sizes ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS photo_sizes_photo_size_id_key ON photo_sizes (photo_size_id, size_type);
CREATE INDEX IF NOT EXISTS photo_sizes_photo_id_idx ON photo_sizes (photo_size_id);
DO $$ BEGIN
  IF EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'photo_sizes' AND column_name = 'volume_id'
  ) THEN
    CREATE INDEX IF NOT EXISTS photo_sizes_volume_id_idx ON photo_sizes (volume_id, local_id);
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.popular_contacts'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE popular_contacts ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS popular_contacts_phone_key ON popular_contacts (phone);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.predefined_users'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE predefined_users ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS predefined_users_phone_key ON predefined_users (phone);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.unregistered_contacts'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE unregistered_contacts ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS unregistered_contacts_phone_key ON unregistered_contacts (phone, importer_user_id);
CREATE INDEX IF NOT EXISTS unregistered_contacts_phone_2_idx ON unregistered_contacts (phone, importer_user_id, imported);
CREATE INDEX IF NOT EXISTS unregistered_contacts_phone_3_idx ON unregistered_contacts (phone);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.username'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE username ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS username_username_key ON username (username);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.users'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE users ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS users_phone_key ON users (phone);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_contacts'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_contacts ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_contacts_owner_user_id_key ON user_contacts (owner_user_id, contact_user_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_global_privacy_settings'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_global_privacy_settings ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_global_privacy_settings_user_id_key ON user_global_privacy_settings (user_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_notify_settings'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_notify_settings ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_notify_settings_user_id_key ON user_notify_settings (user_id, peer_type, peer_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_peer_blocks'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_peer_blocks ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_peer_blocks_user_id_2_key ON user_peer_blocks (user_id, peer_type, peer_id);
CREATE INDEX IF NOT EXISTS user_peer_blocks_user_id_idx ON user_peer_blocks (user_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_peer_settings'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_peer_settings ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_peer_settings_user_id_key ON user_peer_settings (user_id, peer_type, peer_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_presences'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_presences ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_presences_user_id_key ON user_presences (user_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_privacies'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_privacies ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_privacies_user_id_key ON user_privacies (user_id, key_type);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_profile_photos'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_profile_photos ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_profile_photos_user_id_key ON user_profile_photos (user_id, photo_id);
CREATE INDEX IF NOT EXISTS user_profile_photos_user_id_2_idx ON user_profile_photos (user_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_pts_updates'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_pts_updates ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE INDEX IF NOT EXISTS user_pts_updates_user_id_idx ON user_pts_updates (user_id, pts);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.user_settings'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE user_settings ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS user_settings_user_id_key ON user_settings (user_id, key2);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint
    WHERE conrelid = 'public.video_sizes'::regclass AND contype = 'p'
  ) THEN
    ALTER TABLE video_sizes ADD PRIMARY KEY (id);
  END IF;
END $$;
CREATE UNIQUE INDEX IF NOT EXISTS video_sizes_video_size_id_key ON video_sizes (video_size_id, size_type);
DO $$ BEGIN
  IF EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'video_sizes' AND column_name = 'volume_id'
  ) THEN
    CREATE INDEX IF NOT EXISTS video_sizes_volume_id_idx ON video_sizes (volume_id, local_id);
  END IF;
END $$;
CREATE INDEX IF NOT EXISTS video_sizes_video_size_id_2_idx ON video_sizes (video_size_id);
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'auths' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE auths ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'auth_keys' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE auth_keys ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'auth_seq_updates' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE auth_seq_updates ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'auth_users' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE auth_users ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'bots' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE bots ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'bot_commands' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE bot_commands ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'chats' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE chats ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'chat_invites' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE chat_invites ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'chat_invite_participants' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE chat_invite_participants ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'chat_participants' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE chat_participants ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'devices' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE devices ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'dialogs' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE dialogs ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'dialog_filters' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE dialog_filters ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'documents' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE documents ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'encrypted_files' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE encrypted_files ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'hash_tags' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE hash_tags ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'imported_contacts' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE imported_contacts ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'messages' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE messages ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'phone_books' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE phone_books ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'photos' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE photos ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'photo_sizes' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE photo_sizes ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'popular_contacts' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE popular_contacts ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'predefined_users' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE predefined_users ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'unregistered_contacts' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE unregistered_contacts ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'username' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE username ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'users' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE users ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_contacts' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_contacts ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_global_privacy_settings' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_global_privacy_settings ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_notify_settings' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_notify_settings ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_peer_blocks' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_peer_blocks ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_peer_settings' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_peer_settings ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_presences' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_presences ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_privacies' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_privacies ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_profile_photos' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_profile_photos ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_pts_updates' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_pts_updates ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'user_settings' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE user_settings ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
DO $$ BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_schema = 'public' AND table_name = 'video_sizes' AND column_name = 'id' AND is_identity = 'YES'
  ) THEN
    ALTER TABLE video_sizes ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY;
  END IF;
END $$;
CREATE OR REPLACE FUNCTION kgramm_touch_update_at() RETURNS trigger
LANGUAGE plpgsql AS $$
BEGIN
  IF NEW.update_at IS NOT DISTINCT FROM OLD.update_at THEN
    NEW.update_at := now();
  END IF;
  RETURN NEW;
END;
$$;
CREATE OR REPLACE FUNCTION kgramm_touch_updated_at() RETURNS trigger
LANGUAGE plpgsql AS $$
BEGIN
  IF NEW.updated_at IS NOT DISTINCT FROM OLD.updated_at THEN
    NEW.updated_at := now();
  END IF;
  RETURN NEW;
END;
$$;
DROP TRIGGER IF EXISTS auth_key_infos_touch_updated_at ON auth_key_infos;
CREATE TRIGGER auth_key_infos_touch_updated_at BEFORE UPDATE ON auth_key_infos FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS auth_users_touch_updated_at ON auth_users;
CREATE TRIGGER auth_users_touch_updated_at BEFORE UPDATE ON auth_users FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS auths_touch_updated_at ON auths;
CREATE TRIGGER auths_touch_updated_at BEFORE UPDATE ON auths FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS bot_commands_touch_updated_at ON bot_commands;
CREATE TRIGGER bot_commands_touch_updated_at BEFORE UPDATE ON bot_commands FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS bots_touch_updated_at ON bots;
CREATE TRIGGER bots_touch_updated_at BEFORE UPDATE ON bots FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS chat_invite_participants_touch_updated_at ON chat_invite_participants;
CREATE TRIGGER chat_invite_participants_touch_updated_at BEFORE UPDATE ON chat_invite_participants FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS chat_invites_touch_updated_at ON chat_invites;
CREATE TRIGGER chat_invites_touch_updated_at BEFORE UPDATE ON chat_invites FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS chat_participants_touch_updated_at ON chat_participants;
CREATE TRIGGER chat_participants_touch_updated_at BEFORE UPDATE ON chat_participants FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS chats_touch_updated_at ON chats;
CREATE TRIGGER chats_touch_updated_at BEFORE UPDATE ON chats FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS devices_touch_updated_at ON devices;
CREATE TRIGGER devices_touch_updated_at BEFORE UPDATE ON devices FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS dialog_filters_touch_updated_at ON dialog_filters;
CREATE TRIGGER dialog_filters_touch_updated_at BEFORE UPDATE ON dialog_filters FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS dialogs_touch_updated_at ON dialogs;
CREATE TRIGGER dialogs_touch_updated_at BEFORE UPDATE ON dialogs FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS documents_touch_updated_at ON documents;
CREATE TRIGGER documents_touch_updated_at BEFORE UPDATE ON documents FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS hash_tags_touch_updated_at ON hash_tags;
CREATE TRIGGER hash_tags_touch_updated_at BEFORE UPDATE ON hash_tags FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS messages_touch_updated_at ON messages;
CREATE TRIGGER messages_touch_updated_at BEFORE UPDATE ON messages FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS phone_books_touch_updated_at ON phone_books;
CREATE TRIGGER phone_books_touch_updated_at BEFORE UPDATE ON phone_books FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS popular_contacts_touch_update_at ON popular_contacts;
CREATE TRIGGER popular_contacts_touch_update_at BEFORE UPDATE ON popular_contacts FOR EACH ROW EXECUTE FUNCTION kgramm_touch_update_at();
DROP TRIGGER IF EXISTS predefined_users_touch_updated_at ON predefined_users;
CREATE TRIGGER predefined_users_touch_updated_at BEFORE UPDATE ON predefined_users FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS unregistered_contacts_touch_updated_at ON unregistered_contacts;
CREATE TRIGGER unregistered_contacts_touch_updated_at BEFORE UPDATE ON unregistered_contacts FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_contacts_touch_updated_at ON user_contacts;
CREATE TRIGGER user_contacts_touch_updated_at BEFORE UPDATE ON user_contacts FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_global_privacy_settings_touch_updated_at ON user_global_privacy_settings;
CREATE TRIGGER user_global_privacy_settings_touch_updated_at BEFORE UPDATE ON user_global_privacy_settings FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_notify_settings_touch_updated_at ON user_notify_settings;
CREATE TRIGGER user_notify_settings_touch_updated_at BEFORE UPDATE ON user_notify_settings FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_peer_blocks_touch_updated_at ON user_peer_blocks;
CREATE TRIGGER user_peer_blocks_touch_updated_at BEFORE UPDATE ON user_peer_blocks FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_peer_settings_touch_updated_at ON user_peer_settings;
CREATE TRIGGER user_peer_settings_touch_updated_at BEFORE UPDATE ON user_peer_settings FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_presences_touch_updated_at ON user_presences;
CREATE TRIGGER user_presences_touch_updated_at BEFORE UPDATE ON user_presences FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_privacies_touch_updated_at ON user_privacies;
CREATE TRIGGER user_privacies_touch_updated_at BEFORE UPDATE ON user_privacies FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_profile_photos_touch_updated_at ON user_profile_photos;
CREATE TRIGGER user_profile_photos_touch_updated_at BEFORE UPDATE ON user_profile_photos FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS user_settings_touch_updated_at ON user_settings;
CREATE TRIGGER user_settings_touch_updated_at BEFORE UPDATE ON user_settings FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS username_touch_updated_at ON username;
CREATE TRIGGER username_touch_updated_at BEFORE UPDATE ON username FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
DROP TRIGGER IF EXISTS users_touch_updated_at ON users;
CREATE TRIGGER users_touch_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION kgramm_touch_updated_at();
