require 'spec_helper'

describe DeputadoXMLParser do
  let(:parser) { DeputadoXMLParser.new }

  before do
    stub_request(:get, 'http://www.camara.gov.br/SitCamaraWS/Deputados.asmx/ObterDeputados').
      to_return(:status => 200, :body => File.read('spec/fixtures/deputados.xml'), :headers => {:'Content-Type' => 'text/html'})
  end

  describe '#deputados' do
    it 'returns 5 register' do
      expect(parser.deputados.count).to eq(5)
    end

    it 'returns deputados info' do
      expect(parser.deputados.map{|d| d[:idecadastro] }).to eq(["74016", "74210", "74319", "74324", "74421"])
      expect(parser.deputados.first.keys).to eq([:idecadastro, :idparlamentar, :nome, :nomeparlamentar, :sexo, :uf, :partido, :gabinete, :anexo, :fone, :email, :comissoes])
    end
  end
end
