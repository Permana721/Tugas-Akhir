<template>
    <Navbar />
    <div class="pt-24 min-h-screen flex justify-center items-center py-8 px-4">
      <div class="w-full max-w-4xl">
        <h1 class="text-xl font-bold mb-6">Resep</h1>
  
        <!-- Grid Resep -->
        <div class="grid cursor-pointer grid-cols-1 md:grid-cols-2 gap-6">
          <template v-for="(resep, index) in reseps" :key="index">
            <!-- Bungkus item dengan router-link -->
            <router-link
            :to="resep.url"
            class="block bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow"
            >

                <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow cursor-pointer">
                    <img :src="resep.image" alt="Gambar Resep" class="w-full h-32 object-cover" />
                    <div class="p-4">
                    <p class="text-gray-500 text-sm font-medium mb-1">{{ resep.calories }} cal</p>
                    <h2 class="text-lg font-bold text-gray-800 truncate">{{ resep.title }}</h2>
                    <div class="flex justify-between items-center mt-4 text-sm text-gray-600">
                        <div class="flex items-center space-x-2">
                        <div>🧪 {{ resep.protein }}g</div>
                        <div>🔥 {{ resep.fat }}g</div>
                        <div>🍞 {{ resep.carbs }}g</div>
                        </div>
                        <div>⏱️ {{ resep.time }}</div>
                    </div>
                    </div>
                </div>
            </router-link>
          </template>
        </div>
      </div>
    </div>
    <Footer />
  </template>
  
  <script>
  import Navbar from "../components/Navbar.vue";
  import Footer from "../components/Footer.vue";
  
  export default {
    components: {
      Navbar,
      Footer,
    },
    name: "RecipeGrid",
    data() {
      return {
        reseps: [
          {
            title: "Ikan trout mentega lemon | dengan asparagus",
            image: "/img/background/naspad.png",
            calories: 753,
            protein: 46,
            fat: 48,
            carbs: 49,
            time: "30 min",
            url: "/recipe/ikan-trout",
          },
          {
            title: "Taco vegan tanpa daging",
            image: "/img/background/naspad.png",
            calories: 337,
            protein: 14,
            fat: 17,
            carbs: 36,
            time: "30 min",
            url: "/recipe/taco-vegan",
          },
          {
            title: "Blueberry oatmeal | di air",
            image: "/img/background/background.png",
            calories: 227,
            protein: 9,
            fat: 4,
            carbs: 40,
            time: "10 min",
            url: "/recipe/detail",
          },
          {
          title: "Bakso swedia | di atas egg noodles & baby por",
          image: "/img/background/naspad.png",
          calories: 465,
          protein: 16,
          fat: 18,
          carbs: 63,
          time: "40 min",
        },
        {
          title: "Tilapia gaya louisiana | dengan potato wedges",
          image: "/img/background/naspad.png",
          calories: 773,
          protein: 50,
          fat: 29,
          carbs: 88,
          time: "40 min",
        },
        {
          title: "Makanan penutup krim kopi",
          image: "/img/background/naspad.png",
          calories: 69,
          protein: 6.2,
          fat: 1.9,
          carbs: 9.6,
          time: "0 min",
        },
        {
          title: "Bungkus sayuran gandum utuh",
          image: "/img/background/naspad.png",
          calories: 374,
          protein: 13,
          fat: 15,
          carbs: 48,
          time: "20 min",
        },
        {
          title: "Roti panggang apel & selai kacang",
          image: "/img/background/naspad.png",
          calories: 295,
          protein: 12,
          fat: 18,
          carbs: 26,
          time: "5 min",
        },
        {
          title: "Cangkir oatmeal kacang pisang panggang",
          image: "/img/background/naspad.png",
          calories: 160,
          protein: 5.1,
          fat: 5.4,
          carbs: 23,
          time: "65 min",
        },
        {
          title: "Bibimbap babi dan sayuran | dengan zucchini",
          image: "/img/background/naspad.png",
          calories: 813,
          protein: 28,
          fat: 35,
          carbs: 69,
          time: "30 min",
        },
          // Resep lainnya
        ],
      };
    },
  };
  </script>
  
  <style>
  body {
    font-family: "Inter", sans-serif;
  }
  </style>
  