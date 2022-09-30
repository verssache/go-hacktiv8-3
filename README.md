# Assignment 3

Dalam tugas 3 saya menggunakan go-cron dikarenakan bot ini akan jalan terus menerus dengan satuan waktu yaitu 15 detik
Karena cron job sudah begitu terkenal dalam dunia pemrograman maka saya coba cari package cron dalam golang dan menemukan go-cron
https://github.com/go-co-op/gocron

Pada programnya saya mendefiniskan objek go-cron yang dimana menggunakan timezone jakarta kemudian saya menggunakan perintah 
s.Every(5).Seconds().Do(func(){ ... })
seperti yang tertera pada dokumentasi go-cron

Nah karena perlu menjalankan server menggunakan gin maka program go-cron saya lakukan secara asynchronus jadi tidak perlu blocking program
Terima kasih
