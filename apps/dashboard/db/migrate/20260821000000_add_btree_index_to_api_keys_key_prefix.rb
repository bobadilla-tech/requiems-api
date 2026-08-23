# frozen_string_literal: true

class AddBtreeIndexToApiKeysKeyPrefix < ActiveRecord::Migration[8.1]
  disable_ddl_transaction!

  def change
    # The existing index_api_keys_on_key_prefix_trgm (GIN/trigram) is for
    # fuzzy admin search. The Go auth path's candidate-then-verify lookup
    # (WHERE key_prefix = ?) needs a plain btree for efficient exact-match.
    # CONCURRENTLY avoids locking api_keys against writes for the build's
    # duration; it can't run inside a transaction, hence disable_ddl_transaction!.
    add_index :api_keys, :key_prefix, name: "index_api_keys_on_key_prefix_btree", algorithm: :concurrently
  end
end
