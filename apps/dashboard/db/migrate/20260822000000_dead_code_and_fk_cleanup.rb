# frozen_string_literal: true

# Confirmed zero references (re-grepped immediately before writing this
# migration, per docs/plans/2026-08-22-go-auth-foundation-phase-5.md item 5):
# - subscriptions.stripe_customer_id / stripe_subscription_id / credit_limit —
#   dead Stripe-era columns; billing is LemonSqueezy-only.
# - solid_cache_entries — Rails' cache store is always overridden to
#   Redis/null (see config/environments), this table is never read or written.
#
# admin_user_id/resolved_by_id have zero orphaned values as of this writing
# (verified directly against production data before adding these FKs) — same
# pattern as the existing subscriptions.promoted_by_id FK.
class DeadCodeAndFkCleanup < ActiveRecord::Migration[8.1]
  def change
    remove_column :subscriptions, :stripe_customer_id, :string
    remove_column :subscriptions, :stripe_subscription_id, :string
    remove_column :subscriptions, :credit_limit, :integer

    drop_table :solid_cache_entries, force: :cascade do |t|
      t.binary "key", null: false
      t.binary "value", null: false
      t.bigint "key_hash", null: false
      t.integer "byte_size", null: false
      t.datetime "created_at", null: false
      t.index [ "byte_size" ], name: "index_solid_cache_entries_on_byte_size"
      t.index [ "key_hash", "byte_size" ], name: "index_solid_cache_entries_on_key_hash_and_byte_size"
      t.index [ "key_hash" ], name: "index_solid_cache_entries_on_key_hash", unique: true
    end

    add_foreign_key :credit_adjustments, :users, column: :admin_user_id
    add_foreign_key :audit_logs, :users, column: :admin_user_id
    add_foreign_key :abuse_reports, :users, column: :resolved_by_id
  end
end
