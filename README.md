# AppEnv

appenv lets you discover what versions of the various systems your application
needs.

Commit an appenv.conf file to your project and then run appenv to see all the
versions of things you're running

## Supported kinds

This list should be ever growing! If your binary is not supported, please file
an issue.

### Kubernetes

Find container image versions from a Deployment, StatefulSet or Pod

### General

Find various kinds

- name: python
  command: python --version
- name: ruby
  command: ruby

