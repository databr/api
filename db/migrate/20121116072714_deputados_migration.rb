class DeputadosMigration < ActiveRecord::Migration
  def change
    create_table :deputados do |t|
      t.string :name
      t.string :id
    end
  end
end
