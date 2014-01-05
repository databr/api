# encoding: utf-8
class PropositionEntity
  def self.represent(data, options = {})
    result = data.attributes
    result["deputado"] = Deputado.cached_by_cadastro_id(data.cadastro_id)
    result
  end
end
