class Deputado < ActiveRecord::Base
  validates :nome_parlamentar, presence: true
  validates :cadastro_id, uniqueness: true, presence: true

  has_many :cotas
  has_many :videos

  before_save :set_uri

  private
    def set_uri
      self.uri = self.nome_parlamentar.parameterize
    end
end
