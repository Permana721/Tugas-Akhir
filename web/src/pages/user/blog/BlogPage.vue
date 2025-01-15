<template>
    <Navbar />
    <div class="min-h-screen py-16 px-6  flex-1 pt-32 flex flex-col">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <h1 class="text-4xl font-bold text-gray-800 text-center mb-12">Blog Kesehatan & Kebugaran</h1>
        </div>

        <div class="max-w-3xl mx-auto grid gap-8 px-4 sm:px-6 lg:px-8">
            <RouterLink to="/blog" class="bg-white rounded-lg shadow-md overflow-hidden">
                <img src="/img/blog/nasi.png" alt="Apakah Nasi Adalah Penyebab Gagal Diet?" class="w-full"/>
                <div class="p-6">
                    <div class="text-blue-600 text-sm font-medium mb-2">
                        Diet & Nutrition
                    </div>

                    <h2 class="text-2xl font-bold text-gray-800 mb-2 hover:text-blue-600">
                        Apakah Nasi Adalah Penyebab Gagal Diet?
                    </h2>

                    <p class="text-gray-600 mb-4">
                        Banyak yang percaya bahwa nasi adalah penyebab utama berat badan sulit turun. Namun, apakah benar nasi harus dihindari saat diet? Pelajari faktanya di sini!
                    </p>

                    <RouterLink to="/blog" class="text-blue-600 font-semibold hover:underline">
                        Baca Selengkapnya
                    </RouterLink>
                </div>
            </RouterLink>

            <RouterLink to="/blog" class="bg-white rounded-lg shadow-md overflow-hidden">
                <img src="/img/blog/roti-nasi.jpg" alt="Apakah Roti Lebih Sehat Daripada Nasi?" class="w-full"/>
                <div class="p-6">
                    <div class="text-blue-600 text-sm font-medium mb-2">
                    Nutrition
                    </div>

                    <h2 class="text-2xl font-bold text-gray-800 mb-2 hover:text-blue-600">
                    Apakah Roti Lebih Sehat Daripada Nasi?
                    </h2>

                    <p class="text-gray-600 mb-4">
                    Roti dan nasi adalah makanan pokok yang sering kita konsumsi. Namun, manakah yang lebih sehat untuk mendukung pola makan Anda? Pelajari kelebihan dan kekurangan dari keduanya.
                    </p>

                    <RouterLink to="/blog/bread-vs-rice" class="text-blue-600 font-semibold hover:underline">
                    Baca Selengkapnya
                    </RouterLink>
                </div>
            </RouterLink>

            <RouterLink to="/blog" class="bg-white rounded-lg shadow-md overflow-hidden">
                <img src="/img/blog/olahraga.jpg" alt="Jika Ingin Sehat, Apakah Wajib Olahraga?" class="w-full"/>
                <div class="p-6">
                    <div class="text-blue-600 text-sm font-medium mb-2">
                        Health & Fitness
                    </div>
                    <h2 class="text-2xl font-bold text-gray-800 mb-2 hover:text-blue-600">
                        Jika Ingin Sehat, Apakah Wajib Olahraga?
                    </h2>
                    <p class="text-gray-600 mb-4">
                        Banyak yang berpikir bahwa olahraga adalah syarat utama untuk hidup sehat. Namun, apakah benar begitu? Temukan jawaban dan panduan untuk hidup sehat tanpa beban.
                    </p>
                    <RouterLink to="/blog" class="text-blue-600 font-semibold hover:underline">
                        Baca Selengkapnya
                    </RouterLink>
                </div>
            </RouterLink>
        </div>
    </div>
    <Footer />
</template>

<script>
import Navbar from '../../../components/Navbar.vue';
import Footer from '../../../components/Footer.vue';

export default {
    components: {
        Navbar,
        Footer,
    },
};
</script>