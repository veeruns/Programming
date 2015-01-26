# Contributing to Cookbook

We are glad you want to contribute! The first step is the desire to improve the
project. Please visit [yo/code4swag](http://yo/code4swag) for detail.

## Installation and Setup

### Fork and Clone the Cookbook

    # Fork the cookbook to your own organization
    git clone git@git.corp.yahoo.com:<YOUR ORGANIZATION>/helloworld.git
    cd helloworld
    bundle install
    berks install
    rake # Make sure all tests are passing, or contact the maintainer

### Pull Requests

In your forked cookbook, please open a new branch. Make modification follow the [cookbook standard](http://devel.corp.yahoo.com/chef/guide/platform-cookbook-standard.html), test your changes, and make sure it converges on vagrant. Then, send a pull request to the upstream cookbook repo.

### Testing

    cd helloworld
    bundle install
    berks install
    
    # Run linting
    foodcritic .
    
    # Run unit test
    rake

    # Run vagrant convergence
    bundle exec kitchen converge

