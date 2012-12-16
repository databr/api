class Deputado < ActiveRecord::Base
  validates_presence_of :nome_parlamentar, :cadastro_id


end
