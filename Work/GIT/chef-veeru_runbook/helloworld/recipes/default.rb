#
# Cookbook Name:: helloworld
# Recipe:: default
#
# Copyright 2014 Yahoo Inc. All rights reserved.
#
# All rights reserved - Do Not Redistribute
#


template "/etc/config" 	do
	source "config.erb"
	owner "root"
	group "root"
	mode 00644
end
