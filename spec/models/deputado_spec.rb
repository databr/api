require 'spec_helper'

describe Deputado do
  it { should validate_presence_of(:name) }
  it { should validate_presence_of(:deputado_id) }

  describe '.save_from_pesquisa_par,,ser' do
    it 'saves from pesquisa parser' do
      Deputado.save_from_pesquisa_parser
      expect(Deputado.all.count).to eq(5)
      expect(Deputado.all.map(&:deputado_id)).to eq(['141463', '74354', '73933', '74145', '160625'])
    end
  end
end
