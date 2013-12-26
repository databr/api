class Cota < ActiveRecord::Base
  self.table_name = "cotas"
  validates :numero, uniqueness: true

  def deputado
    @deputado ||= Deputado.find(self.deputado_id)
  end

  def self.to_feed(deputado)
    self.where(deputado_id: deputado.id).order("data_emissao DESC")
  end
end
