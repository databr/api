# encoding: utf-8
class DeputadosMigration < ActiveRecord::Migration
  def change
    create_table :deputados do |t|
      t.integer :id
      t.integer :partido_id
      t.integer :cadastro_id
      t.integer :parlamentar_id
      t.string :nome
      t.string :nome_parlamentar
      t.string :sexo
      t.string :uf
      t.integer :gabinete
      t.integer :anexo
      t.string :fone
      t.string :email
    end
  end
end
