# frozen_string_literal: true

class AddBtreeIndexToApiKeysKeyPrefix < ActiveRecord::Migration[8.1]
  def change
    # The existing index_api_keys_on_key_prefix_trgm (GIN/trigram) is for
    # fuzzy admin search. The Go auth path's candidate-then-verify lookup
    # (WHERE key_prefix = ?) needs a plain btree for efficient exact-match.
    add_index :api_keys, :key_prefix, name: "index_api_keys_on_key_prefix_btree"
  end
end
