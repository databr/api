class Event < ActiveRecord::Base
  has_many :videos
  validate :session_id, presence: true, uniqueness: true
end
