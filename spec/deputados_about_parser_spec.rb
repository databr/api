# encoding: utf-8
require 'spec_helper'

describe DeputadoAboutParser do
  let(:deputado) { OpenStruct.new(id: 74_016) }

  before do
    deputados_url = 'http://www.camara.gov.br/SitCamaraWS'
    deputados_url << '/Deputados.asmx/ObterDeputados'
    stub_request(:get, deputados_url)
      .to_return(status: 200, body: File.read('spec/fixtures/deputados.xml'))

    layout_url = 'http://www2.camara.leg.br/deputados/pesquisa'
    layout_url << "/layouts_deputados_biografia?pk=#{deputado.id}"
    stub_request(:get, layout_url)
      .to_return(status: 200, body: File.read('spec/fixtures/biografia.html'))
  end

  subject { DeputadoAboutParser.new(deputado.id) }
  describe '#sections' do
    it 'returns 5 register' do
      title = 'Mandatos (na Câmara dos Deputados)'
      text = 'Deputado Federal, 2007-2011, SP, PSB. Dt. Posse: 01/02/2007; '
      text << 'Deputado Federal, 2011-2015, SP, PSB. Dt. Posse: 01/02/2011.'

      expect(subject.sections.count).to eq(12)
      expect(subject.sections[0].title).to eq(title)
      expect(subject.sections[0].text).to eq(text)
    end
  end
end
