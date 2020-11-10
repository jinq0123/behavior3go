## Common Features

Before creating your own nodes, please read the [General Guidelines](Core 03 General Guidelines)

## Composite

When creating composite nodes, you will need to propagate the tick signal to
the children nodes manually. To do that, override the `tick` method and call 
the `_execute` method on all nodes. For instance, take a look into the code of
the Sequence composite node:
    
    // Creates a new class inheriting from Composite
    var Sequence = b3.Class(b3.Composite);

    // Make sure you use the same name of the class
    Sequence.prototype.name = 'Sequence';
        
    // Override the tick function
    Sequence.prototype.tick = function(tick) {
        // Iterates over the children
        for (var i=0; i<this.children.length; i++) {
            // Propagate the tick
            var status = this.children[i]._execute(tick);
            if (status !== b3.SUCCESS) {
                return status;
            }
        }
        return b3.SUCCESS;
    }

Notice that, you don't need to define a children variable in the composite 
node, because this is done inside the Composite class (including reading the
settings parameters in initialize and getting the children list). For example,
you can initialize the Sequence node like this:

    var sequence = new b3.Sequence({children:[node1, node2, nodeN]});


## Decorator

When creating decorator nodes, you will need to propagate the tick signal to
the child node manually, just like the composite nodes. To do that, override
the `tick` method and call the `_execute` method on the child node. For 
instance, take a look at how the Inverter decorator is implemented:


    // Creates a new class inheriting from Decorator
    var Inverter = b3.Class(b3.Decorator);

    // Remember to set the name of the node. 
    Inverter.prototype.name = 'Inverter';

    // Override the tick function
    Inverter.prototype.tick = function(tick) {
        if (!this.child) {
            return b3.ERROR;
        }

        // Propagate the tick
        var status = this.child._execute(tick);
        if (status == b3.SUCCESS)
            status = b3.FAILURE;
        else if (status == b3.FAILURE)
            status = b3.SUCCESS;
        return status;
    };

Notice that, you don't need to define a child variable in the decorator 
node, because this is done inside the Decorator class (including reading the
settings parameters in initialize and getting the child instance). For example,
you can initialize the Inverter node like this:

    var inverter = new b3.Inverter({child: node1});


## Action and Condition Nodes

Both nodes are leaves, thus, they don't have children. This simplifies the 
implementation of these kind of nodes. Take a look into the Wait node source to
have an idea:


    // Creates a new class inheriting from Action
    var Wait = b3.Class(b3.Action);
    
    // Remember to set the name of the node. 
    Wait.prototype.name = 'Wait';

    // Sets the parameters variable to tell editor who they are
    Wait.prototype.parameters = {'milliseconds': 0};

    // Override the initialize method, remember to call this method on super
    Wait.prototype.__Action_initialize = Wait.prototype.initialize;
    Wait.prototype.initialize = function(settings) {
        settings = settings || {};

        this.__Action_initialize();
        this.endTime = settings.milliseconds || 0;
    }

    // Override the open method, so it can store the time when the node was
    // open    
    Wait.prototype.open = function(tick) {
        var startTime = (new Date()).getTime();
        tick.blackboard.set('startTime', startTime, tick.tree.id, this.id);
    }

    // Override the tick method
    Wait.prototype.tick = function(tick) {
        var currTime = (new Date()).getTime();
        var startTime = tick.blackboard.get('startTime', tick.tree.id, this.id);
        
        if (currTime - startTime > this.endTime) {
            return b3.SUCCESS;
        }
        
        return b3.RUNNING;
    }