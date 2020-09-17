# GM (Game Master) a tool for crawling dungeons

Like a text adventure but with other people (i know MUD's exist, this is IRL people)

## Roll

```sh
$ gm -r d20
1492-04-01 00:00:00 [INFO] You rolled 20
```

## REPL

```sh
$ gm
-> r d20
1492-04-01 00:00:00 [INFO] You rolled 1
-> r d6
1492-04-01 00:00:00 [INFO] You rolled 2
-> q
```





# v2

## Goals

- [ ] Create monster template

- [ ] Create player templates

- [ ] Create persitent game state

- [ ] Create combat rolls/player state updates


## Create monster

`$ gm monster -h 100 -name mon1 -t ./mon1.template`

## Create conflict

`$ gm scene init -f player1.json -f player2.json -monsters ./scene2`

`$ gm attack -from coat -to wind -attack ./attacks/notw.attack`
