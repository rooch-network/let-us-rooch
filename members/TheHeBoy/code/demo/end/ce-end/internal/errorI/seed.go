package errorI

import e "errors"

var SeedAddressLocked = e.New("address is locked")
var SeedNotFound = e.New("seed not found")
