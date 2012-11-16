class FixDeputadoIdToDeputados < ActiveRecord::Migration
  def change
    rename_column :deputados, :deputado_id, :cadastro_id
  end
end
