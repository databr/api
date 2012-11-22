require 'spec_helper'

describe Deputado do
  #it { should validate_uniqueness_of(:cadastro_id) }
  #it { should validate_uniqueness_of(:nome_parlamentar) }
  it { should validate_presence_of(:nome_parlamentar) }
  it { should validate_presence_of(:cadastro_id) }

  before do
    stub_request(:get, 'http://www2.camara.leg.br/deputados/pesquisa').
      to_return(:status => 200, :body => File.read('spec/fixtures/pesquisa.html'), :headers => {:'Content-Type' => 'text/html'})

    stub_request(:get, 'http://www.camara.gov.br/SitCamaraWS/Deputados.asmx/ObterDeputados').
      to_return(:status => 200, :body => File.read('spec/fixtures/deputados.xml'))
  end


  describe '.save_from_pesquisa_parser' do

    it 'saves from pesquisa parser' do
      Deputado.save_from_pesquisa_parser
      expect(Deputado.all.count).to eq(5)
      expect(Deputado.all.map(&:cadastro_id)).to eq([141463, 74354, 73933, 74145, 160625])
    end
  end

  describe ".save_from_deputado_xml_parser" do
    it 'saves from deputado xml parser' do
      Deputado.save_from_deputado_xml_parser
      expect(Deputado.all.count).to eq(5)
      expect(Deputado.all.map(&:parlamentar_id)).to eq([531071, 520939, 522008, 522840, 521856])
    end
  end
end
