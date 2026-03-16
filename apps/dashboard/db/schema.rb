# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[8.1].define(version: 2026_03_16_161000) do
  # These are extensions that must be enabled in order to support this database
  enable_extension "pg_catalog.plpgsql"
  enable_extension "pg_trgm"

  create_table "abuse_reports", force: :cascade do |t|
    t.bigint "api_key_id", null: false
    t.datetime "created_at", null: false
    t.text "description"
    t.string "report_type"
    t.datetime "resolved_at"
    t.integer "resolved_by_id"
    t.string "status"
    t.datetime "updated_at", null: false
    t.bigint "user_id", null: false
    t.index ["api_key_id"], name: "index_abuse_reports_on_api_key_id"
    t.index ["status", "created_at"], name: "index_abuse_reports_on_status_and_created_at"
    t.index ["user_id"], name: "index_abuse_reports_on_user_id"
  end

  create_table "advice", id: :serial, force: :cascade do |t|
    t.text "text", null: false
  end

  create_table "api_keys", force: :cascade do |t|
    t.boolean "active"
    t.datetime "created_at", null: false
    t.string "environment"
    t.string "key_hash"
    t.string "key_prefix"
    t.datetime "last_used_at"
    t.string "last_used_ip"
    t.string "name"
    t.datetime "revoked_at"
    t.string "revoked_reason"
    t.datetime "updated_at", null: false
    t.bigint "user_id", null: false
    t.index ["key_prefix"], name: "index_api_keys_on_key_prefix_trgm", opclass: :gin_trgm_ops, using: :gin
    t.index ["user_id"], name: "index_api_keys_on_user_id"
  end

  create_table "audit_logs", force: :cascade do |t|
    t.string "action"
    t.integer "admin_user_id"
    t.datetime "created_at", null: false
    t.text "details"
    t.string "ip_address"
    t.datetime "updated_at", null: false
    t.bigint "user_id", null: false
    t.index ["admin_user_id"], name: "index_audit_logs_on_admin_user_id"
    t.index ["created_at"], name: "index_audit_logs_on_created_at"
    t.index ["user_id"], name: "index_audit_logs_on_user_id"
  end

  create_table "bin_data", primary_key: "bin_prefix", id: { type: :string, limit: 8 }, force: :cascade do |t|
    t.text "card_level", default: "", null: false
    t.text "card_type", default: "", null: false
    t.decimal "confidence", precision: 3, scale: 2, default: "0.5", null: false
    t.string "country_code", limit: 2, default: "", null: false
    t.text "country_name", default: "", null: false
    t.timestamptz "first_seen", default: -> { "now()" }, null: false
    t.text "issuer_name", default: "", null: false
    t.text "issuer_phone", default: "", null: false
    t.text "issuer_url", default: "", null: false
    t.timestamptz "last_updated", default: -> { "now()" }, null: false
    t.integer "prefix_length", limit: 2, null: false
    t.boolean "prepaid", default: false, null: false
    t.text "scheme", default: "", null: false
    t.text "source", default: "", null: false
    t.index "\"left\"((bin_prefix)::text, 6)", name: "idx_bin_data_prefix6"
    t.index ["country_code"], name: "idx_bin_data_country"
    t.index ["scheme"], name: "idx_bin_data_scheme"
  end

  create_table "counters", primary_key: "namespace", id: :text, force: :cascade do |t|
    t.bigint "total", default: 0, null: false
    t.datetime "updated_at", precision: nil, null: false
  end

  create_table "credit_adjustments", force: :cascade do |t|
    t.string "adjustment_type"
    t.integer "admin_user_id"
    t.integer "amount"
    t.datetime "created_at", null: false
    t.text "reason"
    t.datetime "updated_at", null: false
    t.bigint "user_id", null: false
    t.index ["user_id"], name: "index_credit_adjustments_on_user_id"
  end

  create_table "daily_usage_summaries", force: :cascade do |t|
    t.datetime "created_at", null: false
    t.date "date"
    t.integer "total_credits"
    t.integer "total_requests"
    t.datetime "updated_at", null: false
    t.bigint "user_id", null: false
    t.index ["user_id", "date"], name: "index_daily_usage_on_user_and_date"
    t.index ["user_id"], name: "index_daily_usage_summaries_on_user_id"
  end

  create_table "quotes", id: :serial, force: :cascade do |t|
    t.text "author"
    t.text "text", null: false
  end

  create_table "schema_migrations", primary_key: "version", id: :bigint, default: nil, force: :cascade do |t|
    t.boolean "dirty", null: false
  end

  create_table "solid_cache_entries", force: :cascade do |t|
    t.integer "byte_size", null: false
    t.datetime "created_at", null: false
    t.binary "key", null: false
    t.bigint "key_hash", null: false
    t.binary "value", null: false
    t.index ["byte_size"], name: "index_solid_cache_entries_on_byte_size"
    t.index ["key_hash", "byte_size"], name: "index_solid_cache_entries_on_key_hash_and_byte_size"
    t.index ["key_hash"], name: "index_solid_cache_entries_on_key_hash", unique: true
  end

  create_table "subscriptions", force: :cascade do |t|
    t.boolean "cancel_at_period_end", default: false
    t.datetime "canceled_at"
    t.datetime "created_at", null: false
    t.integer "credit_limit"
    t.datetime "current_period_end"
    t.datetime "current_period_start"
    t.string "lemonsqueezy_customer_id"
    t.string "lemonsqueezy_subscription_id"
    t.string "plan"
    t.string "plan_name"
    t.string "status"
    t.string "stripe_customer_id"
    t.string "stripe_subscription_id"
    t.datetime "trial_ends_at"
    t.datetime "updated_at", null: false
    t.bigint "user_id", null: false
    t.index ["lemonsqueezy_customer_id"], name: "index_subscriptions_on_lemonsqueezy_customer_id"
    t.index ["lemonsqueezy_subscription_id"], name: "index_subscriptions_on_lemonsqueezy_subscription_id", unique: true
    t.index ["plan_name"], name: "index_subscriptions_on_plan_name"
    t.index ["user_id"], name: "index_subscriptions_on_user_id"
  end

  create_table "usage_logs", force: :cascade do |t|
    t.bigint "api_key_id", null: false
    t.string "api_key_name"
    t.datetime "created_at", null: false
    t.integer "credits_used"
    t.string "endpoint"
    t.integer "response_time_ms"
    t.integer "status_code"
    t.datetime "updated_at", null: false
    t.date "usage_date"
    t.datetime "used_at"
    t.bigint "user_id", null: false
    t.index ["api_key_id", "used_at"], name: "index_usage_logs_on_api_key_and_time"
    t.index ["api_key_id"], name: "index_usage_logs_on_api_key_id"
    t.index ["endpoint", "used_at"], name: "index_usage_logs_on_endpoint_and_time"
    t.index ["status_code"], name: "index_usage_logs_on_status_code"
    t.index ["usage_date"], name: "index_usage_logs_on_usage_date"
    t.index ["user_id", "used_at"], name: "index_usage_logs_on_user_and_time"
    t.index ["user_id"], name: "index_usage_logs_on_user_id"
  end

  create_table "users", force: :cascade do |t|
    t.boolean "active", default: true, null: false
    t.boolean "admin", default: false, null: false
    t.datetime "banned_at"
    t.string "banned_reason"
    t.string "company"
    t.datetime "confirmation_sent_at"
    t.string "confirmation_token"
    t.datetime "confirmed_at"
    t.datetime "created_at", null: false
    t.datetime "current_sign_in_at"
    t.string "current_sign_in_ip"
    t.text "deletion_reason"
    t.string "deletion_token"
    t.datetime "deletion_token_sent_at"
    t.string "email", default: "", null: false
    t.boolean "email_notifications", default: true, null: false
    t.string "encrypted_password", default: "", null: false
    t.datetime "last_sign_in_at"
    t.string "last_sign_in_ip"
    t.string "locale"
    t.string "name"
    t.text "notes"
    t.datetime "remember_created_at"
    t.datetime "reset_password_sent_at"
    t.string "reset_password_token"
    t.integer "sign_in_count", default: 0, null: false
    t.string "status", default: "active", null: false
    t.string "unconfirmed_email"
    t.datetime "updated_at", null: false
    t.boolean "usage_alerts", default: true, null: false
    t.boolean "weekly_reports", default: false, null: false
    t.index ["admin"], name: "index_users_on_admin"
    t.index ["confirmation_token"], name: "index_users_on_confirmation_token", unique: true
    t.index ["deletion_token"], name: "index_users_on_deletion_token", unique: true
    t.index ["email"], name: "index_users_on_email", unique: true
    t.index ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true
    t.index ["status"], name: "index_users_on_status"
    t.check_constraint "locale IS NULL OR (locale::text = ANY (ARRAY['en'::character varying, 'es'::character varying]::text[]))", name: "locale_valid_values"
  end

  create_table "words", id: :serial, force: :cascade do |t|
    t.text "definition", null: false
    t.text "part_of_speech"
    t.text "word", null: false
  end

  add_foreign_key "abuse_reports", "api_keys"
  add_foreign_key "abuse_reports", "users"
  add_foreign_key "api_keys", "users"
  add_foreign_key "audit_logs", "users"
  add_foreign_key "credit_adjustments", "users"
  add_foreign_key "daily_usage_summaries", "users"
  add_foreign_key "subscriptions", "users"
  add_foreign_key "usage_logs", "api_keys"
  add_foreign_key "usage_logs", "users"
end
