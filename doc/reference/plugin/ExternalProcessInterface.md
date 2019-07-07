# ExternalProcessInterface

Run an external process, using its stdin and stdous as input and output

## Args

| name    | type     | description                  |
| ------- | -------- | ---------------------------- |
| Command | string   | the command name or pathname |
| Args    | []string | the args                     |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | stream | `-`    |
| O   | stream | `-` |

## Related
