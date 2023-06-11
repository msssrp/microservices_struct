# mircroservices_struct

microservice คือการที่ให้แต่ละ service มี server และ port แยกเป็นของตัวเองไปตามแต่ละ server อย่างเช่น user server = port 8080 , product server = port 8081 นี่จะเป็นการแยกการทำงานของแต่ละ service

ง่ายๆก็คือเราจะสร้าง server แยกของแต่ละ service และให้แต่ละ service มี server port เป็นของตัวเอง ในที่นี้จะช่วยให้ main server ทำงานไม่หนักและทำหน้าที่แค่จัดการว่า request ไหน ไป port ไหน
