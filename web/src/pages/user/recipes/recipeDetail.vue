<template>
    <Navbar />
    <div class="pt-28 max-w-2xl mx-auto min-h-screen p-4">
        <!-- Header -->
        <div class="bg-white rounded-lg shadow-md p-4 md:p-6 text-center">
            <img src="/img/background/background.png" alt="Blueberry Oatmeal" class="w-full h-40 object-cover rounded-lg mb-4">
            <h1 class="text-xl font-bold text-gray-800">Blueberry Oatmeal | di Air</h1>
            <p class="text-gray-500 text-sm">🕒 Waktu: 10 min | 📜 Tingkat: Dasar | ✍️ Penulis: Yamfit</p>
        </div>
    
        <!-- Nutrisi -->
        <div class="bg-white rounded-lg shadow-md mt-6 p-4 md:p-6 text-center">
            <h2 class="text-lg font-bold mb-4">Nutrisi</h2>
            <div class="grid grid-cols-4 gap-2 text-center">
                <div class="p-2 bg-gray-100 rounded-lg">
                    <p class="text-lg font-bold text-green-500">227</p>
                    <p class="text-sm">Kalori</p>
                </div>
                <div class="p-2 bg-gray-100 rounded-lg">
                    <p class="text-lg font-bold text-blue-500">9g</p>
                    <p class="text-sm">Protein</p>
                </div>
                <div class="p-2 bg-gray-100 rounded-lg">
                    <p class="text-lg font-bold text-yellow-500">4g</p>
                    <p class="text-sm">Lemak</p>
                </div>
                <div class="p-2 bg-gray-100 rounded-lg">
                    <p class="text-lg font-bold text-red-500">40g</p>
                    <p class="text-sm">Karbohidrat</p>
                </div>
            </div>
        </div>
    
        <!-- Bahan -->
        <div class="bg-white rounded-lg shadow-md mt-6 p-4 md:p-6">
            <h2 class="text-lg font-bold mb-4">Bahan</h2>
            <ul class="space-y-2">
                <li class="flex justify-between border-b pb-2">
                    <span>Gandum</span>
                    <span class="text-sm text-gray-500">55g (0.5 cangkir)</span>
                </li>
                <li class="flex justify-between border-b pb-2">
                    <span>Blueberi</span>
                    <span class="text-sm text-gray-500">28.3g (1 ons)</span>
                </li>
                <li class="flex justify-between">
                    <span>Air</span>
                    <span class="text-sm text-gray-500">Adct</span>
                </li>
            </ul>
        </div>
    
        <!-- Catatan dan Instruksi -->
        <div class="bg-white rounded-lg shadow-md mt-6 p-4 md:p-6">
            <h2 class="text-lg font-bold mb-4">Catatan</h2>
            <div class="bg-gray-100 p-3 rounded-lg text-sm text-gray-600">
            Anda bisa menggunakan blueberry segar atau beku.
            </div>
            <h2 class="text-lg font-bold mt-6 mb-4">Instruksi</h2>
            <ul class="space-y-4">
                <li class="flex items-center space-x-3">
                    <span class="bg-blue-600 text-white px-3 py-1 rounded-full text-sm">1</span>
                    <p class="text-gray-700">Taruh oat mentah di piring tinggi</p>
                </li>
                <li class="flex items-center space-x-3">
                    <span class="bg-blue-600 text-white px-3 py-1 rounded-full text-sm">2</span>
                    <p class="text-gray-700">Tambahkan air mendidih, dan tutupi oat dengan air matang di atasnya</p>
                </li>
                <li class="flex items-center space-x-3">
                    <span class="bg-blue-600 text-white px-3 py-1 rounded-full text-sm">3</span>
                    <p class="text-gray-700">Taruh blueberry di atasnya</p>
                </li>
            </ul>
        </div>
    
        <!-- Nutrisi Mikro -->
        <div class="bg-white rounded-lg shadow-md mt-6 p-4 md:p-6">
            <h2 class="text-xl font-bold mb-4">Nutrisi Mikro</h2>
            <div class="border-t border-gray-200 pt-4 grid grid-cols-1 md:grid-cols-2 gap-6 text-sm text-gray-700">
                <div>
                    <p class="font-semibold mb-3">Berat keseluruhan: <span class="font-bold">83.35g</span></p>
                    <p class="font-semibold">Karbohidrat: <span class="font-bold">40.39g</span></p>
                    <p>Serat makanan: 6.28g</p>
                    <p class="mb-3">Gula: 3.02g</p>
                    <p class="font-semibold">Lemak: <span class="font-bold">3.61g</span></p>
                    <p>Jenuh: 0.63g</p>
                    <p>Poli: 0.86g</p>
                    <p class="mb-3">Mono: 0.78g</p>
                    <p class="font-semibold mt-7">Protein: <span class="font-bold">8.6g</span></p>
                </div>
                <div class="md:text-right">
                    <p>Kalsium: <span>2.8%</span></p>
                    <p>Besi: <span>12.22%</span></p>
                    <p>Magnesium: <span>0.13%</span></p>
                    <p>Vitamin A: <span>0.31%</span></p>
                    <p>Vitamin C: <span>2.52%</span></p>
                    <p>Vitamin D: <span>0mg</span></p>
                    <p>Vitamin B6: <span>0mg</span></p>
                    <p>Vitamin B12: <span>0mg</span></p>
                    <p>Kolesterol: <span>0mg</span></p>
                    <p>Sodium: <span>0.09%</span></p>
                    <p>Kalium: <span>5.21%</span></p>
                </div>
            </div>
        </div>
    </div>
    <Footer />
</template>

<script>
import Navbar from '../components/Navbar.vue';
import Footer from '../components/Footer.vue';

export default {
    components: {
        Navbar,
        Footer,
    },
};
</script>

<style>
body {
    font-family: "Inter", sans-serif;
}
</style>