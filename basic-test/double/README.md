# Test Double

## Dummy

คือ การ mock parameter ให้ compiler รันผ่าน โดยของด้านในไม่มี logic ใดๆ

## Stubs

คือ เราจะ provide ค่าบ้างอย่างคืนไป

## Spies

คือ เอาไปดูบ้างอย่างคล้ายๆ Stubs แต่ทำอย่างอื่นได้ด้วย spy จะมีการ implement property เพิ่มเติมเพื่อให้เราสามารถทดสอบว่า spy นั้นถูกเรียกหรือไม่

spy จะมีการ implement property เพิ่มเติมเพื่อให้เราสามารถทดสอบว่าของที่รับมามีการเปลี่ยนแปลงหรือไม่

## Fakes

คือ การทำ fake logic

## Mocks

คือ Stubs + Spies + Fakes
