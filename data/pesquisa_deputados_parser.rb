class PesquisaDeputadosParser
  def initialize
    @agent = Mechanize.new
    @pesquisa = @agent.get "http://www2.camara.leg.br/deputados/pesquisa"
  end

  def deputados
    search_deputados
  end

  private

  def search_deputados
    deputados = @pesquisa.search("#deputado option")
    deputados.map do |deputado|
      id =  deputado.attr("value").split("?").last
      {name: deputado.text, id: id} unless id.nil?
    end.delete_if{|i| i.nil?}
  end
end
