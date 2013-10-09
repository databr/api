class AddOtherNameToDeputados < ActiveRecord::Migration
  def change
    add_column :deputados, :other_name, :string
  end
end
