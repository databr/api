class DeputadoXMLParser < CamaraParser
  def initialize
    @url = "http://www.camara.gov.br/SitCamaraWS/Deputados.asmx/ObterDeputados"
    super
  end

  def deputados
    @deputados ||= (@parser/"deputado").map do |deputadoxml|
      map_attributes = {
        "nome" => :nome,
        "fone" => :fone,
        "email" => :email,
        "sexo" => :sexo,
        "gabinete" => :gabinete,
        "idecadastro" => :cadastro_id,
        "anexo" => :anexo,
        "uf" => :uf,
        "partido" => :partido,
        "comissoes" => :comissoes,
        "idparlamentar" => :parlamentar_id,
        "nomeparlamentar" => :nome_parlamentar
      }

      deputado = {}
      deputadoxml.children.each do |c|
        deputado[map_attributes[c.name]] = c.text
      end
      deputado
    end
  end
end
