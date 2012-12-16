require 'spec_helper'

describe Deputado do
  #it { should validate_uniqueness_of(:cadastro_id) }
  #it { should validate_uniqueness_of(:nome_parlamentar) }
  it { should validate_presence_of(:nome_parlamentar) }
  it { should validate_presence_of(:cadastro_id) }
end
