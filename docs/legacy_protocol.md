# Orsum Inflandi II
## 4 in a row with GO and React + Electron

### Communication protocol

#### Connection initialization
| Client I                                | Server                        |
|----------------------------------------:|:------------------------------|
|          connection:connect:[clientName]|                               |
|                                         |success:accepted:[publicKey]   |
|          connection:keyExchange:[aesKey]|                               |
|                     Communication is now|encrypted with aes             |
|                                         |connection:keyExchange:success |
|                        info:requestGames|                               |
|                                         |success:requested:[listOfGames]|

#### Create game
| Client I | Server |
|---:|:---|
| server:newgame:[gameName]||
||success:created|

#### Join game
| Client I | Server |
|---:|:---|
| game:join:[gameName]||
||success:joined:[playerNr]|

#### Set gamestone
| Client I | Server | Client II |
|---:|---|:---|
| game:setstone:[rowNr]|||
||<= success:set||
||game:setstone[rowNr] => ||

#### Send chat message
| Client I | Server | Client II |
|---:|---|:---|
| chat:send:[message]|||
||<= chat:send:success||
||chat:[senderName][message] => ||