class DeputadosUriMigration < ActiveRecord::Migration
  def change
    add_column :deputados, :uri, :string
  end
end

