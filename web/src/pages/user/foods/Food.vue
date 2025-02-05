<template>
    <div>
      <h1 class="text-2xl font-bold text-center my-4">Daftar Makanan dan Nilai Gizi</h1>
      <div class="max-w-7xl mx-auto">
        <input
          type="text"
          v-model="searchQuery"
          @input="searchFood"
          class="p-2 border border-gray-300 rounded-lg w-full mb-4"
          placeholder="Cari bahan makanan..."
        />
        <table class="min-w-full table-auto">
          <thead>
            <tr>
              <th class="border px-4 py-2">Nama Makanan</th>
              <th class="border px-4 py-2">Kalori</th>
              <th class="border px-4 py-2">Karbohidrat</th>
              <th class="border px-4 py-2">Lemak</th>
              <th class="border px-4 py-2">Protein</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="food in foodList" :key="food.code">
              <td class="border px-4 py-2">{{ food.product_name }}</td>
              <td class="border px-4 py-2">{{ food.nutrition_data.calories }} kkal</td>
              <td class="border px-4 py-2">{{ food.nutrition_data.carbohydrates }} g</td>
              <td class="border px-4 py-2">{{ food.nutrition_data.fat }} g</td>
              <td class="border px-4 py-2">{{ food.nutrition_data.protein }} g</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        searchQuery: "",
        foodList: [],
      };
    },
    methods: {
      async searchFood() {
        if (this.searchQuery.length < 3) return; // Minimal 3 karakter untuk pencarian
  
        // Open Food Facts API endpoint
        const url = `https://world.openfoodfacts.org/cgi/search.pl?search_terms=${this.searchQuery}&search_simple=1&action=process&json=1`;
  
        try {
          const response = await fetch(url);
          const data = await response.json();
  
          // Menyiapkan data yang diperlukan
          this.foodList = data.products.map(food => ({
            code: food.code,
            product_name: food.product_name || "Nama tidak ditemukan",
            nutrition_data: {
              calories: food.nutriments?.energy_kcal || "Data tidak tersedia",
              carbohydrates: food.nutriments?.carbohydrates_100g || "Data tidak tersedia",
              fat: food.nutriments?.fat_100g || "Data tidak tersedia",
              protein: food.nutriments?.proteins_100g || "Data tidak tersedia",
            }
          }));
        } catch (error) {
          console.error("Error fetching food data:", error);
        }
      },
    },
  };
  </script>
  