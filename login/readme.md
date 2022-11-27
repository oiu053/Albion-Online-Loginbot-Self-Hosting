# AlbionOnline registration Bot

## Featues

1. Register and link your Ingameacc to you Discordaccount
    - If the IGN is blacklisted, the registeryprocess will get cancled!
    - If the IGN dosn't exists, the registeryprocess will get cancled!

    - applying roles 
        - applying roles based on the Adminlist, Friendlist
        - applying every Memberrole based on the /registered_roles commands
        - creating Guildroles one Step higher than the lowest Memberrole
            - Colour Based on the /guild-colour command. (default: randomly)
            if the player dosn't have a guild the Guild will be equal as "/"
        - creating Allianceroles one Step higher than the lowest Memberrole
            - Colour Based on the /alliance-colour command. (default: randomly)
            if the player dosn't have an alliance the alliance will be equal as "-"


    - changing the server nickname!
        - Format: [Guildname]Ingamename
        - Can't change the nickname of higher ranked Users!
    
    - denying others to register to this IGN
    - denying the registered to register to another IGN 

2. List every registered Player 
3. List every unregistered Player from one ingameguild

4. Blacklist IGN's
    - add IGN's to the Blacklist
    - removing IGN's from the Blacklist

5. Clearachannel:
    - clear / checks if the last 100 Messages have Pinned Messages. Deleates evry non Pinned message of the last 100 Messages

    -clear-all / clears every Message in the range of the last 100 Messages

6. PermissionManagement:
    - some roles could get bounded to the IGN's
        - Based on Frind and or Adminpermissions
            - Pay attention that you don't have the sameroles in The Friendroles as in the Adminroles

7. Discord Timestamps:
    - timestamp-data: 
        - Uses the information of the options to create a discord timestamp who can be coppyed for the use in your own Messages

    - Countdown: 
        - creates a Timestamp in some time. (Connected to your inverted data)

## Permissions:
    - Serverowner: 
        - Blacklist everyone except for himself

        - List every single Ingamename registered (/get-registered-players)
        - List every single Ingamename who isn't registered to the spezific ingame Guild (/get-unregisteredplayers)

    - Only Persons withthe permission to manage the discordserver have the Possibility to Manage the Adminpermissions:
        - adminrrole add
        - adminrole remove
        - ign to adminnames add
        - ign from adminnames remove

        - Friendrole add
        - Freiendrole Remove

    - Persons with adminroles have the possibility to use:
        - Set the registerchannel (the command /register can only be used in this channel!)
        - Remove the registerchannel (the command /register can be used everywhere!)

        - RegisterRole add

        - Change colour of created Guildroles

        - the Clear command (the command /clear reads the last 100 Messages and only deleates them if they arn't pinned! PinnedMessages will get ignored!!!)

        - ign to friendnames add
        - ign from friendnames remove

        - blacklist everyone without anadminrole or Manage Server Permissions
    
    - everyone:
        - register
        - unregister
        - check if an IGN is blacklisted
        - create Timestamps

## Help / Bugs: you can ask in the Discord: 

- https://discord.gg/HPTz2kKTbv

