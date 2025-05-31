-- Insert Categories
INSERT INTO tbl_kategori_produk (id, nama_kategori_produk) VALUES
('1', 'Fixed Satmobile'),
('2', 'Land Satmobile');

-- Insert Products
INSERT INTO tbl_produk (id, kategori_produk_id, nama_produk, deskripsi_produk, gambar_produk, spesifikasi_produk) VALUES
('prod-fixed', '1', 'Fixed Satmobile', 'Data Standar Tanpa Batas, Prioritas Jaringan, Dukungan Prioritas', 'http://example.com/images/fixed.png',
'Layanan internet berkecepatan tinggi berbasis satelit dengan sistem kuota bulanan, ditawarkan dalam bentuk langganan dengan biaya tetap setiap bulan (Monthly Regular Charge) dan minimal kontrak 12 bulan. Selain itu, perangkat Customer Terminal untuk koneksi tetap tersedia dengan skema pembelian (One Time Charge). Solusi ini dirancang untuk memenuhi kebutuhan internet bagi bisnis skala kecil, menengah, hingga besar, terutama di wilayah yang belum terjangkau layanan internet , atau berada di daerah 3T (Terpencil, Terdepan, dan Tertinggal).'),

('prod-land', '2', 'Land Satmobile', 'Data Daratan Tanpa Batas, Dalam Perjalanan + Penggunaan Laut, Prioritas Jaringan, Dukungan Prioritas', 'http://example.com/images/land.png',
'Layanan internet berkecepatan tinggi berbasis satelit dengan sistem kuota bulanan, ditawarkan dalam bentuk langganan dengan biaya tetap setiap bulan (Monthly Regular Charge) dan minimal kontrak 12 bulan. Selain itu, perangkat Customer Terminal untuk koneksi tetap tersedia dengan skema pembelian (One Time Charge). Solusi ini dirancang untuk kebutuhan akses internet bisnis skala kecil, menengah, dan perusahaan besar khususnya sangat diperlukan saat sedang berpergian di kendaraan bergerak dan kapal.');

-- Insert Devices 
INSERT INTO tbl_perangkat (id, produk_id ,kategori_produk_id, nama_produk, deskripsi_produk, harga_produk, gambar_produk) VALUES
('dev-001', 'prod-fixed' ,'1', 'Customer Terminal Enhanched', 'Data Standar Tanpa Batas, Prioritas Jaringan, Dukungan Prioritas', 1000000, '-'),
('dev-002', 'prod-fixed' ,'1', 'Customer Terminal Adaptive', 'Data Standar Tanpa Batas, Prioritas Jaringan, Dukungan Prioritas', 3000000, '-'),
('dev-003', 'prod-land','2', 'Customer Terminal Adaptive', 'Data Daratan Tanpa Batas, Dalam Perjalanan + Penggunaan Laut, Prioritas Jaringan, Dukungan Prioritas', 3000000, '-');

-- Insert Services
INSERT INTO tbl_layanan (id, produk_id, nama_layanan, harga_layanan) VALUES
-- Fixed Satmobile
('lay-001', 'prod-fixed', '50 GB', 750000),
('lay-002', 'prod-fixed', '1TB', 1500000),
('lay-003', 'prod-fixed', '3TB', 2500000),
('lay-004', 'prod-fixed', '5TB', 3000000),
-- Land Satmobile
('lay-005', 'prod-land', '50 GB', 750000),
('lay-006', 'prod-land', '1TB', 1500000),
('lay-007', 'prod-land', '5TB', 3000000);

-- FAQs for kategori_produk_id = 1 (Fixed Satmobile)
INSERT INTO tbl_faq_kategori_produk (id, kategori_produk_id, pertanyaan, jawaban) VALUES
('faq-1', 1, 'Apakah ada batasan data?',
'Semua paket bisnis hadir dengan data standar tanpa batas dan jumlah data prioritas yang dapat dikonfigurasi. Pilihan paket tercantum di bawah ini: <br>Paket Fixed Site: 1TB, 2TB, 6TB, 40GB<br> Data Fixed Site mencakup kecepatan yang lebih tinggi dan Fixed Site jaringan teratas, memastikan cakupan terbaik untuk bisnis Anda'),

('faq-2' ,1, 'Apa yang terjadi jika saya melebihi batas data?',
'Paket Fixed Site akan dikembalikan ke laju data standar tanpa batas.'),

('faq-3', 1, 'Adakah solusi bila kuota data prioritas habis?',
'Anda dapat melakukan upgrade dan downgrade paket data eksisting untuk dapat menggunakan data prioritas kembali, dengan billing cycle bersesuaian. Lebih detail silahkan hubungi Sales Support terkait.'),

('faq-4', 1, 'Apakah saya bisa mengubah paket layanan saya?',
'Ya, paket layanan dapat diubah di dalam akun satmobile Anda. Informasi lebih lanjut tersedia hubungi kami.'),

('faq-5', 1, 'Dapatkah saya menjeda paket layanan saya? Bagaimana cara menjeda/membatalkan layanan?',
'Anda bisa menjeda/membatalkan layanan kapan saja pada paket layanan Fixed Site. Mekanisme penalti bagi pelanggan yang berhenti berlangganan sebelum masa kontrak berakhir, akan dikenakan full subscription sesuai bulan sisa masa kontrak. Informasi lebih lanjut tersedia hubungi kami.'),

('faq-6', 1, 'Dapatkah saya berbagi paket layanan di antara beberapa satmobile Kit?',
'Tidak, setiap satmobile Kit membutuhkan paket layanannya sendiri.'),

('faq-7', 1, 'Apa yang terjadi jika saya perlu memindahkan satmobile saya ke lokasi yang berbeda, tetapi tidak memerlukannya untuk berpindah tempat?',
'Pelanggan dapat mengubah alamat layanan tetap mereka di akun mereka di dalam negara yang sama, tetapi layanan tidak dijamin. Jika Anda berencana untuk sering memindahkan satmobile Anda (lebih dari beberapa kali dalam setahun), kami sarankan untuk beralih ke paket layanan Mobility.'),

('faq-8', 1, 'Berapa banyak pengguna/perangkat yang dapat didukung oleh satu satmobile Kit ?',
'Kami merekomendasikan sekitar 20 perangkat per High Performance Kit, namun hal ini dapat bervariasi tergantung pada penggunaan jaringan.'),

('faq-9', 1, 'Apakah satu satmobile kit dapat digunakan oleh beberapa gedung / lokasi?',
'Router WiFi satmobile mempunyai jangkauan sekitar 2.000 kaki persegi dan direkomendasikan untuk sekitar 20 pengguna. Jika lokasi Anda berada dalam jangkauan dan jumlah pengguna tersebut, maka satu satmobile kit dapat digunakan untuk beberapa lokasi.'),

('faq-10', 1, 'Seberapa handal koneksinya?',
'Layanan satmobile Business Services memiliki Service Availability di atas 98% uptime dengan redundansi jalur dan switching antar satelit.'),

('faq-11', 1, 'Dapatkah saya menggunakan satmobile sebagai opsi pencadangan / failover untuk bisnis saya?',
'Ya, ini dimungkinkan dan banyak bisnis saat ini menggunakan satmobile sebagai koneksi cadangan sekunder atau tersier.'),

('faq-12', 1, 'Seberapa cepat saya bisa mulai menggunakan layanan satmobile?',
'Kit satmobile biasanya memiliki waktu pengiriman 1-2 minggu dan pemasangannya hanya memakan beberapa jam.'),

('faq-13', 1, 'Bagaimana cara melakukan pemesanan dalam jumlah besar?',
'Anda dapat melakukan pembelian satmobile dalam jumlah besar dengan sebelumnya melakukan konfirmasi ke sales support kami.'),

('faq-14', 1, 'Apa saja skema layanan yang disediakan untuk pilihan bisnis saya?',
'Saat ini kami menawarkan User Terminal dengan skema Jual Putus dan Managed Service beserta Service Data Plan dengan sewa layanan minimal 12 bulan.');

-- FAQs for kategori_produk_id = 2 (Land Satmobile)
INSERT INTO tbl_faq_kategori_produk (id, kategori_produk_id, pertanyaan, jawaban) VALUES
('faq-15', 2, 'Apakah ada batasan data?', 
'Semua paket Land Satmobile memiliki data standar tanpa batas dan opsi data prioritas yang dapat disesuaikan. Pilihan layanan mencakup: <br>50GB, 1TB, dan 5TB dengan prioritas jaringan tinggi, dirancang untuk penggunaan dalam kendaraan bergerak dan laut.'),

('faq-16', 2, 'Apa yang terjadi jika saya melebihi batas data?', 
'Jika batas data prioritas terlampaui, layanan akan beralih ke data standar tanpa batas, tetap menjaga konektivitas meskipun dengan kecepatan standar.'),

('faq-17', 2, 'Adakah solusi bila kuota data prioritas habis?', 
'Ya, Anda dapat meningkatkan paket data sesuai kebutuhan langsung melalui akun Anda atau dengan menghubungi Sales Support.'),

('faq-18', 2, 'Apakah saya bisa mengubah paket layanan saya?', 
'Tentu, Anda dapat menyesuaikan paket layanan Anda kapan saja sesuai kebutuhan perjalanan atau lokasi Anda.'),

('faq-19', 2, 'Dapatkah saya menjeda atau membatalkan layanan?', 
'Anda dapat menghentikan atau menjeda layanan kapan pun. Namun, apabila dilakukan sebelum masa kontrak 12 bulan berakhir, akan dikenakan biaya penalti sesuai sisa kontrak.'),

('faq-20', 2, 'Apakah saya bisa berbagi layanan antar kendaraan?', 
'Tidak. Setiap perangkat satmobile memerlukan paket layanannya sendiri, terutama untuk memastikan kualitas jaringan dalam pergerakan.'),

('faq-21', 2, 'Bagaimana jika saya berpindah lokasi atau jalur laut?', 
'Layanan Land Satmobile dirancang untuk mengikuti pergerakan Anda. Anda tetap dapat terhubung di berbagai lokasi selama dalam area cakupan satelit.'),

('faq-22', 2, 'Berapa banyak perangkat yang bisa terkoneksi dengan satu unit?', 
'Satu unit High Performance Kit direkomendasikan untuk sekitar 20 perangkat, namun hal ini tergantung dari intensitas pemakaian masing-masing.'),

('faq-23', 2, 'Bisakah digunakan di kapal atau kendaraan berat?', 
'Ya. Land Satmobile sangat cocok untuk kendaraan berat, kapal, atau armada transportasi yang memerlukan koneksi stabil dan mobile.'),

('faq-24', 2, 'Apakah koneksi tetap stabil saat bergerak?', 
'Ya. Dengan prioritas jaringan dan satelit multi-jalur, koneksi akan tetap stabil saat kendaraan bergerak atau berada di laut.'),

('faq-25', 2, 'Dapatkah Land Satmobile menjadi koneksi cadangan?', 
'Banyak pengguna bisnis menggunakan layanan ini sebagai koneksi cadangan saat koneksi utama tidak tersedia.'),

('faq-26', 2, 'Seberapa cepat layanan bisa digunakan setelah pembelian?', 
'Pengiriman unit memakan waktu 1–2 minggu, dan dapat langsung digunakan setelah instalasi sederhana yang memakan waktu beberapa jam.'),

('faq-27', 2, 'Bagaimana memesan dalam jumlah besar untuk armada?', 
'Untuk kebutuhan pembelian massal bagi armada atau perusahaan, silakan hubungi tim Sales kami untuk pengaturan dan penawaran khusus.'),

('faq-28', 2, 'Apa saja skema layanan yang tersedia?', 
'Kami menyediakan layanan dengan skema Jual Putus dan Managed Service. Kontrak minimal 12 bulan dan berbagai pilihan data sesuai kebutuhan mobilitas bisnis Anda.');
