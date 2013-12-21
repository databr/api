class DeputadoAboutParser < CamaraParser
  def initialize(cadastro_id)
    @url = "http://www2.camara.leg.br/deputados/pesquisa/layouts_deputados_biografia?pk=#{cadastro_id}"
    super()
  end

  def sections
    @parser.search("#bioDeputado .bioOutros").map do |section|
      OpenStruct.new title: (section/".bioOutrosTitulo").text().gsub(':', ''),
                     text: (section/".bioOutrosTexto").text()
    end
  end
end
