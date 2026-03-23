# CreateProduct

## รายละเอียด

API นี้ใช้สำหรับสร้างสินค้าใหม่ในระบบ

เมื่อส่งข้อมูลชื่อสินค้าและราคามาถูกต้อง ระบบจะบันทึกลงฐานข้อมูลและคืนข้อมูลสินค้าที่ถูกสร้างกลับมา

## Endpoint

```http
POST /api/product
```

## Method

- `POST`

## Path

- `/api/product`

## Request Parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| name | String | Yes | ชื่อสินค้า |
| price | Integer | Yes | ราคาสินค้า |

## Request Example

```bash
curl --location --request POST 'http://localhost:3000/api/product' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Keyboard",
  "price": 890
}'
```

## Response Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| id | Integer | ID ของสินค้าที่ถูกสร้าง |
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
| 400 | invalid request body | request body ไม่เป็น JSON ที่ bind ได้ |
| 500 | failed to create product | ไม่สามารถบันทึกข้อมูลลงฐานข้อมูลได้ |

## Error Example

```json
{
  "error": "invalid request body"
}
```