class Deputado < ActiveRecord::Base
  validates_presence_of :nome_parlamentar, :cadastro_id

  has_many :cotas

  before_save :set_uri

  private
    def set_uri
      self.uri = self.nome_parlamentar.parameterize
    end
end
