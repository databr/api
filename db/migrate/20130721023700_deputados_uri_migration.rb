class DeputadosUriMigration < ActiveRecord::Migration
  def change
    add_column :deputados, :uri, :string
    add_index :deputados, :uri
  end
end

