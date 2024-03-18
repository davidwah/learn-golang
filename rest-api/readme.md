# Rest API

## Request and Response API

* Create Order (request):
```json
{
    "orderedAt": "2021-10-06T16:53:27.675931+07:00",
    "customerName": "Test",
    "items": [
        {
            "itemCode": "Item1",
            "description": "ItemDescription",
            "quantity": 1
        },
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        }
    ]
}
```
* Create Order (response):
```
{
    "id": 1,
    "orderedAt": "2021-10-06T16:53:27.675931+07:00",
    "customerName": "Test",
    "items": [
        {
            "itemCode": "Item1",
            "description": "ItemDescription",
            "quantity": 1
        },
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        }
    ]
}
```
* Get Orders (response):
```
[
    {
        "id": 1,
        "orderedAt": "2021-10-06T16:53:27.675931+07:00",
        "customerName": "Test",
        "items": [
            {
                "itemCode": "Item1",
                "description": "ItemDescription",
                "quantity": 1
            },
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            }
        ]
    },
    {
        "id": 2,
        "orderedAt": "2021-10-06T16:53:27.675931+07:00",
        "customerName": "Test",
        "items": [
            {
                "itemCode": "Item1",
                "description": "ItemDescription",
                "quantity": 1
            },
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            }
        ]
    }
]
```
* Update order (request):
```
{
    "orderedAt": "2021-10-06T16:53:27.675931+07:00",
    "customerName": "Test update",
    "items": [
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        },
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        }
    ]
}
```
* Update order (response):
```
{
    "id": 2,
    "orderedAt": "2021-10-06T16:53:27.675931+07:00",
    "customerName": "Test update",
    "items": [
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        },
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        }
    ]
}
```
* Delete Order (response):
```
Success delete
```