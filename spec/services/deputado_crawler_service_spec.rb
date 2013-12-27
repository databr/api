# encoding: utf-8
#
require 'spec_helper'

describe DeputadoCrawlerService do
  subject { DeputadoCrawlerService }
  let(:deputados_ids) { [141_463, 74_354, 73_933, 74_145, 160_625] }
  before do
    stub_request(:get, 'http://www2.camara.leg.br/deputados/pesquisa')
      .to_return(status: 200,
                 body: File.read('spec/fixtures/pesquisa.html'),
                 headers: { content_yype: 'text/html' })

    deputados_ids.each do |id|
      bio_url = 'http://www2.camara.leg.br/deputados/pesquisa/'
      bio_url << "layouts_deputados_biografia?pk=#{id}"
      stub_request(:get, bio_url)
        .to_return(status: 200,
                   body: File.read('spec/fixtures/biografia.html'))
    end

    deputados_url = 'http://www.camara.gov.br/SitCamaraWS/'
    deputados_url << 'Deputados.asmx/ObterDeputados'
    stub_request(:get, deputados_url)
      .to_return(status: 200, body: File.read('spec/fixtures/deputados.xml'))
  end

  describe '.save_from_pesquisa_parser' do
    it 'saves from pesquisa parser' do
      subject.save_from_pesquisa_parser
      expect(Deputado.all.count).to eq(5)
      expect(Deputado.all.map(&:cadastro_id)).to eq(deputados_ids)
    end
  end

  describe '.save_from_deputado_xml_parser' do
    it 'saves from deputado xml parser' do
      subject.save_from_deputado_xml_parser
      ids = [531_071, 520_939, 522_008, 522_840, 521_856]
      expect(Deputado.all.count).to eq(5)
      expect(Deputado.all.map(&:parlamentar_id)).to eq(ids)
    end
  end

  describe '.save_from_deputado_about_parser' do
    it 'saves from deputado about parser' do
      subject.save_from_deputado_about_parser
      expect(About.where(cadastro_id: 74_354).count).to eq(11)
    end
  end
end
