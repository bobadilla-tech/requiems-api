# frozen_string_literal: true

class CreateReferrals < ActiveRecord::Migration[8.1]
  def change
    create_table :referrals do |t|
      t.bigint :referrer_id, null: false
      t.bigint :referred_user_id, null: false
      t.string :status, null: false, default: "pending"
      t.datetime :converted_at
      t.bigint :converting_subscription_id

      t.timestamps
    end

    add_index :referrals, :referred_user_id, unique: true
    add_index :referrals, :referrer_id

    add_foreign_key :referrals, :users, column: :referrer_id
    add_foreign_key :referrals, :users, column: :referred_user_id
    add_foreign_key :referrals, :subscriptions, column: :converting_subscription_id

    add_check_constraint :referrals,
      "referrer_id != referred_user_id",
      name: "chk_referrals_no_self_referral"
  end
end
