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

name = golang-seed-change-me
namespace = kubertnetes-namespace-change-me

kafkatopic = USR.EVENT.blackbeard.article.view
kafkabrokers = worker.i.dev.tda.io:9092
kakfaconsumergroup = flint_test
prometheusendpoint = localhost:8081
redishost = localhost
redisport = 6379
redisdb = 0
rediskey= users
redisusergroups = user_groups

Note that files `tdaci.env` and `tdaci.yml` can be ignored.

Any improvements and suggestions can be sent via pull request or by [email](igor.miletic@tamedia.ch).
