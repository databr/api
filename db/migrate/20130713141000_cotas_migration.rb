# encoding: utf-8
class CotasMigration < ActiveRecord::Migration
  def change
    create_table :cotas do |t|
      t.string :id # TODO: check it
      t.integer :deputado_id
      t.string :carteira_parlamentar
      t.string :legislatura
      t.string :uf
      t.string :partido
      t.integer :partido_id
      t.string :codigo_legislatura
      t.string :sub_cota
      t.string :descricao
      t.string :beneficiario
      t.string :cnpjcpf
      t.string :numero
      t.string :tipo_documento
      t.datetime :data_emissao
      t.decimal :valor_documento, precision: 10, scale: 2
      t.decimal :valor_glossa, precision: 10, scale: 2
      t.decimal :valor_liquido, precision: 10, scale: 2
      t.integer :mes
      t.integer :ano
      t.string :parcela
      t.string :lote
      t.string :ressarcimento
    end
  end
end
