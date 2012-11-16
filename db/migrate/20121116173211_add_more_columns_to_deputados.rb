class AddMoreColumnsToDeputados < ActiveRecord::Migration
  def change
    add_column :deputados, :nome_parlamentar, :string
    add_column :deputados, :sexo, :string
    add_column :deputados, :uf, :string
    add_column :deputados, :partido_id, :integer
    add_column :deputados, :gabinete, :string
    add_column :deputados, :anexo, :string
    add_column :deputados, :fone, :string
    add_column :deputados, :email, :string
  end
end
