# Go Fiber Product API

Created by: Nantadet Wongchaiya  
GitHub: [@Nantadet](https://github.com/Nantadet)

REST API สาหรับจัดการข้อมูลสินค้า (Product) ด้วย Go, Fiber, GORM และ MySQL

รายละเอียดการใช้งาน CRUD ดูได้ที่ [doc](doc/api.md)

## Clone Project

```bash
git clone https://github.com/Nantadet/Go-FrameWork.git
cd Go-FrameWork
```

## Run Project

ก่อนรันโปรเจ็กต์ ควรตรวจสอบว่า:

- MySQL กาลังรันอยู่
- มี database ชื่อ `mydb`
- ค่า config ใน [config/database.config.go](config/database.config.go) ตรงกับเครื่องที่ใช้งาน

### Run normally

```bash
go run main.go
```

### Run with nodemon

```bash
npm install
npm run dev
```

เซิร์ฟเวอร์จะรันที่:

```text
http://localhost:3000
```
