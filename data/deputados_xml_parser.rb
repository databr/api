class DeputadoXMLParser  < CamaraParser
  def initialize
    @url = "http://www.camara.gov.br/SitCamaraWS/Deputados.asmx/ObterDeputados"
    super
  end

  def deputados
    @deputados ||= (@parser/"deputado").map do |deputadoxml|
      deputado = {}
      deputadoxml.children.each do |c|
        deputado[c.name.to_sym] = c.text
      end
      deputado
    end
  end
end
