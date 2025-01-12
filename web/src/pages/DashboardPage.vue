<template>
    <Navbar />
    <div>
        <div class="flex-1 pt-28 flex flex-col items-center px-6 py-12 md:flex-row md:justify-center md:gap-10">
            <div class="text-center md:text-left md:max-w-md">
                <h1 class="text-4xl font-bold text-gray-800 mb-4">
                    Pola makan baik pangkal tubuh sehat.
                </h1>
                <p class="text-lg text-gray-600 mb-6">
                Ingin lebih menjaga pola makan? Catat santapan, pelajari kebiasaan, dan capai sasaran kamu dengan FUFUFAFA.
                </p>
                <RouterLink to="/login" class="px-6 py-3 bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-700 transition duration-300">
                    Mulai dengan gratis
                </RouterLink>
            </div>
            <div class="mt-8 md:mt-0 md:ml-10">
                <img src="/img/background/background.png" alt="Healthy Meal" class="w-full max-w-md rounded-lg"/>
            </div>
        </div>

        <div class="py-16 px-6">
            <div class="max-w-6xl mx-auto flex flex-col items-center md:flex-row md:gap-10">
                <div class="relative w-96 h-72">
                    <img src="/img/background/blueberry.png" alt="Blueberries" class="absolute inset-0 w-full h-full object-cover" />
                </div>
                <div class="mt-10 text-center md:text-left">
                    <h2 class="text-3xl font-bold text-gray-800 mb-4">Catat lebih dari 14 juta makanan.</h2>
                    <p class="text-lg text-gray-600">
                        Lihat perincian kalori dan nutrisi, bandingkan ukuran porsi, dan lihat bagaimana makanan yang Anda makan mewujudkan tujuan Anda.
                    </p>
                </div>
            </div>
        </div>

        <div class="bg-white py-12 px-6">
            <div class="max-w-4xl mx-auto text-center">
                <h2 class="text-3xl font-bold text-gray-800 mb-4">Alat untuk Target Anda</h2>
                <p class="text-gray-600 text-lg mb-8">
                    Ingin menurunkan berat badan, meningkatkan kebugaran, menurunkan BMI, atau memulai pola hidup sehat? Kami punya fitur yang tepat untuk membantu Anda mewujudkannya.
                </p>
            </div>
            <div class="max-w-5xl mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 text-center">
                <div>
                    <div class="flex justify-center mb-4">
                        <div class="bg-blue-100 p-4 rounded-full">
                            <img src="/img/icon/open-book.png" alt="Buku Harian" class="w-12 h-12" />
                        </div>
                    </div>
                    <h3 class="text-xl font-semibold text-gray-800 mb-2">Pelajari, Lacak, Tingkatkan.</h3>
                    <p class="text-gray-600">
                        Menyimpan buku harian makanan Anda akan membantu memahami kebiasaan Anda dan memperbesar kemungkinan Anda mendapatkan tujuan Anda.
                    </p>
                </div>
            <div>
                <div class="flex justify-center mb-4">
                    <div class="bg-blue-100 p-4 rounded-full">
                        <img src="/img/icon/scanner.png" alt="Kode Batang" class="w-12 h-12" />
                    </div>
                </div>
                <h3 class="text-xl font-semibold text-gray-800 mb-2">Pencatatan jadi lebih mudah.</h3>
                <p class="text-gray-600">
                    Pindai barcode, simpan makanan dan resep, dan gunakan Alat Cepat untuk pelacakan makanan yang cepat dan mudah.
                </p>
            </div>
            <div>
                <div class="flex justify-center mb-4">
                    <div class="bg-blue-100 p-4 rounded-full">
                        <img src="/img/icon/sneakers.png" alt="Sepatu" class="w-12 h-12" />
                    </div>
                </div>
                <h3 class="text-xl font-semibold text-gray-800 mb-2">Tetap Termotivasi.</h3>
                <p class="text-gray-600">
                    Bergabung dengan Komunitas Fitness Terbesar di Dunia untuk mendapatkan saran, kiat, dan dukungan 24/7.
                </p>
            </div>
            </div>
        </div>

        <div class="py-12 px-6">
            <div class="max-w-4xl mx-auto text-center">
                <h2 class="text-3xl font-bold text-gray-800 mb-4">Resep & Inspirasi</h2>
                <p class="text-gray-600 text-lg mb-8">
                    Dapatkan resep yang disetujui ahli gizi dan kiat olahraga memotivasi dari para ahli.
                </p>
            </div>
            <div class="max-w-5xl mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
                <RouterLink to="/blog" class="bg-white shadow-md rounded-lg overflow-hidden">
                    <img src="/img/background/menu-diet.jpeg" alt="Semangkuk oatmeal dengan kranberi" class="w-full h-40 object-cover" />
                    <div class="p-4">
                        <h3 class="text-xl font-semibold text-gray-800 mb-2">15 Sarapan Cepat di Bawah 300 Kalori</h3>
                        <a href="#" class="text-blue-600 hover:underline">Blog FUFUFAFA</a>
                    </div>
                </RouterLink>
                <RouterLink to="/blog" class="bg-white shadow-md rounded-lg overflow-hidden">
                    <img src="/img/background/naspad.png" alt="Pengaturan berbagai cheat meal" class="w-full h-40 object-cover" />
                    <div class="p-4">
                        <h3 class="text-xl font-semibold text-gray-800 mb-2">Masalah Cheat Day</h3>
                        <a href="#" class="text-blue-600 hover:underline">Blog FUFUFAFA</a>
                    </div>
                </RouterLink>
                <RouterLink to="/blog" class="bg-white shadow-md rounded-lg overflow-hidden">
                    <img src="/img/background/jogging.jpg" alt="Seorang wanita jogging di kota" class="w-full h-40 object-cover" />
                    <div class="p-4">
                        <h3 class="text-xl font-semibold text-gray-800 mb-2">Panduan Penting agar Tetap Bergerak</h3>
                        <a href="#" class="text-blue-600 hover:underline">Blog FUFUFAFA</a>
                    </div>
                </RouterLink>
            </div>
        </div>

        <div class="px-6 py-12">
            <div class="text-center max-w-4xl mx-auto mb-10">
                <h1 class="text-4xl font-bold text-gray-800 mb-4">Jelajahi FUFUFAFA</h1>
            </div>

            <div class="relative bg-white rounded-lg shadow-lg mx-auto max-w-6xl p-8 md:p-16 lg:p-40 overflow-hidden">
                <img src="/img/background/search-food.jpg" alt="Background Image" class="absolute top-0 left-0 w-full h-full rounded-lg z-0"/>
                <div class="relative z-10 text-center bg-opacity-80 md:bg-opacity-0 p-4 md:p-0 rounded-lg">
                    <h2 class="text-2xl md:text-4xl font-bold text-gray-800 mb-4">Cari Makanan</h2>
                    <p class="text-sm md:text-lg mx-auto mb-7">
                    Penasaran dengan kalori dan nutrisi dalam makanan favorit Anda? Gunakan alat pencarian kami untuk
                    menemukan informasi nutrisi tentang makanan apa pun secara instan.
                    </p>
                    <div class="flex flex-col md:flex-row items-center justify-center gap-4">
                        <input type="text" placeholder="Masukkan makanan untuk mencari" class="w-full max-w-lg px-4 py-2 border rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:outline-none"/>
                        <button class="px-6 py-2 bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-700 transition duration-300">
                            Cari
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <Footer />
</template>

<script>
import Navbar from '../components/Navbar.vue';
import Footer from '../components/Footer.vue';
// import { RouterLink } from 'vue-router';

export default {
    components: {
        Navbar,
        Footer,
    },
};
</script>