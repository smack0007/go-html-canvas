module main

go 1.17

require (
    cards v1.0.0
    types v1.0.0
)

replace (
    cards v1.0.0 => ../cards
    types v1.0.0 => ../types
)