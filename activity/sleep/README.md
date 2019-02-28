---
title: Sleep
weight: 4615
---

# Log
This activity allows you to pause flow execution for given time interval.

## Installation
### Flogo Web
https://github.com/vijaynalawade/flogo-contrib/activity/sleep
### Flogo CLI
```bash
flogo install github.com/vijaynalawade/flogo-contrib/activity/sleep
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
    {
          "name": "interval",
          "type": "integer"
    },
    {
          "name": "intervalType",
          "type": "string",
          "allowed": ["Millisecond", "Second", "Minute"],
          "value": "Millisecond"
    }
  ]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| interval    | True     | Sleep time interval |
| intervalType| True     | Interval type. Supported types are - Millisecond, Second, Minute |
## Examples
The below example logs a message 'test message':

```json
{
      "id": "SleepActivity",
      "name": "SleepActivity",
      "activity": {
       "ref": "github.com/vijaynalawade/flogo-contrib/activity/sleep",
       "settings": {},
       "input": {
        "interval": 30,
        "intervalType": "Second"
       }
      }
}
```