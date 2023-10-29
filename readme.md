#bookingapi 

STATUS DEV

Saya membuat api untuk membooking kamar dan slot yang dipanggil dengan chairs, api ini bisa diperuntukkan untuk membuat service booking kamar pada hotel dan kursi pada sistem penerbangan

Terdapat dua buah endoint yaitu: 
1. "/booking"
Menggunakan method "get" cara kerja nya adalah endpoint "/booking" akan me return json "Response data yang berisi id, isbooking, userbooking, userphone," yang berisi bebrapa chairs sesuai yang diinginkan oleh admin

2. "booking'?id'"
Menggunakan method "post" yang dimana id pada url mengacu pada chairs yang di pilih yang akan menerima request dari user berupa user_booking atau nama si pembooking dan user_phone, dan status is_booking pada url "/booking" akan di update menjadi status telah di book dan request data pada user akan di simpan di database agar tidak terjadi booking yag lebih dari satu