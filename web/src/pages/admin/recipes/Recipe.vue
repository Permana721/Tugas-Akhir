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
        <h2 class="text-xl md:text-2xl font-bold text-blue-600">Kelola Resep</h2>
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
          <button
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto"
            @click="showAddForm = true"
          >
            + Tambah Resep
          </button>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari resep..."
            class="w-full md:w-56 p-2 border rounded-md"
          />
        </div>
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-blue-50">
              <th class="p-4 border-b-2">Nama Resep</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Deskripsi</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Durasi</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Jumlah Bahan</th>
              <th class="p-4 border-b-2">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="recipe in filteredRecipes" :key="recipe.id" class="hover:bg-gray-50">
              <td class="p-4 border-b">{{ recipe.name }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ recipe.description }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ recipe.duration }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ recipe.material }}</td>
              <td class="p-4 border-b flex flex-col md:flex-row space-y-2 md:space-y-0 md:space-x-2">
                <button
                  class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto"
                  @click="editRecipe(recipe)"
                >
                  Edit
                </button>
                <button
                  class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 w-full md:w-auto"
                  @click="deleteRecipe(recipe.id)"
                >
                  Hapus
                </button>
              </td>
            </tr>
            <tr v-if="filteredRecipes.length === 0">
              <td colspan="4" class="text-center p-4 text-gray-500">
                Tidak ada resep ditemukan.
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
      newRecipe: {
        name: "",
        description: "",
        duration: "",
        material: "",
        fat: "",
      },
      recipes: [
        { id: 1, name: "Manfaat Makan Sehat", description: "200", duration: "15", material: "30", fat: "10" },
        { id: 2, name: "Tips Diet Seimbang", description: "250", duration: "20", material: "40", fat: "15" },
        { id: 3, name: "Olahraga yang Baik", description: "180", duration: "10", material: "25", fat: "8" },
      ],
    };
  },
  computed: {
    filteredRecipes() {
      return this.recipes.filter((recipe) =>
        recipe.name.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
  },
  methods: {
    addRecipe() {
      if (!this.newRecipe.name || !this.newRecipe.description) {
        alert("Nama resep dan kalori harus diisi!");
        return;
      }
      this.recipes.push({
        id: this.recipes.length + 1,
        ...this.newRecipe,
      });
      this.newRecipe = {
        name: "",
        description: "",
        duration: "",
        material: "",
        fat: "",
      };
      this.showAddForm = false;
    },
    editRecipe(recipe) {
      alert(`Edit resep: ${recipe.name}`);
    },
    deleteRecipe(id) {
      if (confirm("Apakah Anda yakin ingin menghapus resep ini?")) {
        this.recipes = this.recipes.filter((recipe) => recipe.id !== id);
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
