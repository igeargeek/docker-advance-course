# ประเภทของ Docker Network
โดยปกติแล้ว Docker จะสามารถแบ่ง Network ออกได้เป็น 5 ประเภทประกอบด้วย (ไม่รวม custom และพวก 3rd party)

 - None คือการรัน container ในระบบปิด ที่ไม่สามารถออก network ภายนอกได้ ปกติเรามักสร้าง None network ไว้เชื่อมกับ custom network driver ของเราเอง
 - Bridge คือการรัน container ที่สามารถเชื่อมต่อไปยัง host network หรือถ้าพูดง่ายๆก็คือ าสามารถออก internet ได้ โดยถ้าเราสั่งรัน container โดยไม่ระบุ network driver ก็จะได้ Bridge มาเป็น default network เสมอ ซึ่ง Bridge จะมี Subnet และ Gateway เป็นของตัวเอง
 - Host คือการรัน container บน host network เลย จะไม่มีการแยก network ระหว่าง container และ docker host เหมือนใน Bridge
 - Overlay คือการรัน container ที่มีการคุยกันข้าม docker daemons (host network) ถ้ามีโจทย์แนวๆการทำ distributed network และต้องการให้ container สามารถคุยกันได้ บน host ที่ต่างกัน Overlay ก็จะเป็น network อีกประเภทหนึ่ง ที่เราควรศึกษาเพิ่มดู
 - MacVlan คือการรัน container ที่สามารถทำให้ container วิ่งไปยัง router ที่ต้องการได้ โดยไม่ต้องผ่าน host network


![alt text](https://www.ostechnix.com/wp-content/uploads/2017/01/Portainer-Chromium_019.png)
