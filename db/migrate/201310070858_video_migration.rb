class VideoMigration < ActiveRecord::Migration
  def change
    create_table :videos do |t|
      t.integer :id
      t.integer :event_id
      t.integer :deputado_id
      t.string :video_url
    end
  end
end
