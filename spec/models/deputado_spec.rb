require 'spec_helper'

describe Deputado do
  it { should validate_presence_of(:nome_parlamentar) }
  it { should validate_presence_of(:cadastro_id) }

  describe '.save_from_pesquisa_par,,ser' do
    before do
      stub_request(:get, 'http://www2.camara.leg.br/deputados/pesquisa').
        to_return(:status => 200, :body => File.read('spec/fixtures/pesquisa.html'), :headers => {:'Content-Type' => 'text/html'})
    end

    it 'saves from pesquisa parser' do
      Deputado.save_from_pesquisa_parser
      expect(Deputado.all.count).to eq(5)
      expect(Deputado.all.map(&:cadastro_id)).to eq(['141463', '74354', '73933', '74145', '160625'])
    end
  end
end
