require 'spec_helper'

describe DeputadoPropositionsService do
  subject { DeputadoPropositionsService }

  before do
    stub_request(:get, 'http://www.camara.gov.br/SitCamaraWS/Deputados.asmx/ObterDeputados').
      to_return(:status => 200, :body => File.read('spec/fixtures/deputado.xml'))

    stub_request(:get, "http://www.camara.gov.br/SitCamaraWS/Proposicoes.asmx/ListarProposicoes?sigla=PL&numero=&ano=#{Time.now.year}&datApresentacaoIni=&datApresentacaoFim=&autor=&parteNomeAutor=COSTA&siglaPartidoAutor=&siglaUFAutor=MA&generoAutor=&codEstado=&codOrgaoEstado=&emTramitacao=&idTipoAutor=").
      to_return(status: 200, body: File.read('spec/fixtures/proposicoes.xml'))
  end

  describe ".save_propositions" do
    it 'save proprositions on database' do
      subject.save_propositions
      expect(Proposition.all.count).to eq(2)
    end
  end
end

