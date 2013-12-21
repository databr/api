# encoding: utf-8
require 'spec_helper'

describe DeputadoProposicaoParser do
  def proposicao_url(deputado, year)
    url = "http://www.camara.gov.br/SitCamaraWS/Proposicoes.asmx/ListarProposicoes?sigla=PL&numero=&ano=#{Time.now.year}&datApresentacaoIni=&datApresentacaoFim=&autor=&parteNomeAutor=ABELARDO&siglaPartidoAutor=&siglaUFAutor=SP&generoAutor=&codEstado=&codOrgaoEstado=&emTramitacao=&idTipoAutor="
    stub_request(:get, url).
      to_return(status: 200, body: File.read('spec/fixtures/proposicoes.xml'))
    url
  end

  let(:deputado) { OpenStruct.new id: 74016, nome_parlamentar: "ABELARDO CAMARINHA", uf: "SP" }

  before do
    stub_request(:get, 'http://www.camara.gov.br/SitCamaraWS/Deputados.asmx/ObterDeputados').
      to_return(status: 200, body: File.read('spec/fixtures/deputados.xml'))

    stub_request(:get, "http://www2.camara.leg.br/deputados/pesquisa/layouts_deputados_biografia?pk=#{deputado.id}").
      to_return(status: 200, body: File.read('spec/fixtures/biografia.html'))
  end

  subject { DeputadoProposicaoParser.new(deputado) }
  describe "initialize" do
    it 'accept year' do
      url = proposicao_url(deputado, 2001)
      DeputadoProposicaoParser.new(deputado, 2001)
      expect(a_request(:get, url)).to have_been_made
    end

    it 'default year is the current year' do
      url = proposicao_url(deputado, Time.now.year)
      DeputadoProposicaoParser.new(deputado)
      expect(a_request(:get, url)).to have_been_made
    end
  end

  describe "#propositions" do
    it "get all propositions" do
      proposicao_url(deputado, Time.now.year)
      expect(subject.propositions.count).to eq(2)
    end
  end
end


