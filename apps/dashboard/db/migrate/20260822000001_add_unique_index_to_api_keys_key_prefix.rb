# frozen_string_literal: true

# index_api_keys_on_key_prefix_btree (added in
# 20260821000000_add_btree_index_to_api_keys_key_prefix.rb) is a plain,
# non-unique btree — it speeds up the Go auth path's exact-match lookup but
# enforces nothing. key_prefix uniqueness has only ever been an app-level
# ActiveRecord validation, which cannot close the race two concurrent
# ApiKey#save! calls can hit. ApiKey#save!'s rescue ActiveRecord::RecordNotUnique
# retry (added in Phase 3 specifically for this) has therefore never been able
# to fire — there's no unique DB constraint to violate. Zero existing
# duplicate key_prefix values as of this writing (verified before adding this
# constraint), so this is safe to apply directly.
class AddUniqueIndexToApiKeysKeyPrefix < ActiveRecord::Migration[8.1]
  disable_ddl_transaction!

  def change
    add_index :api_keys, :key_prefix, name: "index_api_keys_on_key_prefix_unique", unique: true, algorithm: :concurrently
    remove_index :api_keys, name: "index_api_keys_on_key_prefix_btree", algorithm: :concurrently
  end
end
