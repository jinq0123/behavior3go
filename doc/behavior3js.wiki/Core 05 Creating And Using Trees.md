## Creating Trees Manually

You can create a tree manually, defining all nodes in the graph in code. For 
example:

    var tree = new b3.BehaviorTree();

    tree.root = new b3.Sequence({children: [
        new b3.Priority({children: [
            new MyCustomNode(),
            new MyCustomNode()
        ]}),
    ]});


## Creating Trees From JSON

You also can create a tree from JSON files, usually created in the editor:

    var tree = new b3.BehaviorTree();

    tree.load({
        'title'       : 'Behavior Tree title'
        'description' : 'My description'
        'root'        : 'node-id-1'
        'nodes'       : {
            'node-id-1' : {
                'name'        : 'Priority',
                'title'       : 'Root Node', 
                'description' : 'Description', 
                'children'    : ['node-id-2', 'node-id-3'], 
            },
            ...
        }
    })


## Using Behavior Trees

Usually, you want to define a target object to be controlled by the behavior 
tree. This target object can be anything: a function, an object, a string, a 
DOM element; this depends on your application.

You will also need to create a blackboard object (the memory), usually one for 
each target.

With the target and the blackboard objects. You just need to "tick" the tree 
and the magic will happen:

    var target = new Image(); // The target will be an Image element
    var blackboard = new b3.Blackboard();

    tree.tick(target, blackboard);
