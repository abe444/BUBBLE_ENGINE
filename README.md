# [PROTOCOL LAYER: BUBBLE]
Source code for an experimental blogging system written in Go.

## CORE IDEA: 
Blogging is fun. It's even more fun when using simple tools to get articles quickly up and running.  
The BUBBLE engine is meant to be that simple tool that allows for extreme ease of use. 

## CORE COMMANDS:

### Initial setup:
This will walk you through some steps to getting the engine properly configured.
```
$ bubble setup
```

### Create a new article entry:
This will create a new markdown file inside of the entries folder. 
```
$ bubble new [article title]
```

### Create a rough draft:
This will create a markdown file inside of the drafts folder. These files won't get displayed in production. Ever. Perfect for brainstorming.
```
$ bubble draft [draft title]
```

### Archive an entry:
Archiving an article means no future updates will be made to said article.
```
$ bubble archive_entry [article title]
```

### Restart engine:
The engine will need to be restarted after writing **anything** to the entries folder.
```
$ bubble restart
```

## TODO:
- Sort article entries by date format: YYYY-MM-DD_[ENTRY].md.
- Engine setup system.
- Setup tag system.
- Setup Archiving system.
- RSS.

## EXPERIMENTAL:
- Add a reply/comment system that sort of acts as a secondary blog in of itself.



