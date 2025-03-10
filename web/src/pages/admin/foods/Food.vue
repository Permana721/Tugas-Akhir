<template>
  <div class="flex flex-col md:flex-row min-h-screen bg-gray-100">
    <aside
      class="inset-y-0 left-0 w-64 bg-white shadow-lg transition-transform transform md:translate-x-0 z-50
            fixed md:static"
      :class="{ '-translate-x-full': !isSidebarOpen }"
    >
      <Sidebar />
    </aside>
    <div
      v-if="isSidebarOpen"
      class="fixed inset-0 bg-black bg-opacity-50 md:hidden"
      @click="isSidebarOpen = false"
    ></div>
    <div class="flex-1 p-4 md:p-6">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-xl md:text-2xl font-bold text-blue-600">Kelola Makanan</h2>
        <button 
          @click="isSidebarOpen = !isSidebarOpen" 
          class="md:hidden text-blue-600 transition-transform duration-300 ease-in-out"
          :class="{ 'rotate-90': isSidebarOpen }"
        >
          ☰
        </button>
      </div>
      <div class="bg-white p-4 shadow-md rounded-md overflow-x-auto">
        <div class="flex flex-col md:flex-row md:justify-between md:items-center gap-2 md:gap-0 mb-4">
          <router-link
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto"
            @click="showAddForm = true"
            to="/admin-pg/food/add-food"
          >
            + Tambah Makanan
          </router-link>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari makanan..."
            class="w-full md:w-56 p-2 border rounded-md"
          />
        </div>
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-blue-50">
              <th class="p-4 border-b-2">Nama Makanan</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Kalori</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Protein</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Karbohidrat</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Lemak</th>
              <th class="p-4 border-b-2">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="food in filteredFoods" :key="food.id" class="hover:bg-gray-50">
              <td class="p-4 border-b">{{ food.name }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ food.calorie }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ food.protein }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ food.carbohydrate }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ food.fat}}</td>
              <td class="p-4 border-b flex flex-col md:flex-row space-y-2 md:space-y-0 md:space-x-2">
                <button
                  class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto"
                  @click="editFood(food)"
                >
                  Edit
                </button>
                <button
                  class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 w-full md:w-auto"
                  @click="deleteFood(food.id)"
                >
                  Hapus
                </button>
              </td>
            </tr>
            <tr v-if="filteredFoods.length === 0">
              <td colspan="4" class="text-center p-4 text-gray-500">
                Tidak ada makanan ditemukan.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import Sidebar from "../components/Sidebar.vue";

export default {
  components: {
    Sidebar,
  },
  data() {
    return {
      isSidebarOpen: false,
      showAddForm: false,
      searchQuery: "",
      newFood: {
        name: "",
        calorie: "",
        protein: "",
        carbohydrate: "",
        fat: "",
      },
      foods: [
        { id: 1, name: "Manfaat Makan Sehat", calorie: "Admin", protein: "2025-02-03", carbohydrate: "2025-02-03", fat: "2025-02-03" },
        { id: 2, name: "Tips Diet Seimbang", calorie: "Admin", protein: "2025-02-01", carbohydrate: "2025-02-03", fat: "2025-02-03" },
        { id: 3, name: "Olahraga yang Baik", calorie: "Admin", protein: "2025-01-29", carbohydrate: "2025-02-03", fat: "2025-02-03" },
      ],
    };
  },
  computed: {
    filteredFoods() {
      return this.foods.filter((food) =>
        food.name.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
  },
  methods: {
    addFood() {
      if (!this.newFood.name || !this.newFood.protein) {
        alert("Judul dan tanggal harus diisi!");
        return;
      }
      this.foods.push({
        id: this.foods.length + 1,
        name: this.newFood.name,
        calorie: this.newFood.calorie,
        protein: this.newFood.protein,
      });
      this.newFood.name = "";
      this.newFood.protein = new protein().toISOString().substr(0, 10);
      this.showAddForm = false;
    },
    editFood(food) {
      alert(`Edit food: ${food.name}`);
    },
    deleteFood(id) {
      if (confirm("Apakah Anda yakin ingin menghapus food ini?")) {
        this.foods = this.foods.filter((food) => food.id !== id);
      }
    },
  },
};
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  padding: 12px;
}
</style>
