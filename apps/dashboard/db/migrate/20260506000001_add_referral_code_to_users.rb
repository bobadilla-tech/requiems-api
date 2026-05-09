# frozen_string_literal: true

class AddReferralCodeToUsers < ActiveRecord::Migration[8.1]
  def up
    add_column :users, :referral_code, :string
    add_index :users, :referral_code, unique: true

    User.find_each do |user|
      begin
        user.update_columns(referral_code: SecureRandom.alphanumeric(8).upcase)
      rescue ActiveRecord::RecordNotUnique
        retry
      end
    end

    change_column_null :users, :referral_code, false
  end

  def down
    remove_column :users, :referral_code
  end
end
