## What is a Behavior Tree?

Well, here is not the best source to discover that! We assume that you already 
know what BTs are to use **Behavior3JS**, but if don't, just follow these links
and you will be good:

- [An Introduction to Behavior Trees (part 1)][tutorial_1]
- [An Introduction to Behavior Trees (part 2)][tutorial_2]
- [An Introduction to Behavior Trees (part 3)][tutorial_3]
- [Implementing a Behavior Tree (part 1)][tutorial_4] - Is a preview for this 
  library.
- [Implementing a Behavior Tree (part 2)][tutorial_5]
- [Behavior Trees for AI - How They Work][tutorial_6]
- [Second Generation Behavior Trees][tutorial_7]
- [Behavior Trees Overview][tutorial_8]


## What is Behavior3JS

**Behavior3JS** is a Behavior Tree library written in JavaScript. It provides 
structures and algorithms that assist you in the task of creating intelligent
agents for your game or application. Check it out some features of Behavior3JS:

- Based on the work of [(Marzinotto et al., 2014)][marzinotto], in which they
  propose a **formal**, **consistent** and **general** definition of Behavior
  Trees;

- **Optimized to control multiple agents**: you can use a single behavior 
  tree instance to handle hundreds of agents;

- It was **designed to load and save trees in a JSON format**, in order to 
  use, edit and test it in multiple environments, tools and languages;

- A **cool visual editor** which you can access online;

- Several **composite, decorator and action nodes** available within the 
  library. You still can define your own nodes, including composites and 
  decorators;

- **Completely free**, the core module and the visual editor are all published
  under the MIT License, which means that you can use them for your open source
  and commercial projects;

- **Lightweight**, only 11.5KB!

Visit http://behavior3js.guineashots.com to know more!


## Core Classes

This library include the following core structures...

- **BehaviorTree**: the structure that represents a Behavior Tree;
- **Blackboard**: represents a "memory" in an agent and is required to to 
  run a `BehaviorTree`;
- **Composite**: base class for all composite nodes;
- **Decorator**: base class for all decorator nodes;
- **Action**: base class for all action nodes;
- **Condition**: base class for all condition nodes;
- **Tick**: used as container and tracking object through the tree during 
  the tick signal;
- **BaseNode**: the base class that provide all common node features;


## Nodes

**Composite Nodes**: 

- Sequence;
- Priority;
- MemSequence;
- MemPriority;


**Decorators**: 

- Inverter;
- Limiter
- MaxTime;
- Repeater;
- RepeaterUntilFailure;
- RepeaterUntilSuccess;


**Actions**:

- Succeeder;
- Failer;
- Error;
- Runner;
- Wait.


[tutorial_1]: http://guineashots.com/2014/07/25/an-introduction-to-behavior-trees-part-1/
[tutorial_2]: http://guineashots.com/2014/08/10/an-introduction-to-behavior-trees-part-2/
[tutorial_3]: http://guineashots.com/2014/08/15/an-introduction-to-behavior-trees-part-3/
[tutorial_4]: http://guineashots.com/2014/09/24/implementing-a-behavior-tree-part-1/
[tutorial_5]: http://guineashots.com/2014/10/25/implementing-a-behavior-tree-part-2/
[tutorial_6]: http://www.gamasutra.com/blogs/ChrisSimpson/20140717/221339/Behavior_trees_for_AI_How_they_work.php
[tutorial_7]: http://aigamedev.com/insider/tutorial/second-generation-bt/
[tutorial_8]: http://aigamedev.com/open/article/bt-overview/
[marzinotto]: http://www.csc.kth.se/~miccol/Michele_Colledanchise/Publications_files/2013_ICRA_mcko.pdf