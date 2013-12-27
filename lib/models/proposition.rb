# encoding: utf-8
#
class Proposition < ActiveRecord::Base
  def deputado
    @deputado ||= Deputado.where(cadastro_id: cadastro_id).first
  end

  def self.to_feed(deputado)
    where(cadastro_id: deputado.cadastro_id).order('presentations_at DESC')
  end
end
