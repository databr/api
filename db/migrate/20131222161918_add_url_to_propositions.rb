# encoding: utf-8
class AddUrlToPropositions < ActiveRecord::Migration
  def change
    add_column :propositions, :url, :string
  end
end
