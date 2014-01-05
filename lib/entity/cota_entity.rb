class CotaEntity
  def self.represent(data, options)
    result = data.attributes
    result['deputado'] = Deputado.find(data.deputado_id)
    result
  end
end
