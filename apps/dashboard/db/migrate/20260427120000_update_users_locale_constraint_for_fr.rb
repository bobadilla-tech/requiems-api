# frozen_string_literal: true

class UpdateUsersLocaleConstraintForFr < ActiveRecord::Migration[8.1]
  def up
    execute "ALTER TABLE users DROP CONSTRAINT IF EXISTS locale_valid_values;"
    execute <<~SQL.squish
      ALTER TABLE users
        ADD CONSTRAINT locale_valid_values
        CHECK (locale IS NULL OR locale IN ('en', 'es', 'fr'));
    SQL
  end

  def down
    execute "ALTER TABLE users DROP CONSTRAINT IF EXISTS locale_valid_values;"
    execute <<~SQL.squish
      ALTER TABLE users
        ADD CONSTRAINT locale_valid_values
        CHECK (locale IS NULL OR locale IN ('en', 'es'));
    SQL
  end
end
