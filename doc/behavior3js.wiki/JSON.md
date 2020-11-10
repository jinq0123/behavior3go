## JSON Format For BTs in **Behavior3JS**

### Tree Specification

**Tree items**:

- **title**: the tree title, useful for identification on editor and debug.
- **description**: the tree description, useful for identification on editor 
  and debug.
- **root**: the root ID (preferably a string), must be a unique identifier.
- **display**:
    - **camera_x**: position X of the camera in canvas.
    - **camera_y**: position Y of the camera in canvas.
    - **camera_z**: zoom factor of the camera in canvas.
    - **x**: position X of the root block in the canvas.
    - **y**: position Y of the root block in the canvas.
- **properties**: object with `<key, value>` pairs.
- **nodes**: object with `<key, value>` pairs, where `key` must be the node ID,
  and `value` must be the node specification.

**Example**:

    {
        "title"       : <String>,
        "description" : <String>,
        "root"        : <String(ID)>,
        "display"     : {
            "camera_x" : <Float>,
            "camera_y" : <Float>,
            "camera_z" : <Float>,
            "x"        : <Float>,
            "y"        : <Float>
        }
        "properties"  : {
            <String> : <String or Number>,
            ...
        },
        "nodes"       : {
            <String(ID)>: <Node Specification>,
            ...
        }
    }

### Node Specification

**Node items**:

- **id**: the node ID (preferably a string), must be a unique identifier.
- **title**: the node title, useful for identification on editor and debug.
- **description**: the node description, useful for identification on editor 
  and debug.
- **display**:
    - **x**: position X of the block in the canvas.
    - **y**: position Y of the block in the canvas.
- **children**: an array of node IDs, only for composite nodes.
- **child**: a node ID, only for decorator nodes.
- **parameters**: object with `<key, value>` pairs. This is needed to specify 
  obligatory parameters when loading trees from JSON.
- **properties**: object with `<key, value>` pairs.

**Example**:

    {
        "id"          : <String(ID)>,
        "name"        : <String>,
        "title"       : <String>,
        "description" : <String>,
        "children"    : [<String(ID)>, ...], // Only for composites
        "child"       : <String(ID)>, // Only for decorators
        "display"     : {
            "x" : <Float>,
            "y" : <Float>,
        },
        "parameters"  : {
            <String> : <String or Number>,
            ...
        },
        "properties"  : {
            <String> : <String or Number>,
            ...
        }
    }


### Example Tree

This example was exported from editor.

    {
        "title": "A Behavior Tree",
        "description": "",
        "root": "2b1c6eda-f0a6-4021-81a1-f1555612d57f",
        "display": {
            "camera_x": 514.5,
            "camera_y": 479,
            "camera_z": 1,
            "x": 0,
            "y": 0
        },
        "properties": {},
        "nodes": {
            "f0254efb-9861-4e0d-822f-b5ca96457dbf": {
                "id": "f0254efb-9861-4e0d-822f-b5ca96457dbf",
                "name": "Limiter",
                "title": "Limiter",
                "description": "",
                "display": {
                    "x": 416,
                    "y": -32
                },
                "parameters": {
                    "maxLoop": 5
                },
                "properties": {},
                "child": "8ad1967e-ea55-49b9-993e-4d65ce3ee459"
            },
            "93edb2f3-47df-4559-99aa-1492bc032ae9": {
                "id": "93edb2f3-47df-4559-99aa-1492bc032ae9",
                "name": "Wait",
                "title": "Wait",
                "description": "",
                "display": {
                    "x": 416,
                    "y": 32
                },
                "parameters": {
                    "milliseconds": 1000
                },
                "properties": {}
            },
            "2b1c6eda-f0a6-4021-81a1-f1555612d57f": {
                "id": "2b1c6eda-f0a6-4021-81a1-f1555612d57f",
                "name": "MemSequence",
                "title": "MemSequence",
                "description": "",
                "display": {
                    "x": 208,
                    "y": 0
                },
                "parameters": {},
                "properties": {},
                "children": [
                    "f0254efb-9861-4e0d-822f-b5ca96457dbf",
                    "93edb2f3-47df-4559-99aa-1492bc032ae9"
                ]
            },
            "8ad1967e-ea55-49b9-993e-4d65ce3ee459": {
                "id": "8ad1967e-ea55-49b9-993e-4d65ce3ee459",
                "name": "Custom Node",
                "title": "Custom Node",
                "description": "",
                "display": {
                    "x": 624,
                    "y": -32
                },
                "parameters": {},
                "properties": {
                    "Some": "Property"
                }
            }
        }
    }