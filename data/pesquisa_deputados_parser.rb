# encoding: utf-8
class PesquisaDeputadosParser < CamaraParser
  URLS = {
    complete_info: 'http://www.camara.gov.br/internet/Deputado/dep_Detalhe.asp',
    bio:           'http://www2.camara.leg.br/deputados/pesquisa/layouts_deputados_biografia'
  }

  def initialize
    @url = 'http://www2.camara.leg.br/deputados/pesquisa'
    super()
  end

  def deputados
    @deputados ||= search_deputados
  end

  def complete_info_url(deputado_id = nil)
    url('complete_info', 'id', deputado_id)
  end

  def bio_url(deputado_id = nil)
    url('bio', 'pk', deputado_id)
  end

  private

  def search_deputados
    deputados = @parser.search('#deputado option')
    deputados = deputados.map do |deputado|
      id =  deputado.attr('value').split('?').last
      { nome_parlamentar: deputado.text, id: id } unless id.nil?
    end
    deputados.delete_if { |i| i.nil? }
  end

  def url(name, param, id = nil)
    url = URLS[name.to_sym]
    url += "?#{param}=#{id}" unless id.nil?
    url
  end
end
