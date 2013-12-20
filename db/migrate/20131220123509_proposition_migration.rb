class PropositionMigration < ActiveRecord::Migration
  def change
    create_table :propositions do |t|
      t.integer :id
      t.integer :cadastro_id
      t.integer :proposition_id
      t.string :name
      t.string :number
      t.string :year
      t.datetime :presentations_at
      t.text :body
    end
    add_index :propositions, :cadastro_id
    add_index :propositions, :proposition_id
  end
end
