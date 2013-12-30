# encoding: utf-8
#
class Proposition < ActiveRecord::Base
  def self.to_feed(deputado)
    where(cadastro_id: deputado['cadastro_id']).order('presentations_at DESC')
  end
end
