# Go Seed project

This is seed project for Golang stack. It provides minimal
but fully functional project structure for project

## How to use it

Project is based on [Gitter8](http://www.foundweekends.org/giter8/) templating engine.

- With [SBT new](https://www.scala-sbt.org/1.0/docs/sbt-new-and-Templates.html)

 `sbt new https://github.com/tamediadigital/go-seed.g8.git`
 
 or with
 
- With [Gitter8](http://www.foundweekends.org/giter8/setup.html)

 `g8 https://github.com/tamediadigital/go-seed.g8.git`


Project contains couple of variables that you have to accept or change during (you will be prompted to change them):

- name = scala-akka-seed-change-me
- namespace = kubertnetes-namespace-change-me
- akka_version=2.5.17
- akka_http_version=10.1.5
- scala_test_version=3.0.5
- jackson_version=2.8.6
- log4j_version=1.2.17
- package=ch.tamedia
- nexus=https://your.nexus.url

Note that files `tdaci.env` and `tdaci.yml` can be ignored.

Any improvements and suggestions can be sent via pull request or by [email](igor.miletic@tamedia.ch).
