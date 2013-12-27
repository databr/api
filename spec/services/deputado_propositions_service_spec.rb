# encoding: utf-8
#
require 'spec_helper'

describe DeputadoPropositionsService do
  subject { DeputadoPropositionsService }

  before do
    deputados_url = 'http://www.camara.gov.br/SitCamaraWS/'
    deputados_url << 'Deputados.asmx/ObterDeputados'
    stub_request(:get, deputados_url)
      .to_return(status: 200, body: File.read('spec/fixtures/deputado.xml'))

    propositions_url = 'http://www.camara.gov.br/SitCamaraWS/Proposicoes.asmx/'
    propositions_url << 'ListarProposicoes?sigla=PL&numero=&'
    propositions_url << "ano=#{Time.now.year}&datApresentacaoIni=&"
    propositions_url << 'datApresentacaoFim=&autor=&parteNomeAutor=COSTA&'
    propositions_url << 'siglaPartidoAutor=&siglaUFAutor=&generoAutor=&'
    propositions_url << 'codEstado=&codOrgaoEstado=&emTramitacao=&idTipoAutor='

    stub_request(:get, propositions_url)
      .to_return(status: 200, body: File.read('spec/fixtures/proposicoes.xml'))
  end

  describe '.save_propositions' do
    it 'save proprositions on database' do
      subject.save_propositions
      expect(Proposition.all.count).to eq(2)
    end
  end
end
