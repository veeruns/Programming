require 'spec_helper'

describe file('/etc/redhat-release') do
  it { should be_file }
end
