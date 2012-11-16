class AddDeputadoIdToDeputados < ActiveRecord::Migration
  def change
    add_column :deputados, :deputado_id, :string
  end
end
