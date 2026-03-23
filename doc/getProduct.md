# GetProduct

## รายละเอียด

API นี้ใช้สำหรับดึงข้อมูลสินค้า 1 รายการตาม `id`

เหมาะสำหรับ:

- เปิดหน้ารายละเอียดสินค้า
- ใช้โหลดข้อมูลก่อนแก้ไข
- ตรวจสอบว่าสินค้าที่ต้องการมีอยู่จริงหรือไม่

## Endpoint

```http
GET /api/product/:id
```

## Method

- `GET`

## Path

- `/api/product/:id`

## Request Parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| id | Integer | Yes | ID ของสินค้าที่ต้องการดึง |

## Request Example

```bash
curl --location --request GET 'http://localhost:3000/api/product/1'
```

## Response Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| id | Integer | ID ของสินค้า |
| name | String | ชื่อสินค้า |
| price | Integer | ราคาสินค้า |

## Response Example

```json
{
  "id": 1,
  "name": "Keyboard",
  "price": 890
}
```

## Error Codes

| Code | Message | Description |
| --- | --- | --- |
| 400 | invalid product id | ค่า `id` ไม่ถูกต้องหรือมีค่าน้อยกว่า 1 |
| 404 | product not found | ไม่พบสินค้าตาม `id` ที่ระบุ |
| 500 | failed to fetch product | เกิดข้อผิดพลาดระหว่าง query database |

## Error Example

```json
{
  "error": "product not found"
}
```