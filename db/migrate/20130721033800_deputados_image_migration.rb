class DeputadosImageMigration < ActiveRecord::Migration
  def change
    add_column :deputados, :image_url, :string
  end
end
