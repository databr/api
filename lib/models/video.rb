# encoding: utf-8
class Video < ActiveRecord::Base
  belongs_to :event
  belongs_to :deputado
  default_scope includes(:event)
end
