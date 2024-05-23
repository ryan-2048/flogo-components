# Sleep

> This activity allows you to pause flow execution for given time interval.

## Installation

```bash
flogo install github.com/ryan-2048/flogo-components/activity/sleep
```

Link for flogo web:

```bash
https://github.com/ryan-2048/flogo-components/activity/sleep
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

## Inputs

| Setting     | Required | Description |
|:------------|:---------|:------------|
| interval    | True     | Sleep time interval |
| intervalType| True     | Interval type. Supported types are - Millisecond, Second, Minute |

## Ouputs