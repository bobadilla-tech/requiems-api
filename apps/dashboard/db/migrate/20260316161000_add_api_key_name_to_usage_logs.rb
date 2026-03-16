# frozen_string_literal: true

class AddApiKeyNameToUsageLogs < ActiveRecord::Migration[8.1]
  def change
    add_column :usage_logs, :api_key_name, :string
  end
end
