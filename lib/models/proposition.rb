class Proposition < ActiveRecord::Base
  def deputado
    Deputado.where(cadastro_id: self.cadastro_id).first
  end
end
