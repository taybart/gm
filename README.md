## Dice

gm roll d20

## Create monster

gm monster -h 100 -name mon1

gm mon -f goblin.json

## Create conflict

gm scene init -f player1.json -f player2.json -monsters ./scene2

gm attack -to monster1 -from gandalf

## cli vs web interface

gm.handlers.roll(20) // cli

-- or --

gm.client.handlers.roll.call(20) // web interface
