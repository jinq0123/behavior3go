## Use The Helper Function `b3.Class`

All classes in Behavior3JS are created using the meta function `b3.Class`, 
which:

- will return a function, which is used as a class definition;
- will do the prototypical inheritance is another class is provided as 
  argument;
- will add a method `initialize` to the class, which will be the constructor \
  (called automatically by the function created by `b3.Class`);
    - the `initialize` method can receive only a single parameter (usually 
    named as `settings`). This parameter is an object containing the actual 
    parameters. This pattern is commonly used in the jQuery community.

For instance, suppose you want to create a new class for any reason:

    var MyCustomClass = b3.Class();
    var custom = new MyCustomClass(); // will call initialize.

If you want to create a custom node, you must inherit from `b3.Composite`,
`b3.Decorator`, `b3.Condition` or `b3.Action`. So, suppose you want to create
a custom composite node class:

    var CustomComposite = b3.Class(b3.Composite);


## Node Name

When creating a new node class, remember to set its name in prototype context.
The node name is important to store and indentify it (for example, when 
loading the tree from JSON).


## Interface Methods

When creating a new node class, you may want to override one or more of these 
methods:

- **enter**: called every time a node is executed;
- **open**: called only when the node is opened (when a node return 
  `b3.RUNNING`, it will stay opened until it return other value or the tree 
  forces the closing);
- **tick**: the real implementation of the node (e.g., the composite nodes will
  call their children in this method). Notice that, this method must return 
  a state value, such as `b3.SUCCESS`, `b3.FAILURE`, `b3.RUNNING`, or 
  `b3.ERROR`;
- **close**: called when the node return `b3.SUCCESS`, `b3.FAILURE` or 
  `b3.ERROR`, or when the tree force the node to close;
- **exit**: called every time at the end of the node execution.

Notice that, you don't need to call these method in the super class when 
overriding them.

All these methods receives as parameter a tick object. The tick stores:

- a reference to the tree (`tick.tree`);
- a reference to the target object (`tick.target`);
- a reference to the blackboard (`tick.blackboard`);


## Blackboard or "the memory"

You may want to store some variables during the node execution. But, if you 
store these variables into the node object itself, this information will
persist for any agent sharing the tree, probably breaking the execution of all
agents. For this reason, **Behavior3JS** provides a memory structure that may
be unique for each agent in the game.

This memory structure is called `Blackboard` and only have 2 public methods:
`set` and `get`. These methods works in 3 different contexts: global, per tree,
and per node per tree.

Suppose you have two different trees controlling a single object with a 
single blackboard, then:

- In the global context, all nodes will access the stored information. 
- In per tree context, only nodes sharing the same tree share the stored 
  information.
- In per node per tree context, the information stored in the blackboard can
  only be accessed by the same node that wrote the data.
  
The context is selected indirectly by the parameters provided to these 
methods, for example:

    // getting/setting variable in global context
    blackboard.set('testKey', 'value');
    var value = blackboard.get('testKey');
    
    // getting/setting variable in per tree context
    blackboard.set('testKey', 'value', tree.id);
    var value = blackboard.get('testKey', tree.id);
    
    // getting/setting variable in per node per tree context
    blackboard.set('testKey', 'value', tree.id, node.id);
    var value = blackboard.get('testKey', tree.id, node.id);

