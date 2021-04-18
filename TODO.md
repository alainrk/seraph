# Random things wondering to do

## High priority
- bug on retry password, an empty vault is initialized
- edited BUT NOT saved fields, are correctly not stored on crypted file, but still remain in app context => they have to be cleared, otherwise visible to the user when getsecret is called

## Medium priority
- handle a bak file to avoid corrupted saving?
- preview editing value

## Low priority
- custom fields?
- better const (mainly steps ones) handling, it's shit now

# Doing

# Done
- infinite loop stuff
- single session => multiple vault handling?
- infinite loop of editing/adding
- multiple vault handling same session
- [getSecret] showing results has to be prettier
- [getSecret] there has to be a any-key waiting press to keep the results on the screen and on key press user can get back to menu
- getSecret step => if no keys have to go back automatically
- handle unique name => now same name overrides the previous one
- edit secret