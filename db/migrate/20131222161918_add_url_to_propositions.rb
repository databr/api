class AddUrlToPropositions < ActiveRecord::Migration
  def change
    add_column :propositions, :url, :string
  end
end
