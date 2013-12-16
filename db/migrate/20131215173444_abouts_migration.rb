class AboutsMigration < ActiveRecord::Migration
  def change
    create_table :abouts do |t|
      t.integer :id
      t.integer :cadastro_id
      t.string :title
      t.text :body
      t.string :section_key
    end
    add_index :abouts, :cadastro_id
  end
end
