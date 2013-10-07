class EventsMigration < ActiveRecord::Migration
  def change
    create_table :events do |t|
      t.integer :id
      t.integer :session_id
      t.datetime :starts_at
      t.string :local
      t.text :title
      t.text :description
    end
  end
end

