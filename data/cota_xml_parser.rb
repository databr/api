# encoding: utf-8
class CotaXMLParser < CamaraParser
  XML2DB = {
    'txNomeParlamentar' => :nome_parlamentar,
    'nuCarteiraParlamentar' => :carteira_parlamentar,
    'nuLegislatura' => :legislatura,
    'sgUF' => :uf,
    'sgPartido' => :partido,
    'codLegislatura' => :codigo_legislatura,
    'numSubCota' => :sub_cota,
    'txtDescricao' => :descricao ,
    'txtBeneficiario' => :beneficiario ,
    'txtCNPJCPF' => :cnpjcpf,
    'txtNumero' => :numero,
    'indTipoDocumento' => :tipo_documento,
    'datEmissao' => :data_emissao,
    'vlrDocumento' => :valor_documento,
    'vlrGlosa' => :valor_glossa,
    'vlrLiquido' => :valor_liquido,
    'numMes' => :mes,
    'numAno' => :ano,
    'numParcela' => :parcela,
    'numLote' => :lote,
    'numRessarcimento' => :ressarcimento
  }

  def initialize(file)
    @file = File.expand_path(file).to_s
    @parser = Nokogiri::XML(File.read(@file))
  end

  def cotas
    @parser.search('DESPESA').map do |despesaxml|
      despesa = {}
      despesaxml.children.each do |c|
        next unless XML2DB[c.name]
        despesa[XML2DB[c.name]] = c.text
      end
      despesa
    end
  end
end
