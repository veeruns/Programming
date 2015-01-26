# helloworld Cookbook

COOKBOOK_DESCRIPTION

## Requirements

Describe requirements for platform, system and external dependencies

## Usage

Give an example on how to use this cookbook

    run_list(
      'recipe[helloworld]'
    )
    default_attributes(
      'key' => 'vaule'
    )

## Attributes

  * `node['key']` - Value Type - descrption of the attribute

## Recipes

  * `default.rb` - describe what does the default recipe do

## Environment Setup

To use this cookbook, we recommend ensuring you development environment is set up with at least the minimum set of tools provided by the [cookbook development environment installer](http://devel.corp.yahoo.com/chef/guide/setup-dev-env.html).

## Tests

#### Bootstrap
    bundle install
    berks install

#### Unit tests
    rake

#### Functional Tests
    kitchen converge
    kitchen verify
    kitchen list


## Build

  * Edit `screwdriver/config.json` to point to your organization on the [Chef server](https://chef.ops.yahoo.com/)
  * Make sure the `by-screwdriver` headless user has access to your organization
  * Create a screwdriver project to build your cookbook and upload to the Chef server

## Bug?

Did you find a bug? Please file an issue in [JIRA Chef](https://jira.corp.yahoo.com/servicedesk/customer/chef). You can login JIRA using your Backyard credential. In addition, you can send us an pull request. Please follow the CONTRIBUTING.md.

## Enjoy
Please reach out to <chef-community@yahoo-inc.com> if you have any questions.

