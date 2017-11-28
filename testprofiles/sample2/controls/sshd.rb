title 'test control for testing test success'

control 'sshd-1.0' do
  impact 0.7
  title 'sshd config'
  desc 'sshd is expected to be running'
  describe service('sshd') do
    it { should be_installed }
    it { should be_enabled }
    it { should be_running }
  end
end
