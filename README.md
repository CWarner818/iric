iric
====
iric (IRI Client) is a command line client for interfacing with
the IOTA Reference Implementation API. It is useful for debugging other programs
by seeing exactly what the API is returning. 

It is also a great tool for exploring the IOTA tangle and seeing how transactions
are related.`


Configuration
-------------

By default it will look in $HOME/.iric.<extension> for a config file. iric
supports many different formats for the config file, such as JSON, TOML, YAML,
and HCL. An example `.iric.yaml` file is shown below. 

```yaml
node: http://nodes.iota.fm:80
timeout: 3s```
