# Metatag details

## Details:

#### @name

> Id parameter to reference the entity.

Names can be used to identify the entity. Multiple names can be set, which can be later referenced by other notes. Please chose unique names or use the complete notepath as the final ID.

```
@name <name>[, ?<names>]
@name example.com, 127.0.0.1, entrypoint-php-system, internal.hostname
@name recon-1
```

#### @ref

> Link between entities

Creates a connection between entities. The keyword `this` can be used, either as the sender, message or recipiant, based on its own @entity.

```
@ref (<node>)->[<action>]->(<node>)
@ref <node>

@ref (this)->[connects]->(other-node)
@ref (node-b)->[connect-to]->(this)
@ref (node-c)->[likes]->(node-d)

# only if @entity relation
@ref (node-a)->[this]->(node-b)

# only if @entity info
@ref parent-node
```

#### @tags

> Give your notes some context and group them

Tags can be used to group notes and give them context. It is helpful for later searches. Tags are separated by , can be have (almost) multiple depth layers.

```
@tags tag1
@tags tag2, tag3, tag4
@tags multi/path/tag/are, also/possible,
```

#### @todo

> Indicates the start of a todo list

connect a todo with a super topic and create todos seperated by newlines. Todos can be define as the todo.txt standard and can contain projects, tags and times and other information. Please see the todo section for additional details.

```
@todo topic-for-todo
task 1
task 2
task 3

@todo other-topic                                                       topic 1
```

#### @icon

> Lets you specify your own icon

Chose your own icons for your entities

```
@icon <url>
@icon http://example.com/icon/info.svg
```

#### @entity

> Specifiy which entity a note should be.

```
@entity <type>
types:

@entity info
# defines the relationship between nodes.
@entity relationship

```

**types**

```
database, node, globe, cserver, site, user, client, dc , win, server
```

#### @label

> Graph label for node

Label can be used to display the node in the graph. Labels can be used multiple times.

```
@label <display>
@label display-name
```
