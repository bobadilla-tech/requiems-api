# frozen_string_literal: true

# New source of truth for plan limits that Go reads at request time (rate
# limiter + quota middleware, see docs/plans/2026-08-21-go-auth-foundation-phase-2.md).
# Deliberately NOT wired up to PlanConfig::PLANS or the Worker's config.ts in
# this migration/phase — those two copies keep their existing hardcoded
# values unchanged, so there are three manually-synced copies of plan limits
# for now (see the plan doc's Context section). Seeded here, in the
# migration itself rather than db/seeds.rb, because production's boot
# command runs `db:prepare` (which runs pending migrations) and never runs
# db:seed — this table has to exist and be populated in every environment,
# including production, the first time Go's rate limiter/quota middleware
# reads it.
class CreatePlans < ActiveRecord::Migration[8.1]
  def up
    create_table :plans, id: false do |t|
      t.string :id, null: false, primary_key: true
      t.integer :request_limit
      t.integer :rate_limit_per_minute
      t.timestamps
    end

    # Values mirrored from PlanConfig::PLANS (requests_per_month /
    # rate_limit_per_minute). enterprise has no PlanConfig entry today; both
    # limits are null here, meaning "unlimited" to the Go middleware.
    execute <<~SQL.squish
      INSERT INTO plans (id, request_limit, rate_limit_per_minute, created_at, updated_at) VALUES
        ('free', 500, 30, NOW(), NOW()),
        ('developer', 100000, 5000, NOW(), NOW()),
        ('business', 1000000, 10000, NOW(), NOW()),
        ('professional', 10000000, 50000, NOW(), NOW()),
        ('enterprise', NULL, NULL, NOW(), NOW())
      ON CONFLICT (id) DO NOTHING
    SQL

    execute <<~SQL.squish
      UPDATE subscriptions
      SET plan_name = 'free'
      WHERE plan_name IS NULL
         OR plan_name NOT IN ('free', 'developer', 'business', 'professional')
    SQL

    add_foreign_key :subscriptions, :plans, column: :plan_name, primary_key: :id
  end

  def down
    drop_table :plans
  end
end
