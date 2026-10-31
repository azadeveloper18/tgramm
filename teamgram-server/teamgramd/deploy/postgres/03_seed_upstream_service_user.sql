-- Modified by Kgramm, 2026-10-31.
-- Derived from teamgram-server SQL, which remains under the Apache License 2.0.
-- MySQL dialect (ENGINE, AUTO_INCREMENT counters, charset, ON DUPLICATE KEY) is not preserved.

-- Upstream service user (name Teamgram, phone 42777). Phase 10 rewrites the name and phone.
-- Counters from the MySQL dump are not in this file. id 777000 is kept.
INSERT INTO users (id, user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, support, scam, fake, premium, about, state, is_bot, account_days_ttl, photo_id, restricted, restriction_reason, archive_and_mute_new_noncontact_peers, emoji_status_document_id, emoji_status_until, deleted, delete_reason, created_at, updated_at) VALUES(777000, 4, 6599886787491911851, 6895602324158323006, 'Teamgram', '', 'teamgram', '42777', '', true, false, false, false, false, '', 0, false, 180, 0, false, '', false, 0, 0, false, '', '2018-09-25 13:43:11', '2021-12-17 12:40:51')
ON CONFLICT (id) DO NOTHING;
