module main

go 1.17

require (
    game v1.0.0
    types v1.0.0
)

replace (
    game v1.0.0 => ../game
    types v1.0.0 => ../types
)