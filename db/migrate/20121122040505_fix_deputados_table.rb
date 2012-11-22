class FixDeputadosTable < ActiveRecord::Migration
  def change
    rename_column :deputados, :name, :nome
    add_column :deputados, :parlamentar_id, :integer
    change_column :deputados, :cadastro_id, :integer
    change_column :deputados, :gabinete, :integer
    change_column :deputados, :anexo, :integer
  end
end
