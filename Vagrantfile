# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.network :public_network,  ip: "192.168.33.22"

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "setup/playbook.yml"
    ansible.extra_vars = { ansible_ssh_user: 'vagrant' }
    ansible.limit = "all"
  end
end
