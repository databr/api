# encoding: utf-8
require 'spec_helper'

describe DeputadoAboutParser do
  let(:deputado) { OpenStruct.new id: 74_016 }
  before do
    stub_request(:get, 'http://www.camara.gov.br/SitCamaraWS/Deputados.asmx/ObterDeputados')
      .to_return(status: 200, body: File.read('spec/fixtures/deputados.xml'))

    stub_request(:get, "http://www2.camara.leg.br/deputados/pesquisa/layouts_deputados_biografia?pk=#{deputado.id}")
      .to_return(status: 200, body: File.read('spec/fixtures/biografia.html'))
  end

  subject { DeputadoAboutParser.new(deputado.id) }
  describe '#sections' do
    it 'returns 5 register' do
      expect(subject.sections.count).to eq(12)
      expect(subject.sections[0].title).to eq('Mandatos (na Câmara dos Deputados)')
      expect(subject.sections[0].text).to eq('Deputado Federal, 2007-2011, SP, PSB. Dt. Posse: 01/02/2007; Deputado Federal, 2011-2015, SP, PSB. Dt. Posse: 01/02/2011.')
    end
  end
end
